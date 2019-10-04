package server

import (
	"fmt"
	"net/http"
)

func Run(port int) {
	mux := http.NewServeMux()
	container := NewContainer()
	logger := container.Logger
	address := fmt.Sprintf("%s:%d", container.HttpAddr, port)

	router := Route(mux, container)

	go container.WebsocketUseCase.Start()

	srv := http.Server{
		Addr:    address,
		Handler: router,
	}

	logger.Infof("Server is starting on: %s", address)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("Server start failed: %s", err)
	}
}
