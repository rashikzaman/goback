package error

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//https://seb-nyberg.medium.com/better-validation-errors-in-go-gin-88f983564a3d
type FieldError struct {
	Err validator.FieldError
}

func (q FieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + q.Err.Field() + "'")
	sb.WriteString(", condition: " + q.Err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if q.Err.Param() != "" {
		sb.WriteString(" { " + q.Err.Param() + " }")
	}

	if q.Err.Value() != nil && q.Err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.Err.Value()))
	}

	return sb.String()
}

func CreateJsonFormError(c *gin.Context, err error) {
	if _, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errors[fieldErr.Field()] = FieldError{fieldErr}.String()
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Abort()
}

func ServerErrorResponse(c *gin.Context, err error) {
	c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	c.Abort()
}
