package route

import (
	"locally/goback/db"
	"locally/goback/pkg/common/middleware"
	"locally/goback/pkg/user/http"
	"locally/goback/pkg/user/repository"
	"locally/goback/pkg/user/usecase"

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
