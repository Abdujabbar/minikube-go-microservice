package main

import (
	"log"
	"net/http"

	"github.com/Abdujabbor/minikube-go-microservice/handlers"
)

func main() {
	log.Println(http.ListenAndServe(":8000", handlers.Router()))
}
