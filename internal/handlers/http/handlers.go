package http

import (
	"github.com/aejoy/prisma-service/internal/interfaces"
)

type Handlers struct {
	prismaService interfaces.Service
}

func NewHTTPHandlers(prismaService interfaces.Service) *Handlers {
	return &Handlers{prismaService}
}
