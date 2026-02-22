package handler

import (
	"context"
	"strconv"
	"github.com/gin-gonic/gin"
	"githukudenis.com/db-test-pooling/internal/service"
	"githukudenis.com/db-test-pooling/internal/store"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.service.GetAll(ctx)
	if err != nil {
		c.Error(err)
	}
	c.JSON(200, users)
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	ctx := c.Request.Context()
	var body struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	user, err := h.service.Save(ctx, &store.User{Name: body.Name})
	if err != nil {
		c.Error(err)
	}
	c.JSON(201, user)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	user, err := h.service.GetById(ctx, id)
	if err != nil {
		c.JSON(404, gin.H{"message": err})
	}
	c.JSON(200, user)
}