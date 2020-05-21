package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Overload()
}

type Config struct {
	GrpcPort     int
	AmqpHost     string
	AmqpPort     int
	AmqpLogin    string
	AmqpPassword string
	AmqpVhost    string
	DbHost       string
	DbPort       int
	DbName       string
	DbUser       string
	DbPassword   string
}

func New() *Config {
	return &Config{
		GrpcPort:     intVal("GRPC_PORT"),
		AmqpHost:     os.Getenv("AMQP_HOST"),
		AmqpPort:     intVal("AMQP_PORT"),
		AmqpLogin:    os.Getenv("AMQP_LOGIN"),
		AmqpPassword: os.Getenv("AMQP_PASSWORD"),
		AmqpVhost:    os.Getenv("AMQP_VHOST"),
		DbHost:       os.Getenv("DB_HOST"),
		DbPort:       intVal("DB_PORT"),
		DbName:       os.Getenv("DB_DATABASE"),
		DbUser:       os.Getenv("DB_USERNAME"),
		DbPassword:   os.Getenv("DB_PASSWORD"),
	}
}

func intVal(variable string) int {
	val, err := strconv.Atoi(os.Getenv(variable))
	panicOnErr(err)

	return val
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}