package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	conf.AddAllowHeaders("Authorization")
	router.Use(cors.New(conf))

	h.RegisterRoutes(router)

	return router
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.SignUp)
			auth.POST("/sign-in", h.SignIn)
		}

		main := api.Group("/")
		{
			main.Use(JwtAuthentication())

			main.GET("/me", h.MyAcc)
			main.GET("/logout", h.Logout)
		}

		event := api.Group("/event")
		{
			event.POST("/add", h.AddEvent)
		}
		test := api.Group("/test")
		{
			test.GET("/my-test", h.test)
		}

	}
}
