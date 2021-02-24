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
)

type Render struct {
	JSON *JSONRender
}

func New() *Render {
	return &Render{
		JSON: &JSONRender{
			Streaming: true,
		},
	}
}

func (r *Render) render(renderer Renderer, w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", renderer.ContentType())
	w.WriteHeader(status)
	return renderer.Render(w, data)
}

var renders = New()

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	return renders.render(renders.JSON, w, status, data)
}
