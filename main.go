package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("hello world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	fmt.Println("Port:",portString)
}