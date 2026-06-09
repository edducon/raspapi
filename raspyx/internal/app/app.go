package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"raspyx2/config"
	_ "raspyx2/docs"
	"raspyx2/internal/handler"
	"raspyx2/internal/parser"
	"raspyx2/internal/repository"
	"raspyx2/internal/service"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Logger setup
	log, err := setupLogger(cfg)
	if err != nil {
		fmt.Println(fmt.Sprintf("error setting up logger: %v", err))
		return
	}

	log.Info(fmt.Sprintf("starting %v v%v", cfg.App.Name, cfg.App.Version), slog.String("logLevel", cfg.Log.Level))
	log.Debug("debug messages are enabled")

	// db connection
	conn, err := InitDBPool(ctx, cfg, log)
	if err != nil {
		log.Error(fmt.Sprintf("error db connection: %v", err))
		return
	}
	defer conn.Close()

	//// redis client
	//redisClient, err := cacheClient(ctx, cfg)
	//if err != nil {
	//	log.Error(fmt.Sprintf("error redis cache: %v", err))
	//	return
	//}
	//defer redisClient.Close()

	repo := repository.NewRepository(conn)
	services := service.NewService(cfg, repo)
	handlers := handler.NewHandler(log, cfg, services)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.HTTP.Port),
		Handler: handlers.InitRoutes(),
	}

	log.Info(fmt.Sprintf("starting server at :%v", cfg.HTTP.Port))

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(fmt.Sprintf("error starting server: %v", err))
			return
		}
	}()

	parser.NewParser(ctx, cfg, log, services).Parse()

	// shutdown
	<-ctx.Done()

	stop()
	log.Info("shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server forced to shutdown", slog.String("error", err.Error()))
	}

	log.Info("server stopped")
}

func InitDBPool(ctx context.Context, cfg *config.Config, log *slog.Logger) (*pgxpool.Pool, error) {
	// Parsing config
	poolConfig, err := pgxpool.ParseConfig(cfg.PG.PGURL)
	if err != nil {
		return nil, err
	}

	// Creating pool

	// Parsing timeout from config
	timeout, err := strconv.Atoi(cfg.PG.Timeout)
	if err != nil {
		return nil, err
	}

	// Ping connection
	var pool *pgxpool.Pool
	for attempt := 1; attempt <= cfg.PG.Attempts; attempt++ {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context canceled while connecting to DB: %v", ctx.Err())
		default:
			pool, err = pgxpool.NewWithConfig(ctx, poolConfig)
			if err == nil {
				err = pool.Ping(ctx)
				if err == nil {
					return pool, nil
				}
				pool.Close()
			}

			log.Info(fmt.Sprintf("failed connect to db, attempt %v", attempt))
			if attempt < cfg.PG.Attempts {
				time.Sleep(time.Duration(timeout) * time.Second)
			}
		}
	}

	return nil, fmt.Errorf("failed to connect after %d attempts: %v", cfg.PG.Attempts, err)
}

func setupLogger(cfg *config.Config) (*slog.Logger, error) {
	var log *slog.Logger
	var err error

	var handler slog.Handler
	level := getLogLevel(strings.TrimSpace(cfg.Log.Level))

	if level == nil {
		return nil, fmt.Errorf("invalid LOG_LEVEL=%v", cfg.Log.Level)
	}

	switch strings.TrimSpace(cfg.Log.Type) {
	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: *level})
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: *level})
	default:
		return nil, fmt.Errorf("invalid LOG_TYPE=%v", cfg.Log.Type)
	}

	log = slog.New(handler)
	return log, err
}

func getLogLevel(level string) *slog.Level {
	var lvl slog.Level
	switch level {
	case "debug":
		lvl = slog.LevelDebug
	case "info":
		lvl = slog.LevelInfo
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		return nil
	}
	return &lvl
}
