module.exports = {
  configureWebpack: {
    module: {
      rules: [
      ]
    }
  }
};

// or passed in via docker build arg and env var (see below dockerfile)
// baseUrl: process.env.APP_BASEPATH,