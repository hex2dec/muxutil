// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func dummyHandler(w http.ResponseWriter, r *http.Request) {}

func TestServerHeaderMiddleware(t *testing.T) {
	name := "TFCloud Go Server"
	mw := &ServerHeader{
		Name: name,
	}

	handler := mw.Middleware(http.HandlerFunc(dummyHandler))
	req := httptest.NewRequest("GET", "/server_header", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	result := w.Header().Get(headerKey)
	if result != name {
		t.Errorf("get server header failed, want %s, got %s", name, result)
	}
}
