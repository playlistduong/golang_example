package respond

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, message string, status int) {
	JSON(w, status, map[string]interface{}{
		"code":    status,
		"message": message,
	})
}
