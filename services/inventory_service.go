package services

import (
	"errors"
	"fmt"
	"inventory-management-system/models"
	"time"
)

func CreateItem(name, description string, stock int) models.Item {
	newItem := models.Item{
		ID:          models.NextItemID,
		Name:        name,
		Description: description,
		Stock:       stock,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	models.Inventory = append(models.Inventory, newItem)
	models.NextItemID++
	return newItem
}

func GetItems() []models.Item {
	return models.Inventory
}

func GetItemByID(id int) (models.Item, error) {
	for _, item := range models.Inventory {
		if item.ID == id {
			return item, nil
		}
	}
	return models.Item{}, errors.New("item not found")
}

func UpdateStock(id int, amount int) (models.Item, error) {
	for i, item := range models.Inventory {
		if item.ID == id {
			action := "plus"
			if amount < 0 {
				action = "minus"
			}

			if item.Stock+amount < 0 {
				return models.Item{}, errors.New("insufficient stock")
			}

			models.Inventory[i].Stock += amount
			models.Inventory[i].UpdatedAt = time.Now()

			// Log activity
			LogActivity(id, action, amount, fmt.Sprintf("Stock updated by %d", amount))

			return models.Inventory[i], nil
		}
	}
	return models.Item{}, errors.New("item not found")
}
