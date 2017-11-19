package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const Addr = ":3000"

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	db, err := DatabaseInit()
	if err != nil {
		logrus.Fatal("Failed to initialize database:", err)
		return
	}

	server := &http.Server{
		Handler:      LoadRoutes(db),
		Addr:         Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer func() {
		db.Close()
		server.Close()
	}()

	logrus.Infof("HTTP server is listening and serving on %v", Addr)
	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalf("Server failed to start: %v", err)
	}
}
