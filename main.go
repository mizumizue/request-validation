package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	apperr "github.com/trewanek/request-validation/apperror"
	"github.com/trewanek/request-validation/model"
	"github.com/trewanek/request-validation/request"
	"github.com/trewanek/request-validation/response"
	"github.com/trewanek/request-validation/validator"
)

func main() {
	server := NewServer(echo.New())
	server.HTTPErrorHandler = ErrorHandler
	server.Validator = validator.NewValidator()
	server.Route()
	server.Logger.Fatal(server.Start(":1323"))
}

func ErrorHandler(err error, echo echo.Context) {
	if errors.As(err, &apperr.ValidationErr{}) {
		v := echo.Echo().Validator.(*validator.Validator)
		errorResponse(echo, http.StatusBadRequest, "Invalid Request", v.ValidationStrings(errors.Unwrap(err))...)
		return
	}
	errorResponse(echo, http.StatusInternalServerError, "Unknown", err.Error())
}

func errorResponse(echo echo.Context, code int, status string, errStrings ...string) {
	err := echo.JSON(code, response.HttpResponse{
		Code:   code,
		Status: status,
		Errors: errStrings,
	})
	if err != nil {
		echo.Logger().Errorf("err: ", err)
	}
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
	req := new(request.SearchArticleRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}

func FindArticleHandler(c echo.Context) error {
	panic("impl")
}

func CreateArticleHandler(c echo.Context) error {
	req := new(request.CreateArticleRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return apperr.NewValidationErr(err)
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
