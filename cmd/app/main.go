package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vk_server/configs"
	"vk_server/internal/handler"
	ad2 "vk_server/internal/handler"
	"vk_server/internal/middleware"
	"vk_server/internal/repository/ad"
	"vk_server/internal/repository/user"
)

func main() {
	storage, err := configs.NewStorage()
	if err != nil {
		log.Fatal("Failed to create storage")
		os.Exit(1)
	}
	adRepo := ad.NewRepoAd(storage)
	userRepo := user.NewUserRepo(storage)
	controllerAd := *ad2.NewControllerAd(adRepo)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	authHandler := handler.NewAuthHandler(userRepo, []byte(configs.JwtConfig().Secret))

	mux := chi.NewRouter()

	mux.Group(func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)

	})
	mux.Group(func(r chi.Router) {
		r.Use(middleware.OptionalAuthMiddleware)
		r.Get("/getAll", controllerAd.GetAll())
	})
	mux.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Post("/addAd", controllerAd.AddAd())

	})
	srv := &http.Server{
		Addr:         ":8080",
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
