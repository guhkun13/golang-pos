package utils

import (
	"log"

	"github.com/go-playground/validator/v10"

	user_model "github.com/guhkun13/go-pos/apps/user/model"
	product_model "github.com/guhkun13/go-pos/apps/product/model"
)

const (
	UserModel string = "user"
	ProductModel string = "product"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
	Details     string
}

func ValidateStruct(dataStruct interface{}, model string) []*ErrorResponse {
	var errors []*ErrorResponse
	
	// log.Println("ValidateStruct 1: ", dataStruct)
	
	if !IsValidModel(model) {
		element := ErrorResponse{
			FailedField: "Model name",
			Details:"Model name is not valid or unregistered",
			Tag :"Struct",
			Value :model,
		}
		errors = append(errors, &element)
		return errors
	}
	
	data := ConvertIface2Struct(dataStruct, model)
	validate := validator.New()
	err := validate.Struct(data)
	
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println("err : ", err)
			
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Value()
			element.Details = err.Namespace()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ConvertIface2Struct(data interface{}, model string) interface{}{
	// log.Println("GetModel", data, model)

	switch model {
	case UserModel:
		return data.(user_model.User)
	case ProductModel:
		return data.(product_model.Product)
	}
	
	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	
	return false
}

func GetValidModels() []string {
	return []string{UserModel, ProductModel}
}

func IsValidModel(model string) bool {
	validModels := GetValidModels()
	
	return contains(validModels, model)
}

