package demo

import (
	"encoding/json"
	"net/http"
)

func Handler(version, started string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "version": version})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page, err := RenderPage(version, started)
		if err != nil {
			http.Error(w, "could not render page", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(page)
	})
	return mux
}
