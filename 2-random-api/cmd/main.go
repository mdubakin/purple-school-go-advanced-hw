package main

import (
	"net/http"
	"random-api/internal/random"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /random/int", random.RandomInt())

	server := http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
