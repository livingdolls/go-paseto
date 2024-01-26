package config

import (
	"os"

	"github.com/joho/godotenv"
)

type HttpServerConfig struct {
	Port uint
}

type (
	Container struct {
		APP  *APP
		DB   *DB
		HTTP *HTTP
	}

	APP struct {
		Name string
		Env  string
	}

	DB struct {
		Driver       string
		Url          string
		MaxLifeTime  string
		MaxOpenConn  string
		MaxIddleConn string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			return nil, err
		}
	}

	app := &APP{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Driver:       os.Getenv("DB_DRIVER"),
		Url:          os.Getenv("DB_URL"),
		MaxLifeTime:  os.Getenv("DB_MAXLIFETIME"),
		MaxOpenConn:  os.Getenv("DB_MAXOPENCONN"),
		MaxIddleConn: os.Getenv("DB_MAXIDLECONN"),
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	return &Container{
		app,
		db,
		http,
	}, nil
}
