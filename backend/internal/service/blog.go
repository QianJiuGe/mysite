package service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/QianJiuGe/mysite/backend/internal/biz"
)

func (s *Service) ListBlogPosts(c *gin.Context) {
	posts, err := s.blog.ListPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "list posts failed"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (s *Service) GetBlogPost(c *gin.Context) {
	slug := c.Param("slug")
	post, err := s.blog.GetPost(slug)
	if err != nil {
		if errors.Is(err, biz.ErrNotFound) {
			c.JSON(http.StatusNotFound, errBody("NOT_FOUND", "post not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "get post failed"))
		return
	}
	c.JSON(http.StatusOK, post)
}

func (s *Service) CreateBlogPost(c *gin.Context) {
	var req struct {
		Slug    string `json:"slug" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid request body"))
		return
	}
	if err := s.blog.CreatePost(req.Slug, req.Content); err != nil {
		if errors.Is(err, biz.ErrAlreadyExists) {
			c.JSON(http.StatusConflict, errBody("VALIDATION_ERROR", "post already exists"))
			return
		}
		if errors.Is(err, biz.ErrValidation) {
			c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid slug"))
			return
		}
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "create post failed"))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"slug": req.Slug})
}

func (s *Service) UpdateBlogPost(c *gin.Context) {
	slug := c.Param("slug")
	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errBody("VALIDATION_ERROR", "invalid request body"))
		return
	}
	if err := s.blog.UpdatePost(slug, req.Content); err != nil {
		if errors.Is(err, biz.ErrNotFound) {
			c.JSON(http.StatusNotFound, errBody("NOT_FOUND", "post not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "update post failed"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"slug": slug})
}

func (s *Service) DeleteBlogPost(c *gin.Context) {
	slug := c.Param("slug")
	if err := s.blog.DeletePost(slug); err != nil {
		if errors.Is(err, biz.ErrNotFound) {
			c.JSON(http.StatusNotFound, errBody("NOT_FOUND", "post not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, errBody("INTERNAL_ERROR", "delete post failed"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"slug": slug})
}
