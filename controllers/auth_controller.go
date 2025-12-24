package controllers

import (
	"encoding/json"
	"inventory-management-system/services"
	"inventory-management-system/utils"
	"net/http"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	if err := services.Register(req.Username, req.Password); err != nil {
		utils.JSONResponse(w, http.StatusConflict, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, "User registered successfully", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	token, err := services.Login(req.Username, req.Password)
	if err != nil {
		utils.JSONResponse(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	utils.JSONResponse(w, http.StatusOK, "Login successful", map[string]string{"token": token})
}
