package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "123456789"

	hp, err := hashPassword(password)
	if err != nil {
		log.Fatalln("Not logged in.")
	}

	if err := comparePassword(password, hp); err != nil {
		log.Fatalln("Not logged in.")
	}
	log.Println("Logged in!")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPassword []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return fmt.Errorf("Invalid Username and Password: %w", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	// var key = []byte{}
	// for i := 1; i <= 64; i++ {
	// 	key = append(key, byte(i))
	// }
	key := []byte("mysecretkey")
	h := hmac.New(sha512.New, key)
	if _, err := h.Write(msg); err != nil {
		return nil, fmt.Errorf("Error in singMessage while hashing message: %w", err)
	}
	signature := h.Sum(nil)
	return signature, nil
}

func checkSignature(msg, signature []byte) (bool, error) {
	newSignature, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSignature while getting signature of message: %w", err)
	}
	same := hmac.Equal(newSignature, signature)
	return same, nil
}
