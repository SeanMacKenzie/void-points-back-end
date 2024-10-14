package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"void-backend/internal/models"
)

func (app *application) Home(writer http.ResponseWriter, request *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json: "message"`
		Version string `json: "version"`
	}{
		Status:  "active",
		Message: "Void Points!",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)
	writer.Write(out)
}

func (app *application) AllRecognitions(writer http.ResponseWriter, request *http.Request) {
	var recognitions []models.Recognition

	di1, _ := time.Parse("2006-01-02", "2024-10-14")
	di2, _ := time.Parse("2006-01-02", "2024-10-13")
	di3, _ := time.Parse("2006-01-02", "2024-10-12")

	recognitions = append(recognitions, models.Recognition{
		Id:         1,
		Amount:     25,
		Recipient:  "Doug",
		Issuer:     "Sean",
		Message:    "Keep up the good work",
		DateIssued: di1,
	}, models.Recognition{
		Id:         2,
		Amount:     25,
		Recipient:  "George",
		Issuer:     "Sean",
		Message:    "Keep up the good work",
		DateIssued: di2,
	}, models.Recognition{
		Id:         3,
		Amount:     5,
		Recipient:  "Phil",
		Issuer:     "Sean",
		Message:    "You could do better",
		DateIssued: di3,
	})

	out, err := json.Marshal(recognitions)
	if err != nil {
		fmt.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")

	writer.WriteHeader(http.StatusOK)
	writer.Write(out)
}
