package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var account Account

	if err := c.BindJSON(&account); err != nil {
		fmt.Printf("failed to bind account: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Insert(&account)

	c.JSON(http.StatusOK, map[string]int{
		"id": account.Id,
	})
}

func (h *Handler) UpdateAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var account Account

	if err := c.BindJSON(&account); err != nil {
		fmt.Printf("failed to bind account: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Update(id, account)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": account.Id,
	})
}

func (h *Handler) GetAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	account, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get account %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Delete(id)

	c.String(http.StatusOK, "account deleted")
}
