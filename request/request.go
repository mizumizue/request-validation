package request

import (
	"github.com/trewanek/request-validation/data_type"
)

type SearchArticleRequest struct {
	Title         *string         `query:"title" search:"title"`
	PublishedDate *data_type.Date `query:"publishedDate" search:"published_date"`
}

type ArticleRequest struct {
	Title string `json:"title" validate:"required,gte=5,lt=20"`
	Body  string `json:"body" validate:"required,lte=100"`
}

type CreateArticleRequest struct {
	*ArticleRequest
	*AuthorRequest
}

type UpdateArticleRequest struct {
	*ArticleRequest
}

type UpdateArticleBodyRequest struct {
	Body string `json:"body" validate:"required"`
}

type AuthorRequest struct {
	AuthorName   string `json:"authorName" validate:"required"`
	Masterpieces []*MasterpieceRequest
}

type MasterpieceRequest struct {
	MasterpieceTitle string `json:"masterpieceTitle" validate:"required"`
	PublishYear      int    `json:"publishYear" validate:"required"`
}
