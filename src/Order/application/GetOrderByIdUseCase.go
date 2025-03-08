package application

import (
	"api-order/src/Order/domain/entities"
	"api-order/src/Order/domain/ports"
)

type GetOrderByIdUseCase struct {
	OrderRepository ports.IOrder
}

func NewGetOrderByIdUseCase(orderRepository ports.IOrder) *GetOrderByIdUseCase {
	return &GetOrderByIdUseCase{OrderRepository: orderRepository}
}

func (o *GetOrderByIdUseCase) Run(id int64) (entities.Order, error) {
	order, err := o.OrderRepository.GetById(id)

	if err != nil {
		return entities.Order{}, err
	}

	return order, nil
}