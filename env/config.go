
package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User   string
	Pass   string
	Net    string
	Addr   string
	DBName string
}

type Config struct {
	DBConfig DBConfig
	JWTSecret string

}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbNet := os.Getenv("DB_NET")
	dbAddr := os.Getenv("DB_ADDR")
	dbName := os.Getenv("DB_NAME")

	jwtSecret := os.Getenv("JWT_SECRET")

	return &Config{
		DBConfig: DBConfig{
			User:   dbUser,
			Pass:   dbPass,
			Net:    dbNet,
			Addr:   dbAddr,
			DBName: dbName,
		},
		JWTSecret: jwtSecret,
	}
}

