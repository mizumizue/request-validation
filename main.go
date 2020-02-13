package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	apperr "github.com/trewanek/request-validation/apperror"
	"github.com/trewanek/request-validation/model"
	"github.com/trewanek/request-validation/request"
	"github.com/trewanek/request-validation/response"
)

func main() {
	server := NewServer(echo.New())
	server.HTTPErrorHandler = ErrorHandler
	server.Route()
	server.Logger.Fatal(server.Start(":1323"))
}

func ErrorHandler(err error, echo echo.Context) {
	if errors.As(err, &apperr.ValidationErr{}) {
		errorResponse(http.StatusBadRequest, "Invalid Request", err, echo)
		return
	}
	errorResponse(http.StatusInternalServerError, "Unknown", err, echo)
}

func errorResponse(code int, status string, err error, echo echo.Context) {
	errs := make([]string, 0, 0)
	ve := func(err error) *apperr.ValidationErr {
		for err != nil {
			if validationErr, ok := err.(apperr.ValidationErr); ok {
				return &validationErr
			}
			err = errors.Unwrap(err)
			continue
		}
		return nil
	}(err)

	if ve != nil {
		errs = ve.ValidationErrs()
	} else {
		errs = append(errs, err.Error())
	}

	e := echo.JSON(code, response.HttpResponse{
		Code:   code,
		Status: status,
		Errors: errs,
	})
	if err != nil {
		echo.Logger().Errorf("err: ", e)
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
		return fmt.Errorf(
			"double wrapped err: %w",
			fmt.Errorf(
				"wrapped err: %w",
				apperr.NewValidationErr(errs),
			),
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
