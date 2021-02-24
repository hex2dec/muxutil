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
	data := map[string]interface{}{
		"key1": "string value",
		"key2": 2,
		"key3": true,
		"key4": []string{"value1", "value2"},
		"key5": map[string]string{"foo": "bar"},
	}
	source, _ := json.Marshal(data)

	w := httptest.NewRecorder()
	err := JSON(w, http.StatusOK, data)
	if err != nil {
		t.Errorf("an error occurred when render data, error: %s", err)
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
	if !bytes.Equal(w.Body.Bytes(), append(source, []byte("\n")...)) {
		t.Error("the json render data don't equal to source data")
	}
}
