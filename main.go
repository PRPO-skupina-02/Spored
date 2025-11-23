package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/orgs/PRPO-skupina-02/Spored/api"
	"github.com/orgs/PRPO-skupina-02/Spored/database"
)

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	slog.Info("Starting server")

	db, err := database.OpenAndMigrate()
	if err != nil {
		return err
	}

	router := gin.Default()
	api.Register(router, db)

	slog.Info("Server startup complete")
	err = router.Run(":8080")
	if err != nil {
		return err
	}

	return nil
}
