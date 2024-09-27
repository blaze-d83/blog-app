package utils

import (
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) []byte {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}
	return hashPass
}

func GetInt(param string) uint  {
	id, _ := strconv.Atoi(param)
	return (uint(id))
}
