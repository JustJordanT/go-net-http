package main

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID         uuid.UUID `gorm:"primaryKey type:uuid"`
	Username   string    `gorm:"unique"`
	Email      string    `gorm:"unique"`
	Created_at  time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
}
