package service

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/QianJiuGe/mysite/backend/internal/biz"
	"github.com/QianJiuGe/mysite/backend/internal/data"
	"github.com/QianJiuGe/mysite/backend/internal/model"
)

func (s *Service) ListMemos(c *gin.Context) {
	sess := c.MustGet("session").(*model.Session)
	var memos []model.Memo
	var err error
	if sess.Role == "admin" {
		memos, err = s.memo.ListAll(c.Request.Context())
	} else {
		memos, err = s.memo.List(c.Request.Context(), sess.UserID)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "list memos failed"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"memos": memos})
}

func (s *Service) CreateMemo(c *gin.Context) {
	sess := c.MustGet("session").(*model.Session)
	var req struct {
		Content  string `json:"content" binding:"required"`
		Priority int    `json:"priority"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid request body"))
		return
	}
	memo, err := s.memo.Create(c.Request.Context(), sess.UserID, req.Content, req.Priority)
	if err != nil {
		if errors.Is(err, biz.ErrValidation) {
			c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "content is required"))
			return
		}
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "create memo failed"))
		return
	}
	c.JSON(http.StatusCreated, memo)
}

func (s *Service) UpdateMemo(c *gin.Context) {
	sess := c.MustGet("session").(*model.Session)
	memoID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || memoID <= 0 {
		c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid memo id"))
		return
	}
	var req struct {
		Content  *string `json:"content"`
		Done     *bool   `json:"done"`
		Priority *int    `json:"priority"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid request body"))
		return
	}
	isAdmin := sess.Role == "admin"
	if err := s.memo.Update(c.Request.Context(), sess.UserID, memoID, isAdmin, req.Content, req.Done, req.Priority); err != nil {
		if errors.Is(err, data.ErrNotFound) {
			c.JSON(http.StatusNotFound, errBody("NOT_FOUND", "memo not found"))
			return
		}
		if errors.Is(err, biz.ErrAdminOnly) {
			c.JSON(http.StatusForbidden, errBody("FORBIDDEN", "not your memo"))
			return
		}
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "update memo failed"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": memoID})
}

func (s *Service) DeleteMemo(c *gin.Context) {
	sess := c.MustGet("session").(*model.Session)
	memoID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || memoID <= 0 {
		c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid memo id"))
		return
	}
	isAdmin := sess.Role == "admin"
	if err := s.memo.Delete(c.Request.Context(), sess.UserID, memoID, isAdmin); err != nil {
		if errors.Is(err, data.ErrNotFound) {
			c.JSON(http.StatusNotFound, errBody("NOT_FOUND", "memo not found"))
			return
		}
		if errors.Is(err, biz.ErrAdminOnly) {
			c.JSON(http.StatusForbidden, errBody("FORBIDDEN", "not your memo"))
			return
		}
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "delete memo failed"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": memoID})
}
