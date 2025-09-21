package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/analopesdev/duochat-service/internal/auth"
	"github.com/analopesdev/duochat-service/internal/config"
	db "github.com/analopesdev/duochat-service/internal/database"
	httpx "github.com/analopesdev/duochat-service/internal/http/router"
	"github.com/analopesdev/duochat-service/internal/room"
	"github.com/analopesdev/duochat-service/internal/user"
	"github.com/analopesdev/duochat-service/internal/ws"
)

func main() {
	config.Load()

	cfg := config.Values

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	poolConfig := db.PoolConfig{
		MaxConnections:  cfg.MaxConnections,
		MinConnections:  cfg.MinConnections,
		MaxConnLifetime: parseDuration(cfg.MaxConnLifetime),
		MaxConnIdleTime: parseDuration(cfg.MaxConnIdleTime),
	}

	pool, err := db.ConnectPool(ctx, cfg.DBURL, poolConfig)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer pool.Close()

	if err := db.Ping(ctx, pool); err != nil {
		log.Fatalf("Ping to database failed: %v", err)
	}
	log.Println("Database connected successfully")

	userRepo := user.NewRepository(pool)
	roomRepo := room.NewRepository(pool)
	authSvc := auth.NewService()

	userSvc := user.NewService(*userRepo)            // user service
	roomSvc := room.NewService(*roomRepo)            // room service
	userHandler := user.NewHandler(userSvc, authSvc) // user handlers HTTP
	roomHandler := room.NewHandler(roomSvc)          // room handlers HTTP

	srv := httpx.NewServer(":"+cfg.AppPort, httpx.RouterDeps{
		UserHandlers: userHandler,
		RoomHandlers: roomHandler,
		AuthHandler:  auth.NewHandler(authSvc),
		WsHandler:    ws.NewHandler(),
	})

	go func() {
		log.Printf("🚀 Server running on port :%s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("⏳ Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}

func parseDuration(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return time.Hour
	}
	return duration
}
