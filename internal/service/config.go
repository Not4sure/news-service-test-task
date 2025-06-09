package service

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MongoPort         string `mapstructure:"MONGO_PORT"`
	MongoHost         string `mapstructure:"MONGO_HOST"`
	MongoUser         string `mapstructure:"MONGO_USER"`
	MongoPassword     string `mapstructure:"MONGO_PASSWORD"`
	MongoDatabaseName string `mapstructure:"MONGO_DATABASE"`

	MongoURL string `mapstructure:"MONGO_URL"`
}

func Load(path string) (config Config, err error) {
	setDefaults()

	err = readConfigFile(path)
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	if config.MongoURL == "" {
		config.MongoURL = construnctMongoURL(
			config.MongoHost,
			config.MongoPort,
			config.MongoUser,
			config.MongoPassword,
		)
	}

	return
}

// setDefaults sets default values of config
func setDefaults() {
	defaults := map[string]any{
		"MONGO_HOST":     "localhost",
		"MONGO_PORT":     "27017",
		"MONGO_USER":     "root",
		"MONGO_PASSWORD": "",
		"MONGO_URL":      "",
		"MONGO_DATABASE": "news",
	}

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
}

func readConfigFile(path string) error {
	// add env file config
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// read env file config
	return viper.ReadInConfig()
}

// construnctMongoURL creates MonogoDB connection URL from params
func construnctMongoURL(host string, port string, user string, password string) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
}
