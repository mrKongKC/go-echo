package common

import "time"

type Customer struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
