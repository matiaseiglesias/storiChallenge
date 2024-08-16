package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/matiaseiglesias/storiChallenge/internal/dto"
)

func TestTransactionSummaryHandler(t *testing.T) {
	app, _ := SetupApp()

	t.Run("Request with empty body", func(t *testing.T) {
		w := httptest.NewRecorder()

		body := dto.SummaryRequest{}

		marshalled, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/transactions/summaries", bytes.NewReader(marshalled))
		req.Header.Set("Content-Type", "application/json")
		app.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Request with empty account", func(t *testing.T) {
		w := httptest.NewRecorder()

		body := dto.SummaryRequest{
			Account: "",
			Email:   "test@test.com",
		}

		marshalled, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/transactions/summaries", bytes.NewReader(marshalled))
		req.Header.Set("Content-Type", "application/json")
		app.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Request with empty email", func(t *testing.T) {
		w := httptest.NewRecorder()

		body := dto.SummaryRequest{
			Account: "test",
			Email:   "",
		}
		marshalled, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/transactions/summaries", bytes.NewReader(marshalled))
		req.Header.Set("Content-Type", "application/json")
		app.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Complete Request", func(t *testing.T) {
		w := httptest.NewRecorder()

		body := dto.SummaryRequest{
			Account: "test",
			Email:   "matiaseiglesias@yahoo.com",
		}

		marshalled, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST", "/transactions/summaries", bytes.NewReader(marshalled))
		req.Header.Set("Content-Type", "application/json")
		app.Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})
}
