package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//	@Summary		Health check endpoint for application.
//	@Description	Perform a health check.
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response
//	@Failure		400	{string}	string	"OK"
//	@Failure		404	{string}	string	"OK"
//	@Failure		500	{string}	string	"OK"
//	@Router			/healthz [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthResponse := response{
		Message: "Success",
		Code:    http.StatusOK,
	}

	log.Println("Health check endpoint has been hit.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthResponse)
}
