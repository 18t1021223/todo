package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	EnvProfile     string `mapstructure:"ENV_PROFILE"`
	ServerPort     int    `mapstructure:"SERVER_PORT"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`

	DBHost   string `mapstructure:"DB_HOST"`
	DBPort   string `mapstructure:"DB_PORT"`
	DBUser   string `mapstructure:"DB_USER"`
	DBPass   string `mapstructure:"DB_PASS"`
	DBName   string `mapstructure:"DB_NAME"`
	DBParams string `mapstructure:"DB_PARAMS"`
	DBDriver string `mapstructure:"DB_DRIVER"`

	JWTSecretKey                    string `mapstructure:"JWT_SECRET_KEY"`
	JwtAccessTokenExpirationSeconds int64  `mapstructure:"JWT_ACCESS_TOKEN_EXPIRATION_SECONDS"`
}

func InitEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err.Error())
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal(err.Error())
	}
	return &env
}
