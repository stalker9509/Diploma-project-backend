package handler

import (
	"Diploma-project-backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *Handler) addGroup(c *gin.Context) {
	var input model.Group

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := handler.service.Group.CreateGroup(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getGroup struct {
	ID string `json:"id" binding:"required"`
}

func (handler *Handler) getGroup(c *gin.Context) {
	var input getGroup

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := handler.service.Group.GetGroup(input.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) updateGroup(c *gin.Context) {
	var input model.Group

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := handler.service.Group.UpdateGroup(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type deleteGroup struct {
	ID string `json:"id" binding:"required"`
}

func (handler *Handler) deleteGroup(c *gin.Context) {
	var input deleteGroup
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := handler.service.Group.DeleteGroup(input.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
