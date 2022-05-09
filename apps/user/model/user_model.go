package user_model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct to describe User object.
type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type=uuid"`
	Email        string    `json:"email" validate:"required,email,lte=255"`
	Name         string    `json:"name" validate:"required,gte=5,lte=255"`
	Password 		 string    `json:"password,omitempty" validate:"lte=255"`
	Status   		 int       `json:"status" validate:"required,len=1"`
	Role     		 string    `json:"role" validate:"required,lte=25"`
}