package application

import (
	"api-order/src/Order/domain/entities"
	"api-order/src/Order/domain/ports"
)

type UpdateStatusOrderUseCase struct {
	OrderRepository ports.IOrder
}

func NewUpdateStatusOrderUseCase(orderRepository ports.IOrder) *UpdateStatusOrderUseCase {
	return &UpdateStatusOrderUseCase{OrderRepository: orderRepository}
}

func (o *UpdateStatusOrderUseCase) Run(id int64, status string) (entities.Order, error) {
	updated, err := o.OrderRepository.UpdateStatus(id, status)

	if err != nil {
		return entities.Order{}, nil
	}

	return updated, nil
}