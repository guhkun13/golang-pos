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
	Name         string    `json:"name" validate:"required,lte=255"`
	PasswordHash string    `json:"password_hash,omitempty" validate:"lte=255"`
	UserStatus   int       `json:"user_status" validate:"required,len=1"`
	UserRole     string    `json:"user_role" validate:"required,lte=25"`
}

// type ErrorResponse struct {
// 	FailedField string
// 	Tag         string
// 	Value       interface{}
// 	Details     string
// }

// func ValidateStruct(user User) []*ErrorResponse {
// 	log.Println("ValidateStruct: ", &user)
// 	var errors []*ErrorResponse
// 	validate := validator.New()
// 	err := validate.Struct(&user)

// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			log.Println("err : ", err)
// 			log.Println("err Param : ", err.Param())
// 			log.Println("err Value : ", err.Value())
// 			var element ErrorResponse
// 			element.FailedField = err.StructNamespace()
// 			element.Tag = err.Tag()
// 			element.Value = err.Value()
// 			element.Details = err.Namespace()
// 			errors = append(errors, &element)
// 		}
// 	}
// 	return errors
// }