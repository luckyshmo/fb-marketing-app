package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/config"
	"github.com/luckyshmo/fb-marketing-app/targetted-back/internal/service"

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
		facebook := api.Group("/facebook")
		{
			facebook.GET("/pending", h.pendingFacebookPages)
			facebook.DELETE("/pending/:id", h.deletePendingFacebookPage)

			facebook.GET("/owned", h.ownedFacebookPages)
			facebook.GET("/owned/:id", h.isPageOwnedByID)

			facebook.POST("/claim/:id", h.claimPage)
		}
		adCampaign := api.Group("/campaign")
		{
			adCampaign.POST("/", h.createAdCampaign)
			adCampaign.GET("/", h.getCampaignList)
			adCampaign.GET("/:id", h.getCampaignByID)
			adCampaign.PUT("/:id", h.updateCampaign)
			adCampaign.DELETE("/:id", h.deleteCampaign)

			adCampaign.POST("/start/:id", h.startAdCampaign)
			adCampaign.POST("/stop/:id", h.stopAdCampaign)

			images := adCampaign.Group(":id/images")
			{
				images.GET("/", h.getCampaignImages)
			}
		}
		users := api.Group("/user")
		{
			users.GET("/", h.getUserList)
			users.GET("/:id", h.getUser)
			user := users.Group("/:id")
			{
				user.POST("/update-balance/:amount", h.updateBalance)

				payment := user.Group("/payment")
				{
					payment.POST("/token", h.getPaymentToken)
					payment.POST("/status/:payment-id", h.getPaymentStatus)
				}

				campaign := user.Group("/campaign")
				{
					campaign.GET("/", h.getCampaignListByUserID)
					images := campaign.Group(":campaign-id/images")
					{
						images.GET("/", h.getUserCampaignImages)
					}
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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
