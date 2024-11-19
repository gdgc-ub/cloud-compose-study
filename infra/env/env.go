package env

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv             string `mapstructure:"APP_ENV"`
	AppPort            string `mapstructure:"APP_PORT"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPass             string `mapstructure:"DB_PASS"`
	DBName             string `mapstructure:"DB_NAME"`
}

var AppEnv = getEnv()

func getEnv() *Env {
	env := &Env{}

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("ERR: " + err.Error())
	}

	if err := viper.Unmarshal(env); err != nil {
		log.Fatal("ERR: " + err.Error())
	}

	return env
}
