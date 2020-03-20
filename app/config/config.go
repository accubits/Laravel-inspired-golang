package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// https://www.ardanlabs.com/blog/2014/03/exportedunexported-identifiers-in-go.html
type Config struct {
	Env *viper.Viper
}

func initViper() (*viper.Viper, error) {

	v := viper.New()
	//Set the file name of the configurations file
	v.SetConfigName("config")
	// Set the path to look for the configurations file
	v.AddConfigPath(".")
	v.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("not found, %s", err)
		} else {
			fmt.Printf("Errors reading config file, %s", err)
		}
	}

	return v, err
}

func New() (*Config, error) {
	config := Config{}
	env, err := initViper()
	config.Env = env
	fmt.Println(config)
	if err != nil {
		return &config, err
	}
	return &config, err
}
