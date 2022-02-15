package tins

import "time"

type ServerOptions struct {
	address           string
	network           string
	protocol          string
	timeout           time.Duration
	serializationType string
}
