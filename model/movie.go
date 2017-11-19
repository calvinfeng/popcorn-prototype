package model

import "github.com/jinzhu/gorm"

// Default table name is movies
type Movie struct {
	gorm.Model
	Title      string    `gorm:"type:varchar(100)"`
	Year       int       `gorm:"type:integer"`
	IMDBRating float32   `gorm:"type:float4;column:imdb_rating"`
	IMDBID     string    `gorm:"type:varchar(100);column:imdb_id"`
	Feature    []float64 `gorm:"type:float8[]"`
	Reviewers  []User    `gorm:"many2many:ratings"`
}
