package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/blaze-d83/blog-app/pkg/logger"
	"github.com/blaze-d83/blog-app/pkg/mysql"
	"github.com/blaze-d83/blog-app/pkg/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if len(os.Args) > 2 && strings.ToLower(os.Args[1]) == "create" && strings.ToLower(os.Args[2]) == "superuser" {
		services.InitCmd()
		services.Execute()
		return
	}

	dbInstance := mysql.InitDB()
	defer func()  {
		sqlDB, err := dbInstance.DB.DB()
		if err != nil {
			log.Println("Failed to obtain db connection for closing: ", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Println("Failed to close the database connection: ", err)
		} else {
			log.Println("Database connection closed successfullly: ", err )
		}
	}()

	e := echo.New()

	customLoggerConfig := logger.GetCustomLoggerConfig(e)

	e.Use(middleware.LoggerWithConfig(*customLoggerConfig))
	e.Use(middleware.Recover())


	go func ()  {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down server due to an error: ", err)
		} 
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Server shutdown failed", err)
	}

	log.Println("Server exited gracefully.")

}

