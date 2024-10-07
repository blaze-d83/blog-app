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

var (
	username string
	email    string
	password string
)

var rootCmd = &cobra.Command{
	Use:   "supercli",
	Short: "A CLI tool to manage superusers",
}

// Initializes the CLI command and passes the dbinstance to be used 
func InitCmd(dbInstance *mysql.Database) {

	// dbInstance (intialized in main) is passed directly to superuserCmd.Run function via closure
	superuserCmd.Run = func(cmd *cobra.Command, args []string) {
		if err := createSuperuser(dbInstance); err != nil {
			log.Fatalf("Failed to create superuser: %v", err)
			os.Exit(1)
		}
	}
	rootCmd.AddCommand(superuserCmd)
}

var superuserCmd = &cobra.Command{
	Use:   "create superuser",
	Short: "Create a superuser for admin access",
}

// Handles the logic for creating superuser
func createSuperuser(dbInstance *mysql.Database) error {
	username, email, password, err := getSuperUserInput()
	if err != nil {
		return err
	}

	// Hashes the password before storing it
	hashedPassword := auth.HashPassword(password)

	// Create a superuser object
	superuser := types.Admin{
		Username: username,
		Email:    email,
		Pass:     string(hashedPassword),
	}

	// Save the superuser to the database
	if err := dbInstance.DB.Create(&superuser).Error; err != nil {
		return fmt.Errorf("faile to create superuser: %v", err)
	}

	fmt.Printf("Superuser %s created successfully\n", username)
	return nil
}

// Handles the user input for creating superuser
func getSuperUserInput() (string, string, string, error) {

	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Email: ")
	fmt.Scanln(&email)

	fmt.Print("Password: ")
	fmt.Scanln(&password)

	var confirmPassword string
	fmt.Print("Re-type Password: ")
	fmt.Scanln(&confirmPassword)

	// Validates if the passwords match
	if password != confirmPassword {
		return "", "", "", fmt.Errorf("password do not match, try again")
	}

	return username, email, password, nil
}

// Runs the CLI command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
