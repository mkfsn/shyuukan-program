package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Program struct {
	gorm.Model

	Day       uint32    `gorm:"not null" json:"day"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	Link      string    `gorm:"not null" json:"link"`
	StartedAt time.Time `gorm:"not null" json:"startedAt"`
	EndedAt   time.Time `gorm:"not null"`
}
