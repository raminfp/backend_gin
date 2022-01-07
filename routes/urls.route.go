package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/raminfp/backend_gin/api/v1"
	"github.com/raminfp/backend_gin/db"
	"github.com/raminfp/backend_gin/middleware"
	"github.com/raminfp/backend_gin/repository"
	"github.com/raminfp/backend_gin/services"
	"gorm.io/gorm"
)

var (
	postDb         *gorm.DB                  = db.ConnectPostgres()
	authRepository repository.AuthRepository = repository.NewAuthRepository(postDb)
	authService    services.AuthService      = services.NewAuthService(authRepository)
	authAPI        v1.AuthAPI                = v1.NewAuthAPI(authService)
)

func Urls() *gin.Engine {

	//defer db.ClosePostgres(postDb)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.NoRoute(middleware.NoRouteHandler())
	r.HandleMethodNotAllowed = true
	r.NoMethod(middleware.NoMethodHandler())

	apiV1 := r.Group("api/v1")
	{
		auth := apiV1.Group("auth")
		{
			// /api/v1/auth/register
			auth.POST("/register", authAPI.Register)
			// /api/v1/auth/login
			auth.POST("/login", authAPI.Login)
			// /api/v1/auth/verify
			//auth.GET("/login", v1.EmailVerify)
			// /api/v1/auth/forgetPass
			//auth.GET("/login", v1.Forget)
		}
		user := apiV1.Group("user")
		{
			// api/v1/user/me
			// api/v1/user/profile
			// api/v1/user/tickets
			//user.POST("/home", v1.Home)
			user.POST("/login", v1.Login)
			//user.POST("/home", v1.Me)
		}

	}
	return r
}
