package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
			adCompany.POST("/", h.newCompany)
			adCompany.GET("/", h.getCompanyList)
		}
		users := api.Group("/user")
		{
			users.GET("/", h.getUserList)
			users.GET("/:id", h.getUser)
		}
	}

	return router
}

type optionsMiddleware struct {
}

func CreateOptionsMiddleware() *optionsMiddleware {
	return &optionsMiddleware{}
}

func (middleware *optionsMiddleware) Response(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://localhost:3000") //TODO separate production and dev
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Access-Control-Allow-Origin, access-control-allow-origin")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
