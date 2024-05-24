package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/sahenulislam/food-delivery/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (service *UserService) LoadData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var users []models.User
	json.Unmarshal(byteValue, &users)

	for _, user := range users {
		service.DB.Create(&user)
		for _, history := range user.PurchaseHistory {
			var menuItem models.MenuItem
			service.DB.Create(&models.Transaction{
				UserID:       user.ID,
				MenuItemID:   menuItem.ID,
				Amount:       history.Amount,
				PurchaseDate: history.PurchaseDate,
			})
		}
	}

	return nil
}

func (service *UserService) Purchase(userID, menuItemID uint) (*models.Transaction, error) {

	tx := service.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user models.User
	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var menuItem models.MenuItem
	if err := tx.First(&menuItem, menuItemID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if user.CashBalance < menuItem.Price {
		tx.Rollback()
		return nil, fmt.Errorf("insufficient balance")
	}

	user.CashBalance -= menuItem.Price
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	transaction := &models.Transaction{
		UserID:       userID,
		MenuItemID:   menuItemID,
		Amount:       menuItem.Price,
		PurchaseDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return transaction, nil
}
