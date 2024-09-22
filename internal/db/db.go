package db

import (
	"fmt"
	"log"

	"github.com/blaze-d83/blog-app/internal/config"
	"github.com/blaze-d83/blog-app/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func InitDB() *Database {

	dsn := GetDSN()
	db, err := NewMySQLStorage(dsn)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	err = db.DB.AutoMigrate(types.Admin{}, types.Category{}, types.Post{})
	if err != nil {
		log.Fatal("Failed to run migrations", err)
	}
	return db

}

func NewMySQLStorage(dsn string) (*Database, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}
	return &Database{DB: db}, err
}

func GetDSN() string {
	dbConfig := config.LoadConfig()
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.DBConfig.User,
		dbConfig.DBConfig.Pass,
		dbConfig.DBConfig.Net,
		dbConfig.DBConfig.Addr,
		dbConfig.DBConfig.DBName)

	return dsn
}
