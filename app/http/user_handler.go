package http

import (
	"locally/goback/app/domain"
	"net/http"

	"locally/goback/app/error"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

func NewUserHandler(us domain.UserUseCase) *UserHandler {
	handler := &UserHandler{
		UserUseCase: us,
	}
	return handler
}

func (a *UserHandler) FetchUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := a.UserUseCase.FetchUsers(c)
		if err != nil {
			error.ServerErrorResponse(c, err)
		}
		c.JSON(http.StatusOK, users)
	}
}

func (a *UserHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"hello": "world"})
	}
}
