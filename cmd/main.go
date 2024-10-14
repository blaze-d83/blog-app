package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blaze-d83/blog-app/internal/handlers"
	"github.com/blaze-d83/blog-app/internal/routes"
	"github.com/blaze-d83/blog-app/pkg/logger"
	middleware "github.com/blaze-d83/blog-app/pkg/middleware"
	"github.com/blaze-d83/blog-app/pkg/mysql"
	"github.com/blaze-d83/blog-app/pkg/services"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// Entry Point of the Application
func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Stderr, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, stdout, stderr io.Writer, args []string) error {

	/**
	 * Main logic of the application.
	 * Accepts a context, writers for stdout and stderr, and the command-line arguments (args).
	 * We also set up a signal handler to gracefully shut down the application when receiving interrupt signals.
	 */

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Intialize the database connection to MySQL
	dbInstance := mysql.InitDB()
	defer func() {
		sqlDB, err := dbInstance.DB.DB()
		if err != nil {
			fmt.Fprintf(stderr, "Failed to obtain db connection for closing: %v\n", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			fmt.Fprintf(stderr, "Failed to close database connection: %v\n", err)
		} else {
			fmt.Fprintf(stdout, "Database connection closed successfully.\n")
		}
	}()

	// CLI-based superuser creation.
	if len(args) > 2 && args[1] == "create" && args[2] == "superuser" {
		services.InitCmd(dbInstance)
		services.Execute()
		return nil
	}

	// Initialize a new echo server
	e := echo.New()

	// Setup custom logger and recovery middleware
	customLogger := logger.NewCustomLogger()

	// Initialize custom middleware
	mw := middleware.Middleware{}

	// User echo's default RequestID middleware
	e.Use(echoMiddleware.RequestID())

	// Use sutom logging middleware globally
	e.Use(mw.LoggingMiddleware(customLogger))

	// Initialize repositories (services)
	publicRepo := services.NewUserRepository(dbInstance)
	adminRepo := services.NewAdminRepository(dbInstance)

	// Initialize the handlers with repositories
	publicHandler := &handlers.PublicHandler{Repository: publicRepo}
	adminHandler := &handlers.AdminHandler{Repository: adminRepo}

	// Register routes
	routes.SetupRouter(e, adminHandler, publicHandler, mw)

	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("Shutting down server due to an error: ", err)
		}
	}()

	// Block until we receive a signal (CTRL+C or system signal)
	<-ctx.Done()

	// Graceful shutdown
	fmt.Fprintf(stdout, "Shutting down server gracefully...\n")
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()

	if err := e.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}

	fmt.Fprintf(stdout, "Server exited gracefully.\n")
	return nil

}
