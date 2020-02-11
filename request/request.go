package request

import (
	"github.com/trewanek/request-validation/validator"
)

type Request interface {
	Validate() []string
}

func validate(target interface{}) []string {
	v := validator.NewValidator()
	err := v.Struct(target)
	if err == nil {
		return nil
	}
	return v.ValidationStrings(err)
}

var _ Request = (*ArticleRequest)(nil)

type ArticleRequest struct {
	Title  string         `json:"title" validate:"required"`
	Body   string         `json:"body" validate:"required"`
	Author *AuthorRequest `json:"author" validate:"required"`
}

func (req *ArticleRequest) Validate() []string {
	return validate(req)
}

var _ Request = (*CreateArticleRequest)(nil)

type CreateArticleRequest struct {
	*ArticleRequest
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
