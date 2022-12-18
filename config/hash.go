package config

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)



func Hashpassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "",fmt.Errorf("failed to hash password %v",err)
	}
	return string(hashpassword), nil
}


func CheckPassword(password string, hashpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashpassword),[]byte(password))
}