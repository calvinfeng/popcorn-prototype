package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"movie-gopher/matfac"
	"net/http"
	"time"
)

const Addr = ":3000"

func StartService() {
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

func main() {
	// StartService()
	movieMap, _ := matfac.LoadMovies("data/movies.csv")
	ratingMap, _ := matfac.LoadRatingsByUserID("data/ratings.csv")
	rec := matfac.NewRecommender(ratingMap, movieMap)

	R := rec.GetRatingMatrix()
	U, M := R.Dims()
	fmt.Printf("R matrix: %v by %v\n", U, M)
}
