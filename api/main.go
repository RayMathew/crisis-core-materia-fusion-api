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

//	@title			Crisi Core Materia Fusion API
//	@description	A server for simulating Materia Fusion outputs in the game Crisis Core: Final Fantasy VII
//	@version	0.1

//	@contact.name	Ray Mathew
//	@contact.url	https://github.com/RayMathew/
//	@contact.email	ray10mathew@gmail.com

//	@license.name	GNU GENERAL PUBLIC LICENSE
//	@license.url	https://www.gnu.org/licenses/gpl-3.0.en.html

// @host		crisis-core-materia-fusion-api-546461677134.us-central1.run.app
// @schemes https
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
	baseURL string
	db      struct {
		dsn string
	}
	httpPort                 int
	apiTimeout               int
	apiCallsAllowedPerSecond float64
}

type application struct {
	db     *database.DB
	logger *slog.Logger
	cache  map[string]interface{}
	wg     sync.WaitGroup
	mu     sync.Mutex
	config config
}

func run(logger *slog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL")
	cfg.httpPort = env.GetInt("HTTP_PORT")
	cfg.db.dsn = env.GetString("DB_DSN")
	cfg.apiTimeout = env.GetInt("API_TIMEOUT_SECONDS")
	cfg.apiCallsAllowedPerSecond = float64(env.GetInt("API_CALLS_ALLOWED_PER_SECOND"))

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
		cache:  make(map[string]interface{}),
	}

	return app.serveHTTP()
}
