package main

import (
	http2 "gopaseto/internal/controller/http"
	"gopaseto/internal/core/config"
	"gopaseto/internal/core/server/http"
	"gopaseto/internal/core/service"
	"gopaseto/internal/infra/logger"
	"gopaseto/internal/infra/storages"
	"gopaseto/internal/infra/storages/repository"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func main() {
	// Load Config

	conf, err := config.New()

	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
	}

	logger.Set(conf.APP)

	slog.Info("Starting the application", "app", conf.APP.Name, "env", conf.APP.Env)

	// Disable debug mode in production
	if conf.HTTP.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// cors
	ginConfig := cors.DefaultConfig()
	allowedOrigins := conf.HTTP.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	instance := gin.New()
	instance.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	// Start Database

	db, err := storages.NewDB(
		config.DB{
			Driver:       conf.DB.Driver,
			Url:          conf.DB.Url,
			MaxLifeTime:  conf.DB.MaxLifeTime,
			MaxOpenConn:  conf.DB.MaxOpenConn,
			MaxIddleConn: conf.DB.MaxIddleConn,
		},
	)

	if err != nil {
		slog.Error("error initializing database connection", "error", db)
		os.Exit(1)
	}

	// Dependensi Injection

	userRepo := repository.NewUsersRepository(db)
	userService := service.NewUserService(userRepo)
	userController := http2.NewUserController(instance, userService)

	userController.InitRouter()

	// Start Server
	httpServer := http.NewHttpServer(
		instance,
		config.HttpServerConfig{
			Port: 8000,
		},
	)

	httpServer.Start()

	defer func(h http.HttpServer) {
		err := h.Close()

		if err != nil {
			log.Printf("failed to close server %v", err)
		}
	}(httpServer)

	slog.Info("Listenning signals...")

	c := make(chan os.Signal, 1)

	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)

	<-c
	log.Println("gracefull shutdown")
}
