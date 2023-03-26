package config

import (
	"os"

	"log"

	"github.com/osscameroon/opencollective/billing/internal/graphql"
	"github.com/spf13/viper"
)

var conf Config = newDefaultConfig()

type Env struct {
	//OCKey is the key to access the OpenCollective API
	OCKey string `mapstructure:"OPEN_COLLECTIVE_API_KEY"`
	OCURL string `mapstructure:"OPEN_COLLECTIVE_API_URL"`
}

//Config contains the entire cli dependencies
type Config struct {
	gql graphql.IClient
	env Env
}

func GetEnv() Env {
	return conf.env
}

//NewDefaultConfig creates a new default config
func newDefaultConfig() Config {
	url := os.Getenv("OPEN_COLLECTIVE_API_URL")
	key := os.Getenv("OPEN_COLLECTIVE_API_KEY")

	conf := Config{
		gql: graphql.NewClient(url, key),
		env: loadEnv(),
	}

	return conf
}

func loadEnv() Env {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading .env file", err)
	}

	env := Env{}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal(err)
	}

	return env
}
