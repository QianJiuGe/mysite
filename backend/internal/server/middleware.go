package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/QianJiuGe/mysite/backend/internal/model"
	"github.com/QianJiuGe/mysite/backend/internal/service"
)

func authRequired(svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.TrimSpace(c.GetHeader("Authorization"))
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "UNAUTHORIZED", "message": "missing token"})
			return
		}
		token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "UNAUTHORIZED", "message": "empty token"})
			return
		}
		sess, err := svc.GetSession(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "UNAUTHORIZED", "message": "invalid token"})
			return
		}
		c.Set("session", sess)
		c.Next()
	}
}

func adminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessVal, exists := c.Get("session")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "UNAUTHORIZED", "message": "no session"})
			return
		}
		sess, ok := sessVal.(*model.Session)
		if !ok || sess.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"code": "FORBIDDEN_ADMIN_ONLY", "message": "admin only"})
			return
		}
		c.Next()
	}
}
