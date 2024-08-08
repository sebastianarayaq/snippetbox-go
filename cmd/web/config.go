package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "4000"
	}
}
