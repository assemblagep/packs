package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAPI_AmountWithDefaultPacks(t *testing.T) {

	tests := []struct {
		name               string
		url                string
		expectedStatusCode int
		expectedBody       string
		amount             string
	}{
		{
			name:               "amount ok",
			url:                "/amount",
			expectedStatusCode: 200,
			expectedBody:       "map[250:1 2000:1]",
			amount:             "2001",
		},
		{
			name:               "amount empty",
			url:                "/amount",
			expectedStatusCode: 200,
			expectedBody:       "map[]",
			amount:             "0",
		},
		{
			name:               "amount error",
			url:                "/amount",
			expectedStatusCode: 400,
			expectedBody:       "Bad amount parameter",
			amount:             "asd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := url.Values{}
			data.Set("amount", tt.amount)
			encodedData := data.Encode()
			req, err := http.NewRequest("POST", tt.url, strings.NewReader(encodedData))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			if err != nil {
				t.Fatal(err)
			}

			handler := http.HandlerFunc(Amount)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if status := w.Code; status != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatusCode)
			}
			if strings.TrimSpace(w.Body.String()) != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					w.Body.String(), tt.expectedBody)
			}
		})
	}
}

func TestAPI_AmountWithAdjustedPacks(t *testing.T) {
	tests := []struct {
		name               string
		urlAmount          string
		urlPacks           string
		expectedStatusCode int
		expectedBody       string
		amount             string
		packs              string
	}{
		{
			name:               "amount ok",
			urlAmount:          "/amount",
			urlPacks:           "/packs",
			expectedStatusCode: 200,
			expectedBody:       "map[4:0 13:2]",
			amount:             "20",
			packs:              "13,4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// change packs
			data := url.Values{}
			data.Set("packs", tt.packs)
			encodedData := data.Encode()
			req, err := http.NewRequest("POST", tt.urlPacks, strings.NewReader(encodedData))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			if err != nil {
				t.Fatal(err)
			}

			handler := http.HandlerFunc(Packs)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if status := w.Code; status != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatusCode)
			}
			// amount
			data = url.Values{}
			data.Set("amount", tt.amount)
			encodedData = data.Encode()
			req, err = http.NewRequest("POST", tt.urlAmount, strings.NewReader(encodedData))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			if err != nil {
				t.Fatal(err)
			}

			handler1 := http.HandlerFunc(Amount)
			w = httptest.NewRecorder()
			handler1.ServeHTTP(w, req)

			if status := w.Code; status != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatusCode)
			}
			if strings.TrimSpace(w.Body.String()) != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					w.Body.String(), tt.expectedBody)
			}

		})
	}
}
