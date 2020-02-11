package request

import (
	"github.com/trewanek/request-validation/validator"
	"reflect"
)

type Request interface {
	Validate() []string
}

// Validate を実装した複数の構造体を組み合わせると本体の構造体のみしか Validate が実装されない
// 再帰的に実行するように reflect を使っている
func validate(target interface{}) []string {
	errs := make([]string, 0, 0)
	v := validator.NewValidator()

	tv := reflect.ValueOf(target)
	if tv.Type().Kind() == reflect.Ptr {
		tv = tv.Elem()
	}

	for i := 0; i < tv.NumField(); i++ {
		field := tv.Field(i)
		if field.Type().Kind() == reflect.Struct {
			errs = append(errs, validate(tv.Interface())...)
		}
	}

	err := v.Struct(target)
	if err != nil {
		errs = append(errs, v.ValidationStrings(err)...)
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}

var _ Request = (*ArticleRequest)(nil)

type ArticleRequest struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

func (req *ArticleRequest) Validate() []string {
	return validate(req)
}

var _ Request = (*CreateArticleRequest)(nil)

type CreateArticleRequest struct {
	*ArticleRequest
	*AuthorRequest
}

func (req *CreateArticleRequest) Validate() []string {
	return validate(req)
}

type UpdateArticleRequest struct {
	*ArticleRequest
}

var _ Request = (*UpdateArticleRequest)(nil)

type UpdateArticleBodyRequest struct {
	Body string `json:"body" validate:"required"`
}

var _ Request = (*AuthorRequest)(nil)

type AuthorRequest struct {
	AuthorName string `json:"authorName" validate:"required"`
}

func (req *AuthorRequest) Validate() []string {
	return validate(req)
}
