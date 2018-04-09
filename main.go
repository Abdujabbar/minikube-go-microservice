package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Abdujabbor/minikube-go-microservice/handlers"
)

func main() {
	port := os.Getenv("PORT")
	log.Println(http.ListenAndServe(":"+port, handlers.Router()))
}
