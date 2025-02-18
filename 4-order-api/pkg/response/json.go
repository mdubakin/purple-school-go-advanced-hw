package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func EncodeJSON(w http.ResponseWriter, content any) {
	EncodeJSONWithHeaders(w, content, http.StatusOK, nil)
}

func EncodeJSONWithHeaders(w http.ResponseWriter, content any, status int, headers map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(content); err != nil {
		log.Printf("json encode error: %v\n", err)
		return
	}
}
