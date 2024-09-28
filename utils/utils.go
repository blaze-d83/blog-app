package utils

import (
	"log"
	"strconv"

	"github.com/blaze-d83/blog-app/types"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) []byte {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}
	return hashPass
}

func CompareHashPassword(admin *types.Admin, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Pass), []byte(pass)); err != nil {
		return err
	}
	return nil
}

func GetInt(param string) uint {
	id, _ := strconv.Atoi(param)
	return (uint(id))
}
