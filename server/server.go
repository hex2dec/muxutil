// Copyright 2021 TFCloud Co.,Ltd. All rights reserved.
// This source code is licensed under Apache-2.0 license
// that can be found in the LICENSE file.

package server

import (
	"log"
	"net"
	"net/http"
)

type server struct {
	Host string
	Port string
}

func NewServer(host, port string) *server {
	// TODO: validate host and port
	return &server{
		Host: host,
		Port: port,
	}
}

func (s *server) Run(handler http.Handler) error {
	addr := net.JoinHostPort(s.Host, s.Port)

	// TODO: supports starting https server

	log.Printf("Staring server on %s\n", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Printf("Staring server failed: %v\n", err)
		return err
	}

	// TODO: shutdowning server gracefully

	return nil
}
