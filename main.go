package main

import (
	"context"
	"github.com/jerevick8/routes"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	s := &http.Server{
		Addr:         ":8080",
		Handler:      routes.Routes(),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			routes.L.Fatal(err)
		}
	}()
	// the following code ensures that the server is shutdown properly without disrupting any requests or functions running. it will wait for all currently running requests to complete before shutting down.
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)
	sig := <-signalChan
	routes.L.Println("Received terminate, graceful shutdown", sig)
	shutdownCtx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := s.Shutdown(shutdownCtx)
	if err != nil {
		routes.L.Fatal(err)
	}
}
