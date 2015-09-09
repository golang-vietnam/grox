package validator

import (
	"errors"
	"gopkg.in/bluesuncorp/validator.v8"
)

var validate = validator.New(&validator.Config{TagName: "validate"})

func Validate(v interface{}) (errorList []error) {
	errs := validate.Struct(v)
	if errs == nil {
		return
	}
	for _, v := range errs.(validator.ValidationErrors) {
		switch v.Field {
		case "Username":
			switch v.Tag {
			case "required":
				errorList = append(errorList, errors.New("Username required"))
			case "max":
				errorList = append(errorList, errors.New("Username max length "+v.Param+" charater"))
			case "min":
				errorList = append(errorList, errors.New("Username min length "+v.Param+" charater"))
			}
		case "Email":
			switch v.Tag {
			case "required":
				errorList = append(errorList, errors.New("Email required"))
			case "max":
				errorList = append(errorList, errors.New("Email max length "+v.Param+" character"))
			case "min":
				errorList = append(errorList, errors.New("Email min length "+v.Param+" character"))
			case "email":
				errorList = append(errorList, errors.New("Email invalid"))
			}
		case "Name":
			switch v.Tag {
			case "required":
				errorList = append(errorList, errors.New("Name required"))
			case "max":
				errorList = append(errorList, errors.New("Name max length "+v.Param+" character"))
			case "min":
				errorList = append(errorList, errors.New("Name min length "+v.Param+" character"))
			}
		case "Password":
			switch v.Tag {
			case "required":
				errorList = append(errorList, errors.New("Password required"))
			case "max":
				errorList = append(errorList, errors.New("Password max length "+v.Param+" character"))
			case "min":
				errorList = append(errorList, errors.New("Password min length "+v.Param+" character"))
			}
		}
	}
	return
}
