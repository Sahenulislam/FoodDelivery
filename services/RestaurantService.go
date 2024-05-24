package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sahenulislam/food-delivery/models"
	"gorm.io/gorm"
)

type RestaurantService struct {
	DB *gorm.DB
}

func NewRestaurantService(db *gorm.DB) *RestaurantService {
	return &RestaurantService{DB: db}
}

func (service *RestaurantService) LoadData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var restaurants []models.Restaurant
	json.Unmarshal(byteValue, &restaurants)

	for _, restaurant := range restaurants {
		service.DB.Create(&restaurant)
		for _, menuItem := range restaurant.Menu {
			menuItem.RestaurantID = restaurant.ID
			service.DB.Create(&menuItem)
		}
	}

	return nil
}

func (service *RestaurantService) GetOpenRestaurants(day string, dateTime string) ([]models.Restaurant, error) {

	var openRestaurants []models.Restaurant

	var allRestaurants []models.Restaurant
	if err := service.DB.Preload("Menu").Find(&allRestaurants).Error; err != nil {
		return nil, err
	}

	for _, restaurant := range allRestaurants {

		if isOpen := isOpenAtDateTime(restaurant.OpeningHours, day, dateTime); isOpen {
			openRestaurants = append(openRestaurants, restaurant)
		}
	}

	return openRestaurants, nil
}

func isOpenAtDateTime(openingHours string, day string, dateTime string) bool {

	schedules := strings.Split(openingHours, "/")

	for _, schedule := range schedules {
		parts := strings.Split(strings.TrimSpace(schedule), " ")
		if len(parts) < 4 {
			continue
		}
		if parts[2] == "pm" && parts[5] == "am" {
			continue
		}

		if parts[0] != day {
			continue
		}

		startTimeStr := parts[1] + " " + parts[2]
		endTimeStr := parts[4] + " " + parts[5]

		parsedStartTime, err := time.Parse("03:04 pm", startTimeStr)
		if err != nil {
			continue
		}
		parsedEndTime, err := time.Parse("03:04 pm", endTimeStr)
		if err != nil {
			continue
		}
		partsx := strings.Split(strings.TrimSpace(dateTime), " ")
		//println(partsx[1])
		if parts[2] == "am" && parts[5] == "am" && partsx[1] == "pm" {
			continue
		}
		if parts[2] == "pm" && parts[5] == "pm" && partsx[1] == "am" {
			continue
		}
		//fmt.Println(day + " " + dateTime + " " + startTimeStr + " " + endTimeStr)

		if dateTime >= parsedStartTime.Format("03:04 pm") && dateTime <= parsedEndTime.Format("03:04 pm") {
			return true
		}
	}

	return false
}

func (service *RestaurantService) GetTopRestaurants(numDishes, priceRange string, moreThan string) ([]models.Restaurant, error) {
	var topRestaurants []models.Restaurant

	isMoreThan, err := strconv.ParseBool(moreThan)
	if err != nil {
		return nil, errors.New("invalid more_than parameter")
	}

	query := service.DB.Preload("Menu")

	if numDishes != "" {
		num, err := strconv.Atoi(numDishes)
		if err != nil {
			return nil, errors.New("invalid num_dishes parameter")
		}
		if isMoreThan {
			query.Having("COUNT(menu_items.id) > ?", num)
		} else {
			query.Having("COUNT(menu_items.id) < ?", num)
		}
	}

	if priceRange != "" {

		minPrice, maxPrice, err := parsePriceRange(priceRange)
		if err != nil {
			return nil, errors.New("invalid price_range parameter")
		}

		query.Joins("JOIN menu_items ON menu_items.restaurant_id = restaurants.id")
		query.Where("menu_items.price BETWEEN ? AND ?", minPrice, maxPrice)
	}

	query.Group("restaurants.id")

	query.Order("restaurants.name").Find(&topRestaurants)

	return topRestaurants, nil
}

func parsePriceRange(priceRange string) (float64, float64, error) {
	prices := strings.Split(priceRange, "-")
	if len(prices) != 2 {
		return 0, 0, errors.New("invalid price range format")
	}
	minPrice, err := strconv.ParseFloat(strings.TrimSpace(prices[0]), 64)
	if err != nil {
		return 0, 0, errors.New("invalid min price in price range")
	}
	maxPrice, err := strconv.ParseFloat(strings.TrimSpace(prices[1]), 64)
	if err != nil {
		return 0, 0, errors.New("invalid max price in price range")
	}
	return minPrice, maxPrice, nil
}

func (service *RestaurantService) SearchRestaurants(query string) ([]models.Restaurant, error) {
	var restaurants []models.Restaurant

	if err := service.DB.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(query)+"%").Find(&restaurants).Error; err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (service *RestaurantService) SearchDishes(query string) ([]models.MenuItem, error) {
	var dishes []models.MenuItem

	if err := service.DB.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(query)+"%").Find(&dishes).Error; err != nil {
		return nil, err
	}
	return dishes, nil
}
