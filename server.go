package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dereulenspiegel/go-brewchild"
	"github.com/dereulenspiegel/taplist/filestore"
	"github.com/dereulenspiegel/taplist/graph"
	"github.com/dereulenspiegel/taplist/graph/generated"
	"github.com/dereulenspiegel/taplist/graph/model"
	"github.com/gobuffalo/packr/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	defaultHttpTimeout = time.Second * 10
)

var Version = "undefined"

const (
	frontendPath = "./frontend"
)

var (
	frontendBox = packr.New("Frontend", frontendPath)
)

type ContextKey string

const (
	CtxAuthUser ContextKey = "auth-user"
)

func setDefaults() {
	viper.SetDefault("kegerator.name", "My Smart Kegerator")
	viper.SetDefault("num.taps", 2)
	viper.SetDefault("http.addr", ":8088")
	viper.SetDefault("log.level", logrus.InfoLevel.String())
	viper.SetDefault("no.auth", false)
	viper.SetDefault("http.user.header", "X-Auth-User")
	viper.SetDefault("data.path", "./data/data.json")
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	setDefaults()
	flag.Parse()

	viper.SetEnvPrefix("TAPLIST")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Warn("Failed to read configuration")
	}
	logLevel, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.WithError(err).Error("Failed to parse log level from config")
	} else {
		logrus.SetLevel(logLevel)
	}

	logrus.WithFields(logrus.Fields{
		"version":       Version,
		"numberOfTaps":  viper.GetInt("num.taps"),
		"kegeratorName": viper.GetString("kegerator.name"),
	}).Info("Starting Taplist")

	fsStore := filestore.New(viper.GetString("data.path"), viper.GetString("kegerator.name"), viper.GetInt("num.taps"))

	brewfatherUserID := viper.GetString("brewfather.userid")
	brewfatherSecret := viper.GetString("brewfather.secret")

	var bfClient *brewchild.Client
	if brewfatherSecret != "" && brewfatherUserID != "" {
		logrus.Info("Enabling Brewfather support")
		bfClient, err = brewchild.New(brewfatherUserID, brewfatherSecret)
	} else {
		logrus.Info("Brewfather credentials unconfigured")
	}

	graphqlConf := generated.Config{
		Resolvers: graph.NewResolver(fsStore, bfClient),
		Directives: generated.DirectiveRoot{
			HasRole: HasRole,
		},
	}

	var frontendHandler http.Handler
	if frontendPath := viper.GetString("frontend.path"); frontendPath != "" {
		frontendHandler = http.FileServer(http.Dir(frontendPath))
	} else {
		frontendHandler = http.FileServer(frontendBox)
	}

	mux := http.NewServeMux()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graphqlConf))
	mux.Handle("/", frontendHandler)
	mux.Handle("/query", srv)
	mux.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	httpServer := &http.Server{
		Addr:              viper.GetString("http.addr"),
		Handler:           checkUserHeader(mux),
		ReadTimeout:       defaultHttpTimeout,
		ReadHeaderTimeout: defaultHttpTimeout,
		WriteTimeout:      defaultHttpTimeout,
		IdleTimeout:       defaultHttpTimeout,
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logrus.WithField("listenAddr", viper.GetString("http.addr")).Info("Listening")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Fatal("Http server failed listening")
		}
	}()

	select {
	case <-sigs:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		httpServer.Shutdown(ctx)
	}
}

func checkUserHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if headerVal := r.Header.Get(viper.GetString("http.user.header")); headerVal != "" {
			ctx := context.WithValue(r.Context(), CtxAuthUser, headerVal)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (res interface{}, err error) {
	if viper.GetBool("no.auth") {
		return next(ctx)
	}

	if val := ctx.Value(CtxAuthUser); val != nil {
		if user, ok := val.(string); ok {
			if user == viper.GetString("admin.user") {
				next(ctx)
			}
		}
	}
	return nil, fmt.Errorf("Access denied")
}
