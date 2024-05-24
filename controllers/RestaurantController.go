package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sahenulislam/food-delivery/services"
)

type RestaurantController struct {
	Service *services.RestaurantService
}

func NewRestaurantController(service *services.RestaurantService) *RestaurantController {
	return &RestaurantController{Service: service}
}

func (controller *RestaurantController) GetOpenRestaurants(c *gin.Context) {
	dateTimeStr := c.Query("datetime")
	dateTime, err := time.Parse("2006-01-02 15:04:05", dateTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datetime format"})
		return
	}

	openRestaurants, err := controller.Service.GetOpenRestaurants(dateTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch open restaurants"})
		return
	}

	c.JSON(http.StatusOK, openRestaurants)
}

func (controller *RestaurantController) GetTopRestaurants(c *gin.Context) {

	numDishes := c.Query("number_of_dishes")
	priceRange := c.Query("price_range")
	moreThan := c.Query("more_than")

	restaurants, err := controller.Service.GetTopRestaurants(numDishes, priceRange, moreThan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get top restaurants"})
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

func (controller *RestaurantController) SearchRestaurants(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'query' is required"})
		return
	}

	restaurants, err := controller.Service.SearchRestaurants(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search restaurants"})
		return
	}

	c.JSON(http.StatusOK, restaurants)
}

func (controller *RestaurantController) SearchDishes(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'query' is required"})
		return
	}

	dishes, err := controller.Service.SearchDishes(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search dishes"})
		return
	}

	c.JSON(http.StatusOK, dishes)
}
