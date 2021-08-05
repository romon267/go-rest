package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/romon267/go-rest/internal/entities"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input entities.User
	// Validate request body
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {

}
