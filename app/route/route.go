package route

import (
	"locally/goback/app/http"
	"locally/goback/app/http/middleware"
	"locally/goback/app/repository"
	"locally/goback/app/usecase"
	"locally/goback/db"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	apiGroup := r.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			userGroup := v1Group.Group("/users", middleware.AdminMiddleware())
			{
				userRepository := repository.NewPostgresUserRepository(db.GetDb())
				userUseCase := usecase.NewUserUseCase(userRepository)
				handler := http.NewUserHandler(userUseCase)
				userGroup.GET("/", handler.FetchUsers())
			}

			businessGroup := v1Group.Group("/businesses")
			{
				businessRepository := repository.NewPostgresBusinessRepository(db.GetDb())
				businessUseCase := usecase.NewBusinessUseCase(businessRepository)
				handler := http.NewBusinessHandler(businessUseCase)
				businessGroup.GET("/", handler.FetchBusinesses())
				businessGroup.POST("/", handler.CreateBusiness())
			}
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
