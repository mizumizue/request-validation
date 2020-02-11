package main

import (
	"github.com/labstack/echo/v4"
	"github.com/trewanek/request-validation/model"
	"github.com/trewanek/request-validation/request"
	"github.com/trewanek/request-validation/response"
	"net/http"
)

func main() {
	server := NewServer(echo.New())
	server.Route()
	server.Logger.Fatal(server.Start(":1323"))
}

type Server struct {
	*echo.Echo
}

func NewServer(echo *echo.Echo) *Server {
	return &Server{Echo: echo}
}

func (server *Server) Route() {
	server.GET("/articles", FetchArticlesHandler)
	server.GET("/articles/:article_id", FindArticleHandler)
	server.POST("/articles", CreateArticleHandler)
	server.PUT("/articles/:article_id", UpdateAllFieldsArticleHandler)
	server.PATCH("/articles/:article_id", UpdateArticleHandler)
	server.DELETE("/articles/:article_id", DeleteArticleHandler)
}

func FetchArticlesHandler(c echo.Context) error {
	panic("impl")
}

func FindArticleHandler(c echo.Context) error {
	panic("impl")
}

func CreateArticleHandler(c echo.Context) error {
	req := new(request.CreateArticleRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}

	errs := req.Validate()
	if errs != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.HttpResponse{
				Code:     http.StatusBadRequest,
				Status:   "Invalid Request",
				Response: nil,
				Errors:   errs,
			},
		)
	}
	return c.JSON(
		http.StatusOK,
		response.HttpResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Response: model.Article{
				ArticleID: "article1",
				Title:     req.Title,
				Body:      req.Body,
			},
		})
}

func UpdateAllFieldsArticleHandler(c echo.Context) error {
	panic("impl")
}

func UpdateArticleHandler(c echo.Context) error {
	panic("impl")
}

func DeleteArticleHandler(c echo.Context) error {
	panic("impl")
}
