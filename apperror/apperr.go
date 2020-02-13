package apperr

import "strings"

type ValidationErr struct {
	errs []string
}

func NewValidationErr(errs []string) ValidationErr {
	return ValidationErr{errs: errs}
}

func (e ValidationErr) Error() string {
	return strings.Join(e.errs, "\n")
}

func (e ValidationErr) ValidationErrs() []string {
	return e.errs
}
