package bootstrap

import (
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Env struct {
	DB     DBEnv  `envPrefix:"DB_"`
	Server Server `envPrefix:"SERVER_"`
	Domain string `env:"DOMAIN"`
}

func NewEnv() *Env {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	var e Env
	if err := env.ParseWithOptions(&e, env.Options{
		RequiredIfNoDef: true,
		Prefix:          "APP_",
	}); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return &e
}
