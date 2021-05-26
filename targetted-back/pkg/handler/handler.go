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

	// router.Use(CreateOptionsMiddleware().Response)
	router.Use(CORS())
	// router.Use(cors.New(cors.Config{
	// 	AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowAllOrigins:  false,
	// 	AllowOriginFunc:  func(origin string) bool { return true },
	// 	MaxAge:           86400,
	// }))

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
		users := api.Group("/user")
		{
			users.GET("/", h.getUserList)
			users.GET("/:id", h.getUser)
		}
	}

	// router.Run()
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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://localhost:8080")
		// c.Writer.Header().Set()
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
