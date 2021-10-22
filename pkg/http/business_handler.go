package http

import (
	"locally/goback/pkg/domain"
	"net/http"

	"locally/goback/pkg/error"

	"github.com/gin-gonic/gin"
)

type BusinessHandler struct {
	BusinessUseCase domain.BusinessUseCase
}

type BusinessCreateForm struct {
	Name           string `form:"name" json:"name" binding:"required"`
	OwnerName      string `form:"owner_name" json:"owner_name" binding:"required"`
	BusinessTypeId uint   `form:"business_type_id" json:"business_type_id" binding:"required"`
	Street         string `form:"street" json:"street"`
	Area           string `form:"area" json:"area"`
	Thana          string `form:"thana" json:"thana"`
	District       string `form:"district" json:"district"`
	Division       string `form:"division" json:"division"`
	Lattitude      string `form:"lattitude" json:"lattitude"`
	Longitude      string `form:"longitude" json:"longitude"`
}

func NewBusinessHandler(bs domain.BusinessUseCase) *BusinessHandler {
	handler := &BusinessHandler{
		BusinessUseCase: bs,
	}
	return handler
}

func (a *BusinessHandler) FetchBusinesses() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := a.BusinessUseCase.FetchBusinesses(c)
		if err != nil {
			error.ServerErrorResponse(c, err)
		}
		c.JSON(200, users)
	}
}

func (a *BusinessHandler) CreateBusiness() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json BusinessCreateForm
		if err := c.ShouldBindJSON(&json); err != nil {
			error.CreateJsonFormError(c, err)
		} else {
			c.JSON(http.StatusCreated, gin.H{"success": "business created"})
		}
	}
}
