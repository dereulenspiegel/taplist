// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Beer struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Style       *string  `json:"style"`
	Abv         float64  `json:"abv"`
	BuGuRatio   *float64 `json:"buGuRatio"`
	Ibu         *int     `json:"ibu"`
	ColorEbc    *int     `json:"colorEbc"`
	Og          *float64 `json:"og"`
	Fg          *float64 `json:"fg"`
	GravityUnit *string  `json:"gravityUnit"`
}

type BeerData struct {
	Name        string   `json:"name"`
	Abv         float64  `json:"abv"`
	BuGuRatio   *float64 `json:"buGuRatio"`
	Ibu         *int     `json:"ibu"`
	ColorEbc    *int     `json:"colorEbc"`
	Og          *float64 `json:"og"`
	Fg          *float64 `json:"fg"`
	GravityUnit *string  `json:"gravityUnit"`
}

type BrewfatherBatch struct {
	ID    string  `json:"id"`
	State *string `json:"state"`
	Beer  *Beer   `json:"beer"`
}

type Kegerator struct {
	Name    *string   `json:"name"`
	Sensors []*Sensor `json:"sensors"`
	Taps    []*Tap    `json:"taps"`
}

type Sensor struct {
	ID    string     `json:"id"`
	Type  SensorType `json:"type"`
	Name  string     `json:"name"`
	Value float64    `json:"value"`
	Unit  string     `json:"unit"`
}

type Tap struct {
	ID               string    `json:"id"`
	Number           int       `json:"number"`
	Empty            bool      `json:"empty"`
	VolumeAvailable  *float64  `json:"volumeAvailable"`
	PercentAvailable *float64  `json:"percentAvailable"`
	Name             string    `json:"name"`
	Sensors          []*Sensor `json:"sensors"`
	Beer             *Beer     `json:"beer"`
}

type TapData struct {
	Empty            bool      `json:"empty"`
	VolumeAvailable  *float64  `json:"volumeAvailable"`
	PercentAvailable *float64  `json:"percentAvailable"`
	Name             *string   `json:"name"`
	Beer             *BeerData `json:"beer"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
)

var AllRole = []Role{
	RoleAdmin,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SensorType string

const (
	SensorTypeTempRoom         SensorType = "TEMP_ROOM"
	SensorTypeTempKegerator    SensorType = "TEMP_KEGERATOR"
	SensorTypeTempTap          SensorType = "TEMP_TAP"
	SensorTypeHumidityRoom     SensorType = "HUMIDITY_ROOM"
	SensorTypeHumidtyKegerator SensorType = "HUMIDTY_KEGERATOR"
	SensorTypePressureTap      SensorType = "PRESSURE_TAP"
	SensorTypeFlowrateTap      SensorType = "FLOWRATE_TAP"
)

var AllSensorType = []SensorType{
	SensorTypeTempRoom,
	SensorTypeTempKegerator,
	SensorTypeTempTap,
	SensorTypeHumidityRoom,
	SensorTypeHumidtyKegerator,
	SensorTypePressureTap,
	SensorTypeFlowrateTap,
}

func (e SensorType) IsValid() bool {
	switch e {
	case SensorTypeTempRoom, SensorTypeTempKegerator, SensorTypeTempTap, SensorTypeHumidityRoom, SensorTypeHumidtyKegerator, SensorTypePressureTap, SensorTypeFlowrateTap:
		return true
	}
	return false
}

func (e SensorType) String() string {
	return string(e)
}

func (e *SensorType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SensorType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SensorType", str)
	}
	return nil
}

func (e SensorType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
