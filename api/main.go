package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/RayMathew/crisis-core-materia-fusion-api/internal/database"
	"github.com/RayMathew/crisis-core-materia-fusion-api/internal/env"
	"github.com/RayMathew/crisis-core-materia-fusion-api/internal/version"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

func main() {
	err := godotenv.Load(".env") // Loads variables from .env file
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	err = run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL  string
	httpPort int
	db       struct {
		dsn string
	}
}

type application struct {
	config config
	db     *database.DB
	logger *slog.Logger
	wg     sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL")
	cfg.httpPort = env.GetInt("HTTP_PORT")
	cfg.db.dsn = env.GetString("DB_DSN")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.NewConnection(cfg.db.dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
	}

	return app.serveHTTP()
}
