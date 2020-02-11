package validator

import (
	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/locales/ja_JP"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	*validator.Validate
	ut.Translator
}

func NewValidator() *Validator {
	trans := ut.New(ja_JP.New(), ja_JP.New(), en_US.New())
	ja, _ := trans.GetTranslator("ja_JP")

	validate := validator.New()
	err := validate.RegisterTranslation("required", ja, func(ut ut.Translator) error {
		return ut.Add("required", "{0}は必須項目です", false)
	}, transFunc)

	if err != nil {
		panic("validator init is failed. detail: " + err.Error())
	}

	return &Validator{
		validate,
		ja,
	}
}

func transFunc(ut ut.Translator, fe validator.FieldError) string {
	translator, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		return fe.(error).Error()
	}
	return translator
}

func (v *Validator) ValidationStrings(err error) []string {
	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		panic("passed not validator error. arg should be validator.ValidationErrors.")
	}
	errs := make([]string, 0, len(validationErrs))
	for _, validationErr := range validationErrs {
		errs = append(errs, validationErr.Translate(v.Translator))
	}
	return errs
}
