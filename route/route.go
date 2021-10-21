package route

import (
	"locally/goback/db"
	"locally/goback/pkg/http"
	"locally/goback/pkg/http/middleware"
	"locally/goback/pkg/repository"
	"locally/goback/pkg/usecase"

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
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
