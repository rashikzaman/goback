package http

import (
	"locally/goback/pkg/domain"

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
		users := a.UserUseCase.Fetch(c)
		c.JSON(200, users)
	}
}
