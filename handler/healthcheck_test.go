// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	handler := &healthCheckHandler{}

	var checkStatusFunc = func(rr *httptest.ResponseRecorder) {
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %d want %d",
				status, http.StatusOK)
		}
	}

	var checkBodyFunc = func(rr *httptest.ResponseRecorder, expected []byte) {
		body := bytes.TrimRight(rr.Body.Bytes(), "\n")
		if !bytes.Equal(body, expected) {
			t.Errorf("handler returned unexpected body: got %s want %s",
				body, expected)
		}
	}

	t.Run("the default format (text format)", func(tt *testing.T) {
		req, err := http.NewRequest("GET", "/healthcheck", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler.CheckStatus(rr, req)

		checkStatusFunc(rr)

		expected := []byte("ok")
		checkBodyFunc(rr, expected)
	})

	t.Run("the json format", func(tt *testing.T) {
		req, err := http.NewRequest("GET", "/healthcheck", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Accept", "application/json")

		rr := httptest.NewRecorder()

		handler.CheckStatus(rr, req)

		checkStatusFunc(rr)

		contentType := rr.Header().Get("Content-Type")
		if !strings.HasPrefix(contentType, "application/json") {
			t.Errorf("handler returned wrong Content-Type Header, got %s want %s",
				contentType, "application/json")
		}

		expected := []byte(`{"status":"ok"}`)
		checkBodyFunc(rr, expected)
	})
}
