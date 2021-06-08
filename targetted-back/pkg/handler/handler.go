package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/pkg/service"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	_ "github.com/luckyshmo/fb-marketing-app/targetted-back/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(CORS())

	//Open-Api endpoints documentation using GIN swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity) //JWT Auth
	{
		adCompany := api.Group("/company")
		{
			adCompany.POST("/", h.createAdCompany)
			adCompany.GET("/", h.getCompanyList)
			adCompany.GET("/:id", h.getCompanyByID)

			images := adCompany.Group(":id/images")
			{
				images.GET("/", h.getCompanyImages)
			}
			// adCompany.DELETE("/:id")
		}
		users := api.Group("/user")
		{
			users.GET("/", h.getUserList)
			users.GET("/:id", h.getUser)
		}
	}

	return router
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Get().CorsHeader) //TODO separate production and dev
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
