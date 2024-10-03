package handler

import (
	"Diploma-project-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) InitRout() *gin.Engine {
	rout := gin.New()

	users := rout.Group("/users")
	{
		users.POST("/sign-up", handler.signUp)
		users.POST("/sign-in", handler.signIn)
		users.POST("/update-user", handler.updateUser)
		users.POST("/delete-user", handler.deleteUser)
	}

	return rout
}
