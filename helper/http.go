// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package helper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tfcloud-go/muxutil/render"
)

func ResponseJSON(w http.ResponseWriter, code int, payload interface{}) {
	_ = render.JSON(w, code, payload)
}

func ResponseErrorJSON(w http.ResponseWriter, code int, err error) {
	message := map[string]interface{}{
		"error": err.Error(),
	}
	ResponseJSON(w, code, message)
}

func ResponseText(w http.ResponseWriter, code int, payload string) {
	_ = render.Text(w, code, payload)
}

func ResponseErrorText(w http.ResponseWriter, code int, err error) {
	message := fmt.Sprintf("error: %v", err)
	ResponseText(w, code, message)
}

func ParseRequestBodyJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		err = fmt.Errorf("parse request body failed: %v", err)
		return err
	}

	return nil
}

func ParseResponseBodyJSON(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		err = fmt.Errorf("parse response body failed: %v", err)
		return err
	}

	return nil
}
