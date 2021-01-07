package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/go-redis/redis"
)

var (
	ctx            = context.Background()
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

type redisService struct {
	client *redis.Client
}

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
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 5)
	rand.Read(b)
	slug := hex.EncodeToString(b)

	// user expiretime
	userExpireTime := getArgumentTime()
	// either 24 or i
	expireTime := time.Hour * time.Duration(userExpireTime)

	fmt.Println(expireTime)

	client.Set(ctx, slug, pw, expireTime)

	fmt.Println("PW: " + pw)

	redisPW := client.Get(ctx, slug)
	fmt.Println(redisPW)
}

func getArgumentTime() int {
	if len(os.Args) > 1 {
		arg1 := os.Args[1]

		i, _ := strconv.Atoi(arg1)
		return i
	}

	return 24
}
