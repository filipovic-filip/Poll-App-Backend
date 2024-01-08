package main

import (
	"net/http"

	"filip.filipovic/polling-app/config"
	"filip.filipovic/polling-app/db"
	"filip.filipovic/polling-app/logging"
	"filip.filipovic/polling-app/middleware"
)

func main() {
	db.SetupDb()
	defer config.AppConfig.Client.Close()

	router := middleware.RoutePaths()

	logging.Info("Starting the server on port :8080")
	logging.Fatal(http.ListenAndServe("localhost:8080", router))
}
