// Copyright 2020-2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package render

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	v := map[string]interface{}{
		"key1": "string value",
		"key2": 2,
		"key3": true,
		"key4": []string{"value1", "value2"},
		"key5": map[string]string{"foo": "bar"},
	}
	w := httptest.NewRecorder()
	encode, _ := json.Marshal(v)

	if err := JSON(w, http.StatusOK, v); err != nil {
		t.Errorf("an error occurred when render json data, error: %s", err)
	}

	// check status code
	if w.Code != http.StatusOK {
		t.Errorf("http status error, want %d, got %d", http.StatusOK, w.Code)
	}

	// check content type
	ct := w.Header().Get("Content-Type")
	if ct != renders.JSON.ContentType() {
		t.Errorf("http content-type header error, want %s, got %s", renders.JSON.ContentType(), ct)
	}

	// check body
	if !bytes.Equal(w.Body.Bytes(), append(encode, []byte("\n")...)) {
		t.Error("the json render data don't equal to source data")
	}
}

func TestText(t *testing.T) {
	v := "tfcloud-go"
	w := httptest.NewRecorder()

	if err := Text(w, http.StatusOK, v); err != nil {
		t.Errorf("an error occurred when render text data, error: %s", err)
	}

	// check status code
	if w.Code != http.StatusOK {
		t.Errorf("http status error, want %d, got %d", http.StatusOK, w.Code)
	}

	// check content type
	ct := w.Header().Get("Content-Type")
	if ct != renders.Text.ContentType() {
		t.Errorf("http content-type header error, want %s, got %s", renders.Text.ContentType(), ct)
	}

	// check body
	if !bytes.Equal(w.Body.Bytes(), []byte(v)) {
		t.Error("the text render data don't equal to source data")
	}
}
