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

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateAccount(c *gin.Context) {
	account, err := h.parseAccount(c)
	if err != nil {
		return
	}

	Db.Create(&account)

	c.JSON(http.StatusOK, map[string]int{
		"id": account.Id,
	})
}

func (h *Handler) UpdateAccount(c *gin.Context) {
	id, parse_id_error := h.parseId(c)
	account, parse_account_error := h.parseAccount(c)
	account.Id = id

	if parse_id_error != nil || parse_account_error != nil {
		return
	}

	Db.Save(&account)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": account.Id,
	})
}

func (h *Handler) GetAccount(c *gin.Context) {
	id, err := h.parseId(c)
	if err != nil {
		return
	}

	var account Account
	Db.First(&account, id)
	c.JSON(http.StatusOK, gin.H{"data": account})
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	id, err := h.parseId(c)
	if err != nil {
		return
	}

	Db.Delete(&Account{}, id)

	c.String(http.StatusOK, "account deleted")
}

// help functions

func (h *Handler) parseId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return id, err
	}

	return id, err
}

func (h *Handler) parseAccount(c *gin.Context) (*Account, error) {
	var account Account

	if err := c.BindJSON(&account); err != nil {
		fmt.Printf("failed to bind account: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return &account, err
	}

	return &account, nil
}
