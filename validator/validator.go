package validator

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

type ValidatorInterface interface {
	ValidateRequest(s interface{}) (err error)
	ValidateRequestWithGetBody(c *gin.Context, s interface{}) (err error)
}

func CreateValidator(validator *validator.Validate) ValidatorInterface {
	return &Validator{validator}
}

func (v *Validator) ValidateRequestWithGetBody(c *gin.Context, s interface{}) (err error) {
	if err = c.ShouldBindJSON(&s); err != nil {
		return
	}
	return v.ValidateRequest(s)
}

func (v *Validator) ValidateRequest(s interface{}) (err error) {
	err = v.Validator.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return fmt.Errorf(err.Error())
		}

		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf(err.Field() + " is " + err.Tag())
		}
	}
	return
}
