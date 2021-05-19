// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package handler

import (
	"net/http"
	"strings"

	"github.com/tfcloud-go/muxutil/render"
)

type HealthCheckHandler interface {
	CheckStatus(w http.ResponseWriter, r *http.Request)
}

type healthCheckHandler struct{}

func NewHealthCheckHandler() HealthCheckHandler {
	handler := &healthCheckHandler{}

	return handler
}

func (h *healthCheckHandler) CheckStatus(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")

	if strings.HasPrefix(accept, "application/json") {
		payload := map[string]string{
			"status": "ok",
		}
		_ = render.JSON(w, http.StatusOK, payload)
		return
	}

	_ = render.Text(w, http.StatusOK, "ok")
}
