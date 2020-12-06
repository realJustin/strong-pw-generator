package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/gomodule/redigo/redis"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
	rand.Seed(time.Now().Unix())
	minSpecialCharacter := 2
	minNum := 2
	minUpperCase := 2
	pwLength := 16
	pw := generatePassword(pwLength, minSpecialCharacter, minNum, minUpperCase)
	clipboard.WriteAll(pw)
	storePW(pw)
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func storePW(pw string) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	_, err = conn.Do("SET", "pwkey", pw)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("PW: " + pw)
}
