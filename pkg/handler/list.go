package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/romon267/go-rest/internal/entities"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	id, err := h.getUserId(c)
	if err != nil {
		return
	}

	var input entities.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.sendResponse(c, gin.H{"id": listId}, http.StatusCreated)
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.sendResponse(c, gin.H{"lists": lists}, http.StatusOK)
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			h.sendResponse(c, gin.H{"list": ""}, http.StatusNotFound)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.sendResponse(c, gin.H{"list": list}, http.StatusOK)
}

func (h *Handler) updateListById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input entities.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.TodoList.UpdateById(userId, id, input)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			h.sendResponse(c, gin.H{"list": ""}, http.StatusNotFound)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.sendResponse(c, gin.H{"list": list}, http.StatusOK)
}

func (h *Handler) deleteListById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.DeleteById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.sendResponse(c, gin.H{}, http.StatusNoContent)
}
