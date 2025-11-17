package server

// Здесь буит обработчики HTPS запросов

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	log.Printf("[%s] %s %s from %s", time.Now().Format("15:04:05"), r.Method, r.URL.Path, r.RemoteAddr)

	fmt.Fprintf(w, "Салем! Это защищённый HTTPS-сервер с автоматическим ратотором сертификатов от Let's Encrypt.\n")
	fmt.Fprintf(w, "Время на сервере: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(response)
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"message": "Это JSON API энд поинт",
		"version": "1.0.0",
		"secure":  true,
	}

	json.NewEncoder(w).Encode(data)
}
