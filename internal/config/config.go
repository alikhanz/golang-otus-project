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
	AMQPHost     string
	AMQPPort     int
	AMQPLogin    string
	AMQPPassword string
	AMQPVhost    string
	DBHost       string
	DBPort       int
	DBName       string
	DBUser       string
	DBPassword   string
}

func New() *Config {
	return &Config{
		GrpcPort:     intVal("GRPC_PORT"),
		AMQPHost:     os.Getenv("AMQP_HOST"),
		AMQPPort:     intVal("AMQP_PORT"),
		AMQPLogin:    os.Getenv("AMQP_LOGIN"),
		AMQPPassword: os.Getenv("AMQP_PASSWORD"),
		AMQPVhost:    os.Getenv("AMQP_VHOST"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       intVal("DB_PORT"),
		DBName:       os.Getenv("DB_DATABASE"),
		DBUser:       os.Getenv("DB_USERNAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
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
