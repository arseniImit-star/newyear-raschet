package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCountdownEndpoint(t *testing.T) {
	srv := httptest.NewServer(NewHandler())
	defer srv.Close()

	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantDays   int
		checkDays  bool
		wantErr    bool
	}{
		{name: "без параметра", query: "", wantStatus: http.StatusOK, checkDays: false},
		{name: "конкретная дата", query: "?date=2024-12-31", wantStatus: http.StatusOK, wantDays: 1, checkDays: true},
		{name: "некорректная дата", query: "?date=not-a-date", wantStatus: http.StatusBadRequest, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(srv.URL + "/api/v1/countdown" + tt.query)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Fatalf("status = %d, want %d", resp.StatusCode, tt.wantStatus)
			}

			var body struct {
				Date      string `json:"date"`
				DaysUntil int    `json:"days_until_new_year"`
				Error     string `json:"error"`
			}
			if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
				t.Fatalf("decode: %v", err)
			}

			if tt.wantErr {
				if body.Error == "" {
					t.Errorf("expected error message, got empty")
				}
				return
			}

			if tt.checkDays {
				if body.DaysUntil != tt.wantDays {
					t.Errorf("days = %d, want %d", body.DaysUntil, tt.wantDays)
				}
			} else {
				if body.DaysUntil < 1 || body.DaysUntil > 366 {
					t.Errorf("days = %d, out of range [1, 366]", body.DaysUntil)
				}
			}
		})
	}
}
