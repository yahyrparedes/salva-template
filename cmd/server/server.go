package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	env "github.com/yahyrparedes/salva-template/cmd/config"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

func InitializeServer() *fiber.App {
	var timeout = os.Getenv("server.timeout")
	var name = os.Getenv("server.name")
	readTimeoutSecondsCount, _ := strconv.Atoi(timeout)

	// Return Fiber configuration.
	config := fiber.Config{
		ServerHeader: name,
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
	}

	return fiber.New(config)
}

func generateUrl() string {
	var host = os.Getenv("server.host")
	var port = os.Getenv("server.port")
	return fmt.Sprintf("%s:%s", host, port)
}

func startServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	if err := a.Listen(generateUrl()); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

func startServer(a *fiber.App) {
	if err := a.Listen(generateUrl()); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

func RunServer(a *fiber.App) {
	appEnv := os.Getenv("APP_ENV")
	if strings.EqualFold(appEnv, env.Production) {
		startServerWithGracefulShutdown(a)
	} else {
		startServer(a)
	}
}
