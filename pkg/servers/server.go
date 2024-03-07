package servers

import (
	"FlowerHive/pkg/service"
	"context"
)

type Server interface {
	Run() (err error)
}

type httpServer struct {
	ctx context.Context
	src service.Service
}

func New(src service.Service) Server {
	return &httpServer{
		src: src,
	}
}

func (h *httpServer) Run() (err error) {
	return
}
