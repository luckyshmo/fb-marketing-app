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
			adCompany.PUT("/:id", h.updateCompany)
			adCompany.DELETE("/:id", h.deleteCompany)

			adCompany.POST("/start/:id", h.startAdCompany)
			adCompany.POST("/stop/:id", h.stopAdCompany)

			//TODO payment service
			adCompany.POST("/paymentToken", h.getPaymentToken)
			adCompany.POST("/paymentStatus/:id", h.getPaymentStatus)

			images := adCompany.Group(":id/images")
			{
				images.GET("/", h.getCompanyImages)
			}
		}
		users := api.Group("/user")
		{
			users.GET("/", h.getUserList)
			users.GET("/:id", h.getUser)

			users.POST("/:id/update-balance/:amount", h.updateBalance)

			company := users.Group(":id/company")
			{
				company.GET("/", h.getCompanyListByUserID)
				images := company.Group(":company-id/images")
				{
					images.GET("/", h.getUserCompanyImages)
				}
			}
		}
	}

	return router
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Get().CorsHeader)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
