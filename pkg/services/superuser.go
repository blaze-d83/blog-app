package services

import (
	"fmt"
	"log"
	"os"

	"github.com/blaze-d83/blog-app/pkg/auth"
	"github.com/blaze-d83/blog-app/pkg/mysql"
	"github.com/blaze-d83/blog-app/pkg/types"
	"github.com/spf13/cobra"
)

var dbInstance *mysql.Database

var (
	username string
	email    string
	password string
)

var rootCmd = &cobra.Command{
	Use:   "supercli",
	Short: "A CLI tool to manage superusers",
}

var superuserCmd = &cobra.Command{
	Use:   "create superuser",
	Short: "Create a superuser for admin access",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("Username: ")
		fmt.Scanln(&username)

		fmt.Print("Email: ")
		fmt.Scanln(&email)

		fmt.Print("Password: ")
		fmt.Scanln(&password)

		var confirmPassword string
		fmt.Print("Re-type Password: ")
		fmt.Scanln(&confirmPassword)

		if password != confirmPassword {
			log.Fatalf("Password do not match, please try again")
			os.Exit(1)
		}

		dbInstance := mysql.InitDB()
		defer func() {
			sqlDB, err := dbInstance.DB.DB()
			if err != nil {
				log.Fatalf("Failed to obtain db connection  for closing: %v", err)
			}
			if err := sqlDB.Close(); err != nil {
				log.Fatalf("Failed to close the database connection: %v", err)
			}
		}()

		hashedPassword := auth.HashPassword(password)

		user := types.Admin{
			Username: username,
			Email:    email,
			Pass:     string(hashedPassword),
		}

		if err := dbInstance.DB.Create(&user).Error; err != nil {
			log.Fatalf("Failed to create superuser: %v", err)
		}

		fmt.Printf("Superuser %s created successfully\n", username)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func InitCmd() {
	rootCmd.AddCommand(superuserCmd)
}
