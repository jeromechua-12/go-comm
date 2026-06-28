package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"net/http"

	"github.com/jeromechua-12/go-comm/internal/auth"
	"github.com/jeromechua-12/go-comm/internal/gateway"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	dsn := flag.String("dsn", os.Getenv("DSN"), "MySQL data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	// initialise handlers
	authRepo := auth.NewRepository(db)
	authSvc := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authSvc)

	svr := http.Server{
		Addr: ":8080",
		Handler: gateway.NewRouter(logger, authHandler),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", svr.Addr)

	err = svr.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	if len(dsn) == 0 {
		return nil, fmt.Errorf("missing dsn flag")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
