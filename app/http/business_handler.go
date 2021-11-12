package http

import (
	"locally/goback/app/domain"
	"net/http"

	"locally/goback/app/error"

	"github.com/gin-gonic/gin"
)

type BusinessHandler struct {
	BusinessUseCase domain.BusinessUseCase
}

type BusinessCreateForm struct {
	Name           string  `form:"name" json:"name" binding:"required,max=255"`
	OwnerName      string  `form:"owner_name" json:"owner_name" binding:"required,max=255"`
	BusinessTypeId uint    `form:"business_type_id" json:"business_type_id" binding:"required"`
	Street         string  `form:"street" json:"street" binding:"omitempty,max=255"`
	Area           string  `form:"area" json:"area" binding:"omitempty,max=255"`
	Thana          string  `form:"thana" json:"thana" binding:"omitempty,max=255"`
	District       string  `form:"district" json:"district" binding:"omitempty,max=255"`
	Division       string  `form:"division" json:"division" binding:"omitempty,max=255"`
	Latitude       float32 `form:"latitude" json:"latitude" binding:"omitempty,latitude"`
	Longitude      float32 `form:"longitude" json:"longitude" binding:"omitempty,longitude"`
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
			business := convertBusinessFormToDomainBusiness(json)
			storedData, err := a.BusinessUseCase.StoreBusiness(c, business)
			if err != nil {
				error.ServerErrorResponse(c, err)
			} else {
				c.JSON(http.StatusCreated, storedData)
			}

		}
	}
}

func convertBusinessFormToDomainBusiness(form BusinessCreateForm) *domain.Business {
	business := domain.Business{
		Name:           &form.Name,
		OwnerName:      &form.OwnerName,
		BusinessTypeId: &form.BusinessTypeId,
		Street:         form.Street,
		Area:           form.Area,
		Thana:          form.Thana,
		District:       form.District,
		Latitude:       form.Latitude,
		Longitude:      form.Longitude,
	}
	return &business
}
