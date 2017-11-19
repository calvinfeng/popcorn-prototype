package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"movie-gopher/model"
)

func DatabaseInit() (*gorm.DB, error) {
	args := "user=cfeng password=cfeng dbname=movie_gopher_development sslmode=disable"
	if db, err := gorm.Open("postgres", args); err != nil {
		return nil, err
	} else {
		db.AutoMigrate(&model.User{}, &model.Movie{}, &model.Rating{})
		return db, nil
	}
}
