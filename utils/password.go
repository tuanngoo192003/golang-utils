package utils

import (
    "golang.org/x/crypto/bcrypt"
    "errors"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
    if err != nil {
        return "", errors.New("failed to hash password")
    }
    return string(bytes), nil 
}

func CheckPasswordHash (password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) 
    return err == nil
}

func ValidatePassword(password string) error {
    if len(password) < 8 {
        return errors.New("password must at least 8 characters")
    }
    return nil
}
