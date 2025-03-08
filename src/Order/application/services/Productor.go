package services

import "api-order/src/Order/domain/entities"

type IOrderProducer interface {
	PublishOrderCreated(order entities.Order) error
}
