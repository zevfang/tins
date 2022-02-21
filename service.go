package tins

import (
	"context"
	"fmt"
	"github.com/zevfang/tins/interceptor"
)

type Service interface {
	Register(string, Handler)
	Serve(options *ServerOptions)
	Close()
	Name() string
}

type service struct {
	server      interface{}
	ctx         context.Context
	cancel      context.CancelFunc
	serviceName string
	handlers    map[string]Handler
	opts        *ServerOptions
	closing     bool
}

type Handler func(context.Context, interface{}, func(interface{}) error, []interceptor.ServerInterceptor) (interface{}, error)

func (s *service) Register(handlerName string, handler Handler) {
	if s.handlers == nil {
		s.handlers = make(map[string]Handler)
	}
	s.handlers[handlerName] = handler
}

func (s *service) Serve(options *ServerOptions) {

}

func (s *service) Close() {
	s.closing = true
	if s.cancel != nil {
		s.cancel()
	}
	fmt.Println("service closing")
}

func (s *service) Name() string {
	return s.serviceName
}
