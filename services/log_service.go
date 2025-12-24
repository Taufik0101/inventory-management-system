package services

import (
	"inventory-management-system/models"
	"time"
)

func GetLogs() []models.ActivityLog {
	return models.Logs
}

func LogActivity(itemID int, action string, amount int, message string) {
	newLog := models.ActivityLog{
		ID:        models.NextLogID,
		ItemID:    itemID,
		Action:    action,
		Amount:    amount,
		Timestamp: time.Now(),
		Message:   message,
	}
	models.Logs = append(models.Logs, newLog)
	models.NextLogID++
}
