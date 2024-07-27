package config

import (
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Consul  ConsulConfig
	DB      DBConfig
	Service ServiceConfig
}

type ConsulConfig struct {
	Addr           string `validate:"required"`
	DeregisterTime string `validate:"required"`
	IntervalTime   string `validate:"required"`
}

type DBConfig struct {
	DSN string `validate:"required"`
}

type ServiceConfig struct {
	Name string `validate:"required"`
	Port int    `valdiate:"required"`
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	log.Println("loaded .env file successfully")

	config := Config{
		Consul: ConsulConfig{
			Addr:           getEnv("Consul_Addr"),
			DeregisterTime: getEnv("Consul_Deregister_Time"),
			IntervalTime:   getEnv("Consul_Interval_Time"),
		},
		DB: DBConfig{
			DSN: getEnv("DB_DSN"),
		},
		Service: ServiceConfig{
			Name: getEnv("Service_Name"),
			Port: getEnvInt("Service_Port"),
		},
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return ""
}

func getEnvInt(key string) int {
	valueStr := getEnv(key)

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return -1
}
