package main

import (
	"context"
	"github.com/jerevick8/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	serveMux := http.NewServeMux()
	l := log.New(os.Stdout, "Product API", log.LstdFlags)
	hh := handlers.NewHello(l)
	ah := handlers.NewAbout(l)
	ph := handlers.NewProducts(l)

	serveMux.Handle("/", hh)
	serveMux.Handle("/about", ah)
	serveMux.Handle("/products", ph)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// the following code ensures that the server is shutdown properly without disrupting any requests or functions running. it will wait for all currently running requests to complete before shutting down.
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)
	sig := <-signalChan
	l.Println("Received terminate, graceful shutdown", sig)
	shutdownCtx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := s.Shutdown(shutdownCtx)
	if err != nil {
		l.Fatal(err)
	}
}
