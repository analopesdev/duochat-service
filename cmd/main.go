package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/analopesdev/duochat-service/internal/config"
	"github.com/analopesdev/duochat-service/internal/db"
)

func main() {
	cfg := config.Load()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Configurar o pool de conexões
	poolConfig := db.PoolConfig{
		MaxConnections:  cfg.MaxConnections,
		MinConnections:  cfg.MinConnections,
		MaxConnLifetime: parseDuration(cfg.MaxConnLifetime),
		MaxConnIdleTime: parseDuration(cfg.MaxConnIdleTime),
	}

	// Criar o pool de conexões
	pool, err := db.ConnectPool(ctx, cfg.DBURL, poolConfig)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer pool.Close()
	log.Println("✅ Pool de conexões criado com sucesso")

	// Testar a conectividade
	if err := db.Ping(ctx, pool); err != nil {
		log.Fatalf("Ping ao banco de dados falhou: %v", err)
	}
	log.Println("✅ Conexão com banco de dados verificada")

	// Log das estatísticas do pool
	stats := db.GetStats(pool)
	log.Printf("📊 Pool Stats - Max: %d, Total: %d, Idle: %d, Acquired: %d",
		stats.MaxConns(), stats.TotalConns(), stats.IdleConns(), stats.AcquiredConns())

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "duochat up")
	})

	srv := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("🚀 Server running on port :%s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
}

// parseDuration converte uma string de duração para time.Duration
func parseDuration(durationStr string) time.Duration {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Printf("Erro ao fazer parse da duração '%s', usando padrão de 1h: %v", durationStr, err)
		return time.Hour
	}
	return duration
}
