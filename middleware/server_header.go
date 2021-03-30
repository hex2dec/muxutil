// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package middleware

import "net/http"

const headerKey = "Server"

type ServerHeader struct {
	Name string
}

func (s *ServerHeader) Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerKey, s.Name)

		h.ServeHTTP(w, r)
	})
}
