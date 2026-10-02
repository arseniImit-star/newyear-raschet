package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arseniImit-star/newyear-raschet/internal/countdown"
)

type response struct {
	Date      string `json:"date"`
	DaysUntil int    `json:"days_until_new_year"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// NewHandler собирает http.Handler с маршрутами сервиса.
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/countdown", handleCountdown)
	return mux
}

func handleCountdown(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	dateStr := r.URL.Query().Get("date")

	var (
		days int
		ref  time.Time
		err  error
	)

	if dateStr == "" {
		ref = time.Now().UTC()
		days = countdown.DaysUntilNewYear(ref)
	} else {
		ref, err = time.Parse(time.DateOnly, dateStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid date, expected YYYY-MM-DD")
			return
		}
		days = countdown.DaysUntilNewYear(ref)
	}

	writeJSON(w, http.StatusOK, response{
		Date:      ref.Format(time.DateOnly),
		DaysUntil: days,
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, errorResponse{Error: msg})
}

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
