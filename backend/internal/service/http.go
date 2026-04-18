package service

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/QianJiuGe/mysite/backend/internal/biz"
	"github.com/QianJiuGe/mysite/backend/internal/model"
)

type Service struct {
	uc   *biz.Usecase
	blog *biz.BlogUsecase
	memo *biz.MemoUsecase
}

func New(uc *biz.Usecase, blog *biz.BlogUsecase, memo *biz.MemoUsecase) *Service {
	return &Service{uc: uc, blog: blog, memo: memo}
}

func (s *Service) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid request body"))
		return
	}
	id, status, err := s.uc.Register(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, biz.ErrValidation) {
			c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid username or password"))
			return
		}
		if errors.Is(err, biz.ErrAlreadyExists) {
			c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "username already exists"))
			return
		}
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "register failed"))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"userId": id, "status": status})
}

func (s *Service) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid request body"))
		return
	}
	token, role, err := s.uc.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, biz.ErrPendingApproval):
			c.JSON(http.StatusForbidden, errBody("FORBIDDEN_PENDING_APPROVAL", "account pending approval"))
		case errors.Is(err, biz.ErrUnauthorized):
			c.JSON(http.StatusUnauthorized, errBody("UNAUTHORIZED", "invalid credentials"))
		default:
			c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "login failed"))
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "role": role})
}

func (s *Service) PendingUsers(c *gin.Context) {
	current, err := s.currentSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errBody("UNAUTHORIZED", "invalid token"))
		return
	}
	users, err := s.uc.ListPendingUsers(c.Request.Context(), current)
	if err != nil {
		if errors.Is(err, biz.ErrAdminOnly) {
			c.JSON(http.StatusForbidden, errBody("FORBIDDEN_ADMIN_ONLY", "admin only"))
			return
		}
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "list pending users failed"))
		return
	}

	type pendingUser struct {
		UserID   int64  `json:"userId"`
		Username string `json:"username"`
		Status   string `json:"status"`
	}
	resp := make([]pendingUser, 0, len(users))
	for _, u := range users {
		resp = append(resp, pendingUser{UserID: u.ID, Username: u.Username, Status: u.Status})
	}
	c.JSON(http.StatusOK, gin.H{"users": resp})
}

func (s *Service) ApproveUser(c *gin.Context) {
	current, err := s.currentSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errBody("UNAUTHORIZED", "invalid token"))
		return
	}

	userID, convErr := strconv.ParseInt(c.Param("userId"), 10, 64)
	if convErr != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid userId"))
		return
	}

	if err := s.uc.Approve(c.Request.Context(), current, userID); err != nil {
		switch {
		case errors.Is(err, biz.ErrAdminOnly):
			c.JSON(http.StatusForbidden, errBody("FORBIDDEN_ADMIN_ONLY", "admin only"))
		case errors.Is(err, biz.ErrNotFound):
			c.JSON(http.StatusNotFound, errBody("NOT_FOUND", "user not found"))
		default:
			c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "approve failed"))
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"userId": userID, "status": "approved"})
}

func (s *Service) Home(c *gin.Context) {
	current, err := s.currentSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errBody("UNAUTHORIZED", "invalid token"))
		return
	}
	if current.Status != "approved" {
		c.JSON(http.StatusForbidden, errBody("FORBIDDEN_PENDING_APPROVAL", "account pending approval"))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"welcome": "Welcome, " + current.Username,
		"modules": []string{"profile-card", "feed-placeholder", "announcement"},
	})
}

func (s *Service) Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Service) GetSession(ctx context.Context, token string) (*model.Session, error) {
	return s.uc.GetSessionByToken(ctx, token)
}

func (s *Service) currentSession(c *gin.Context) (*model.Session, error) {
	auth := strings.TrimSpace(c.GetHeader("Authorization"))
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		return nil, biz.ErrUnauthorized
	}
	token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
	if token == "" {
		return nil, biz.ErrUnauthorized
	}
	return s.uc.GetSessionByToken(c.Request.Context(), token)
}

func errBody(code, msg string) gin.H {
	return gin.H{"code": code, "message": msg}
}
