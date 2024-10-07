package handler

import (
	"Diploma-project-backend/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Handler) signUp(c *gin.Context) {
	var input model.Users

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := handler.service.Users.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) updateUser(c *gin.Context) {
	var input model.Users

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := handler.service.Users.UpdateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (handler *Handler) signIn(c *gin.Context) {
	var input signInInput

	// Логирование входящего запроса
	fmt.Printf("Received signIn request: %+v\n", input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user, err := handler.service.Users.GetUser(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Проверка типа пользователя: профессор или студент
	isProfessor := user.JobTitle != "" // Например, если есть JobTitle, значит это профессор

	// Логирование перед отправкой ответа
	fmt.Printf("Sending response: { id: %v, isProfessor: %v }\n", user.ID, isProfessor)

	// Отправка полного JSON-ответа с данными о пользователе
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          user.ID,
		"firstName":   user.FirstName,
		"lastName":    user.LastName,
		"email":       user.Email,
		"isProfessor": isProfessor,
		"jobTitle":    user.JobTitle,
		"imageAvatar": user.ImageAvatar,
	})
}

type deleteUser struct {
	ID          string `json:"id" binding:"required"`
	IsProfessor bool   `json: "IsProfessor" binding:"required"`
}

func (handler *Handler) deleteUser(c *gin.Context) {
	var user deleteUser
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := handler.service.Users.DeleteUser(user.ID, user.IsProfessor)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
