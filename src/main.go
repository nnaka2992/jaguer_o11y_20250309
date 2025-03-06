package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// Create channel to listen for signals.
var (
	signalChan chan (os.Signal) = make(chan os.Signal, 1)
)


func main() {
	shutdown, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown()

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start HTTP server.
	ctx := context.Background()
	srv := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	// start the server
	go func() {
		http.Handle("/health", otelhttp.NewHandler(http.HandlerFunc(healthHandler), "health"))
		http.Handle("/user/add", otelhttp.NewHandler(http.HandlerFunc(postUserAddHandler), "postUserAddHandler"))
		http.Handle("/user", otelhttp.NewHandler(http.HandlerFunc(getUserHandler), "getUserHandler"))
		http.Handle("/item/add", otelhttp.NewHandler(http.HandlerFunc(postItemAddHandler), "postItemAddHandler"))
		http.Handle("/item", otelhttp.NewHandler(http.HandlerFunc(getItemHandler), "getItemHandler"))
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("server exited: %+v", err)
		}
	}()
	
	// Receive output from signalChan.
	sig := <-signalChan
	log.Printf("%s signal caught. Graceful Shutdown.", sig)

	// Gracefully shutdown the server by waiting on existing requests (except websockets).
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %+v", err)
	}
	log.Print("server exited")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
