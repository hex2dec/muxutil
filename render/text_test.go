// Copyright 2020-2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package render

import (
	"bytes"
	"testing"
)

func TestTextRender(t *testing.T) {
	v := "tfcloud-go"
	b := bytes.NewBuffer(nil)

	textRender := TextRender{}
	err := textRender.Render(b, v)
	if err != nil {
		t.Errorf("an error occurred when text data render, error: %s", err)
	}

	if !bytes.Equal(b.Bytes(), []byte(v)) {
		t.Errorf("the text render data don't equal to source data")
	}
}
