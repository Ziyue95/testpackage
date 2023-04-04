package test

import (
	"log"
	"math/rand"
)

func LogInfo(msg string) {
	log.Printf("INFO - %v", msg)
}

func LogWarning(msg string) {
	log.Printf("WARN - %v", msg)
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}