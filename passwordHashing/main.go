package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	for {
		pwd := getPwd()
		hash := hasAndSalt(pwd)

		pwd2 := getPwd()
		pwdMatch := comparePasswords(hash, pwd2)

		fmt.Println("passwords match?", pwdMatch)
	}

}

func getPwd() []byte {
	fmt.Println("Enter your password: ")

	var pwd string
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}

	return []byte(pwd)
}

func hasAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPw string, plainPw []byte) bool {
	byteHash := []byte(hashedPw)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPw)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
