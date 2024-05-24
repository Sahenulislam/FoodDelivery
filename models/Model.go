package models

import (
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name         string     `json:"restaurantName"`
	CashBalance  float64    `json:"cashBalance"`
	OpeningHours string     `json:"openingHours"`
	Menu         []MenuItem `json:"menu" gorm:"foreignKey:RestaurantID"`
}

type MenuItem struct {
	gorm.Model
	RestaurantID uint    `json:"restaurant_id"`
	Name         string  `json:"dishName"`
	Price        float64 `json:"price"`
}

type User struct {
	gorm.Model
	Name            string        `json:"name"`
	CashBalance     float64       `json:"cashBalance"`
	PurchaseHistory []Transaction `json:"purchaseHistory" gorm:"foreignKey:UserID"`
}

type Transaction struct {
	gorm.Model
	UserID       uint    `json:"user_id"`
	MenuItemID   uint    `json:"menu_item_id"`
	Amount       float64 `json:"amount"`
	PurchaseDate string  `json:"purchaseDate"`
}
