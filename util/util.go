package util

import (
	"log"
	"math/rand"
	"time"
)

func RandomString(n int) string {
	log.Println("random")
	var letters = []byte("asdfghjklzxcvbnmqwertyuioZXCVVBNMASDFGHJKLQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
