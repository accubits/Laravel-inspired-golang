package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/spf13/viper"
)

// https://www.ardanlabs.com/blog/2014/03/exportedunexported-identifiers-in-go.html
type Config struct {
	Env *viper.Viper
	Db  *gorm.DB
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
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
func initDb(env *viper.Viper) (*gorm.DB, error) {
	conn := env.GetString("database.user") + ":" + env.GetString("database.password") + "@(" + env.GetString("database.host") + ")/" + env.GetString("database.name") + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func New() (*Config, error) {
	config := Config{}
	env, err := initViper()
	db, err := initDb(env)
	config.Env = env
	config.Db = db
	if err != nil {
		return &config, err
	}
	return &config, err
}
