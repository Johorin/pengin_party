package dbmodel

import "time"

type User struct {
	ID        uint      `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email;uniqueIndex"`
	UID       string    `gorm:"column:uid"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}