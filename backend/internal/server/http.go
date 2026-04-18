package server

import (
	"github.com/gin-gonic/gin"

	"github.com/QianJiuGe/mysite/backend/internal/service"
)

func NewEngine(svc *service.Service) *gin.Engine {
	r := gin.Default()

	r.Use(cors())

	r.GET("/healthz", svc.Healthz)

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", svc.Register)
			auth.POST("/login", svc.Login)
		}

		admin := v1.Group("/admin")
		{
			admin.GET("/users/pending", svc.PendingUsers)
			admin.POST("/users/:userId/approve", svc.ApproveUser)
		}

		v1.GET("/home", svc.Home)
	}

	return r
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
