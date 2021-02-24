// Copyright 2020-2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package render

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestJSONRender(t *testing.T) {
	data := map[string]interface{}{
		"key1": "string value",
		"key2": 2,
		"key3": true,
		"key4": []string{"value1", "value2"},
		"key5": map[string]string{"foo": "bar"},
	}
	source, _ := json.Marshal(data)

	jsonRender := &JSONRender{}
	b := bytes.NewBuffer(nil)
	err := jsonRender.Render(b, data)
	if err != nil {
		t.Errorf("an error occurred when encode data, error: %s", err)
	}
	if !bytes.Equal(b.Bytes(), source) {
		t.Error("the json render data don't equal to source data")
	}
}
