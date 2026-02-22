package routes

import (
	"github.com/gin-gonic/gin"
	"githukudenis.com/db-test-pooling/internal/handler"
)

func RegisterUserRoutes(r *gin.Engine, h *handler.UserHandler) {
	users := r.Group("/api/v1/users")
	users.GET("", h.GetAllUsers)
	users.GET("/:id", h.GetUserById)
	users.POST("", h.RegisterUser)
}
