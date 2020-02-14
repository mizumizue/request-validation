package model

import "time"

type Article struct {
	ArticleID     string    `json:"articleId"`
	Title         string    `json:"title"`
	Body          string    `json:"body"`
	PublishedDate time.Time `json:"publishedDate"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
}
