package models

import (
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Title string `json:"title"`
    Desc  string `json:"desc"`
}