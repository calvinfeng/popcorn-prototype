package model

import "github.com/jinzhu/gorm"

// Default table name is ratings
type Rating struct {
	gorm.Model
	User  User
	Movie Movie
	Value float32 `gorm:"type:float4"`
}
