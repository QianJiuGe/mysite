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
		// Auth (public)
		auth := v1.Group("/auth")
		auth.POST("/register", svc.Register)
		auth.POST("/login", svc.Login)

		// Blog (public read)
		v1.GET("/blog/posts", svc.ListBlogPosts)
		v1.GET("/blog/posts/:slug", svc.GetBlogPost)

		// Authenticated routes
		authed := v1.Group("", authRequired(svc))

		// Home
		authed.GET("/home", svc.Home)

		// Memos (per-user)
		authed.GET("/memos", svc.ListMemos)
		authed.POST("/memos", svc.CreateMemo)
		authed.PUT("/memos/:id", svc.UpdateMemo)
		authed.DELETE("/memos/:id", svc.DeleteMemo)

		// Admin routes — all under /admin prefix with admin middleware
		admin := authed.Group("/admin", adminRequired())
		admin.GET("/users/pending", svc.PendingUsers)
		admin.POST("/users/:userId/approve", svc.ApproveUser)
		admin.POST("/blog/posts", svc.CreateBlogPost)
		admin.PUT("/blog/posts/:slug", svc.UpdateBlogPost)
		admin.DELETE("/blog/posts/:slug", svc.DeleteBlogPost)
	}

	return r
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
