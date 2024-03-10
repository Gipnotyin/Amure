package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (routes *services) createHashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("create hashed password was failed: %v", err.Error())
	}

	return string(hashPassword), nil
}
