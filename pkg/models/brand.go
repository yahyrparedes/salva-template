package models

import (
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model
	Name   string `gorm:"varchar(255);uniqueIndex;not null" json:"name,omitempty"`
	Active bool   `gorm:"default:false;not null" json:"active"`
}
