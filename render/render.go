// Copyright 2020-2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package render

import (
	"io"
	"net/http"
)

type Renderer interface {
	Render(io.Writer, interface{}) error
	ContentType() string
}

var (
	_ Renderer = &JSONRender{}
	_ Renderer = &TextRender{}
)

type Render struct {
	JSON *JSONRender
	Text *TextRender
}

func New() *Render {
	return &Render{
		JSON: &JSONRender{
			Streaming: true,
		},
		Text: &TextRender{},
	}
}

func (r *Render) render(renderer Renderer, w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", renderer.ContentType())
	w.WriteHeader(status)
	return renderer.Render(w, v)
}

var renders = New()

// JSON renders a data as json to response
func JSON(w http.ResponseWriter, status int, v interface{}) error {
	return renders.render(renders.JSON, w, status, v)
}

// Text renders a string data to response
func Text(w http.ResponseWriter, status int, v interface{}) error {
	return renders.render(renders.Text, w, status, v)
}
