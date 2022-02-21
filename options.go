package tins

import "time"

type ServerOptions struct {
	address           string
	network           string
	protocol          string
	timeout           time.Duration
	serializationType string
}

type ServerOption func(options *ServerOptions)

func WithAddress(address string) ServerOption {
	return func(options *ServerOptions) {
		options.address = address
	}
}

func WithNetwork(network string) ServerOption {
	return func(options *ServerOptions) {
		options.network = network
	}
}

func WithProtocol(protocol string) ServerOption {
	return func(options *ServerOptions) {
		options.protocol = protocol
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(options *ServerOptions) {
		options.timeout = timeout
	}
}

func WithSerializationType(serializationType string) ServerOption {
	return func(options *ServerOptions) {
		options.serializationType = serializationType
	}
}
