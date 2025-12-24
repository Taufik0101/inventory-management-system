package controllers

import (
	"encoding/json"
	"inventory-management-system/services"
	"inventory-management-system/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CreateItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
}

type UpdateStockRequest struct {
	Amount int `json:"amount"`
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	item := services.CreateItem(req.Name, req.Description, req.Stock)
	utils.JSONResponse(w, http.StatusCreated, "Item created successfully", item)
}

func ListItemsHandler(w http.ResponseWriter, r *http.Request) {
	items := services.GetItems()
	utils.JSONResponse(w, http.StatusOK, "Success", items)
}

func GetItemDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	item, err := services.GetItemByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Success", item)
}

func UpdateStockHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var req UpdateStockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	item, err := services.UpdateStock(id, req.Amount)
	if err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Stock updated successfully", item)
}
