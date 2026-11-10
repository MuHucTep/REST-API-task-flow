package handler

import (
	"net/http"

	"github.com/MuHucTep/REST-API-task-flow/internal/pkg/response"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func Root(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"message": "TaskFlow API",
	})
}
