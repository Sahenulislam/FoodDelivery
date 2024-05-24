package controllers

import (
	"net/http"

	"github.com/sahenulislam/food-delivery/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}
func (controller *UserController) Purchase(c *gin.Context) {
	var request struct {
		UserID     uint `json:"userId"`
		MenuItemID uint `json:"menuItemId"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := controller.Service.Purchase(request.UserID, request.MenuItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
