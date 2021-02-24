// Copyright 2020-2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package render

import (
	"encoding/json"
	"io"
)

type JSONRender struct {
	Streaming bool
}

// Render writes a json data response
func (j *JSONRender) Render(w io.Writer, v interface{}) error {
	if j.Streaming {
		return json.NewEncoder(w).Encode(v)
	}

	result, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Write(result)
	return nil
}

// ContentType returns json content-type
func (j *JSONRender) ContentType() string {
	return "application/json"
}
