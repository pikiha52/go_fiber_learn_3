package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Name     string    `validate:"required" json:"name"`
	Username string    `validate:"required" json:"username"`
	Password string    `validate:"required" json:"password"`
}
