package test

import "log"

func LogInfo(msg string) {
	log.Printf("INFO - %v", msg)
}

func LogWarning(msg string) {
	log.Printf("WARN - %v", msg)
}
