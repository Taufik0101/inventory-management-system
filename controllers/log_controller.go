package controllers

import (
	"inventory-management-system/services"
	"inventory-management-system/utils"
	"net/http"
)

func ListLogsHandler(w http.ResponseWriter, r *http.Request) {
	logs := services.GetLogs()
	utils.JSONResponse(w, http.StatusOK, "Success", logs)
}
