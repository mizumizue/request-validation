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

var jaJPValidations = map[string]map[string]interface{}{
	"required": {
		"text":     "{0}は必須項目です",
		"override": false,
	},
	"gt": {
		"text":     "{0}は{1}文字より大きい数で入力してください",
		"override": false,
	},
	"lt": {
		"text":     "{0}は{1}文字未満で入力してください",
		"override": false,
	},
	"gte": {
		"text":     "{0}は{1}文字以上で入力してください",
		"override": false,
	},
	"lte": {
		"text":     "{0}は{1}文字以下で入力してください",
		"override": false,
	},
}

func NewValidator() *Validator {
	trans := ut.New(ja_JP.New(), ja_JP.New(), en_US.New())
	ja, _ := trans.GetTranslator("ja_JP")
	validate := validator.New()

	for key, dic := range jaJPValidations {
		if err := validate.RegisterTranslation(key, ja, func(ut ut.Translator) error {
			return ut.Add(key, dic["text"].(string), dic["override"].(bool))
		}, transFunc); err != nil {
			panic("RegisterTranslation is failed. please check your dictionary. validator init failed... detail: " + err.Error())
		}
	}

	return &Validator{
		validate,
		ja,
	}
}

func transFunc(ut ut.Translator, fe validator.FieldError) string {
	translator, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
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
