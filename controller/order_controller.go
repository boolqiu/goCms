package controller

import (
	"context"
	"goCms/service"
)

type OrderController struct {
	Ctx     context.Context
	Service service.OrderService
}
