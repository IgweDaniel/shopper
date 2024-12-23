package internal

import (
	"github.com/IgweDaniel/shopper/internal/contracts"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Port int `env:"PORT,required"`
	Jwt  struct {
		Access  string `env:"JWT_ACCESS,required"`
		Refresh string `env:"JWT_REFRESH,required"`
	}
	DbUri string `env:"DB_URI,required"`
	Env   string `env:"ENV,required"`

	DbHost     string `env:"DB_HOST,required"`
	DbPort     string `env:"DB_PORT,required"`
	DbDatabase string `env:"DB_DATABASE,required"`
	DbUsername string `env:"DB_USERNAME,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbSchema   string `env:"DB_SCHEMA,required"`
}

func LoadConfig() (Config, error) {

	cfg := Config{} // 👈 new instance of `Config`
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}

	err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

type Application struct {
	Config       Config
	Repositories contracts.Repositories
}
