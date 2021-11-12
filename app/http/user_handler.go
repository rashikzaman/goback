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

type UserBasicForm struct {
	Name        string `form:"name" json:"name" binding:"required,max=255"`
	Email       string `form:"email" json:"email" binding:"required,max=255,email"`
	PhoneNumber string `form:"phone_number" json:"phone_number" binding:"required,max=255"`
	Gender      string `form:"gender" json:"gender" binding:"max=20"`
	Dob         string `form:"dob" json:"dob" binding:"max=20"`
	Photo       string `form:"photo" json:"photo"`
}

type UserAdminForm struct {
	UserBasicForm
	IsVerified *bool `form:"is_verified" json:"is_verified" binding:"required"`
	IsActive   *bool `form:"is_active" json:"is_active" binding:"required"`
	IsSeller   *bool `form:"is_seller" json:"is_seller" binding:"required"`
	ChatActive *bool `form:"chat_active" json:"chat_active" binding:"required"`
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
		var json UserAdminForm
		if err := c.ShouldBindJSON(&json); err != nil {
			error.CreateJsonFormError(c, err)
		} else {
			user := convertAdminUserToDomainUser(json)
			user, err := a.UserUseCase.StoreUser(c, user)
			if err != nil {
				error.ServerErrorResponse(c, err)
			} else {
				c.JSON(http.StatusCreated, user)
			}
		}
	}
}

func convertAdminUserToDomainUser(data UserAdminForm) *domain.User {
	user := &domain.User{
		Name:         &data.Name,
		Email:        &data.Email,
		PhoneNumber:  &data.PhoneNumber,
		Photo:        data.Photo,
		Gender:       data.Gender,
		Dob:          data.Dob,
		IsVerified:   data.IsVerified,
		IsActive:     data.IsActive,
		IsSeller:     data.IsSeller,
		IsChatActive: data.ChatActive,
	}
	return user
}
