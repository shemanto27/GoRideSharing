package config

import (
	"log"
	"os"
)

type config struct {
	Port string
}

func Loadconfig() *config {
	port := os.Getenv("PORT")

	if port=="" {
		port = "8080"
		log.Println("PORT not set in env, using default port 8080")

	}

	return &config{
		Port: port,
	}
}