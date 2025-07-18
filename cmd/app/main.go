package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vk_server/configs"
	ad2 "vk_server/internal/handler"
	"vk_server/internal/repository/ad"
)

func main() {
	storage, err := configs.NewStorage()
	if err != nil {
		log.Fatal("Failed to create storage")
		os.Exit(1)
	}
	adRepo := ad.NewRepoAd(storage)
	controllerAd := *ad2.NewControllerAd(adRepo)
	mux := chi.NewRouter()
	mux.Post("/addAd", controllerAd.New())
	mux.Get("/getAll", controllerAd.GetAll())
	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  4 * time.Second,
		WriteTimeout: 4 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	serverErr := make(chan error)
	go func() {
		log.Println("Server starting on", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("Shutting down server...")
	case err := <-serverErr:
		log.Fatal("Server error: ", err)
	}

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown error: ", err)
	}
	log.Println("Server stopped")

}
