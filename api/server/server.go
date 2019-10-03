package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run(port int) {
	mux := http.NewServeMux()
	container := NewContainer()
	address := fmt.Sprintf("%s:%d", container.HttpAddr, port)

	router := Route(mux, container)

	go container.WebsocketUseCase.Start()

	srv := http.Server{
		Addr:    address,
		Handler: router,
	}

	log.Printf("Server is starting on: %s", address)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
