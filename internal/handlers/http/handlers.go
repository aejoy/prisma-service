package http

import (
	"github.com/aejoy/prisma-service/internal/interfaces"
)

type PrismaHandlers struct {
	prismaService interfaces.Service
}

func NewHTTPHandlers(prismaService interfaces.Service) *PrismaHandlers {
	return &PrismaHandlers{prismaService}
}
