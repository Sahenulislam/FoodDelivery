package main

import (
	"github.com/sahenulislam/food-delivery/controllers"
	"github.com/sahenulislam/food-delivery/models"
	"github.com/sahenulislam/food-delivery/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:pass1234@tcp(localhost:3306)/food-delivery?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Restaurant{}, &models.MenuItem{}, &models.User{}, &models.Transaction{})

	restaurantService := services.NewRestaurantService(db)
	userService := services.NewUserService(db)

	// if err := restaurantService.LoadData("data/restaurant_with_menu.json"); err != nil {
	// 	panic(err)
	// }

	// if err := userService.LoadData("data/users_with_purchase_history.json"); err != nil {
	// 	panic(err)
	// }

	restaurantController := controllers.NewRestaurantController(restaurantService)
	userController := controllers.NewUserController(userService)

	r := gin.Default()

	restaurantRoutes := r.Group("/restaurants")
	{
		restaurantRoutes.GET("/open", restaurantController.GetOpenRestaurants)
		restaurantRoutes.GET("/top", restaurantController.GetTopRestaurants)
		restaurantRoutes.GET("/search/restaurant", restaurantController.SearchRestaurants)
		restaurantRoutes.GET("/search/dishes", restaurantController.SearchDishes)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/purchase", userController.Purchase)
	}

	r.Run(":8080")
}
