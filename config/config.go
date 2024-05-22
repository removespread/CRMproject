package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	Env    string
	Server struct {
		Host int
		Port int
	}
	Database struct {
		Dsn int
	}
}

func LoadConfiguration() (*Configuration, error) {
	viper.SetConfigFile("config.yaml") // name of config file (without extension)
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")           // path to look for the config file in
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadConfig()

	if err != nil {
		return nil, errors.Wrap(err, "Error reading configuration file")
	}

	if err := viper.Unmarshal(&Configuration{}); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling configuration")
	}
	return &Configuration{}, nil

}
