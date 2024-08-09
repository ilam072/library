package main

import (
	"github.com/sirupsen/logrus"
	"islamic-library/internal/api"
	"islamic-library/internal/config"
	"islamic-library/internal/repository/books"
	"islamic-library/internal/repository/users"
	authservice "islamic-library/internal/service/auth"
	bookservice "islamic-library/internal/service/books"
	userservice "islamic-library/internal/service/users"
	"islamic-library/pkg/db"
	"net/http"
	"os"
)

var logger = logrus.New()

func main() {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("config got successfully")

	storage, err := db.NewDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("DB created successfully")

	defer func() {
		if err := storage.Close(); err != nil {
			logger.Error(err)
		}
	}()

	logger.Debugln("Up migrations...")
	if err := db.Migrate(storage); err != nil {
		logger.Fatal(err)
	}

	userRepo := users.New(storage)
	bookRepo := books.New(storage)

	r := api.New(logger, authservice.New(
		userRepo,
		logger,
	),
		userservice.New(
			userRepo,
			logger,
		),
		bookservice.New(
			userRepo,
			bookRepo,
			logger,
		),
	)

	srv := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("listen: %s\n", err)
	}
}
