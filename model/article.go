package model

type Article struct {
	ArticleID string `json:"articleId"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}
