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

		main := api.Group("/") //for authorized users //short info means without any preloads
		{
			main.Use(JwtAuthentication())
			//actions with personal account //for any authorized user
			main.GET("/me", h.MyAcc)                     //+personal account short info
			main.GET("/logout", h.Logout)                //+logged out or not
			main.GET("/events", h.ShowMyEvents)          //+short info about current user events
			main.PUT("/manage-event", h.ManageMyEvent)   //+updated short info about current user events
			main.PUT("/edit-profile", h.UpdateMyProfile) //+personal account short info

			user := main.Group("/user") //actions with another users' profiles
			{
				user.GET("/:userId", h.ShowUser)                     //+short user info //for any authorized user
				user.PUT("/:userId", h.UpdateUser)                   //+updated user short info //for super admin only
				user.PUT("/manage-event/:userId", h.manageUserEvent) //+updated short info about user events //for super admin only
				user.DELETE("/:userId", h.DeleteUser)                //+deleted user email/id //for super admin only
			}

			event := main.Group("/event") //short info means without any preloads
			{
				event.POST("/add", h.AddEvent) //+short info about event //admin only
				//event.GET("/:eventId/show", h.ShowEvent)                     //all info about event (users and location)
				//event.GET("/all-events", h.GetAllEvents)                     //short info about events
				//event.GET("/in-area", h.GetEventsInArea)                     //info about event without users data
				//event.GET("/in-location/:locationId", h.GetEventsInLocation) //short info about events in location
				event.PUT("/:eventId", h.UpdateEvent)    //+info about event without users data //admin only
				event.DELETE("/:eventId", h.DeleteEvent) //+deleted event id //admin only
			}

			location := main.Group("/location")
			{
				location.POST("/add", h.AddLocation) //+short info about location //admin only
				//location.GET("/:locationId", h.ShowLocation)      //all info about location (with short events data)
				//location.GET("/all", h.ShowAllLocations)          //short info about locations
				location.PUT("/:locationId", h.UpdateLocation)    //+all info about location (with short events data) //admin only
				location.DELETE("/:locationId", h.DeleteLocation) //+deleted location id //admin only
			}
		}
		public := api.Group("/public")
		{
			event := public.Group("/event")
			{
				event.GET("/:eventId/show", h.ShowEvent)
				event.GET("/all-events", h.GetAllEvents)
				event.GET("/in-area", h.GetEventsInArea)
				event.GET("/in-location/:locationId", h.GetEventsInLocation)
			}

			location := public.Group("/location")
			{
				location.GET("/:locationId", h.ShowLocation)
				location.GET("/all", h.ShowAllLocations)
			}
		}
	}
}
