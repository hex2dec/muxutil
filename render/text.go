// Copyright 2020-2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package render

import "io"

type TextRender struct{}

// Render writes a text response
func (t *TextRender) Render(w io.Writer, v interface{}) error {
	w.Write([]byte(v.(string)))
	return nil
}

// ContentType returns text content-type
func (t *TextRender) ContentType() string {
	return "text/plain"
}
