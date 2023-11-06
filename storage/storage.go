package storage

import "github.com/ignoxx/sl3/checkout/types"

type Storage interface {
	CreateOrder(order types.Order) error
	GetOrder(id int) (types.Order, error)
	UpdateOrder(order types.Order) error
}
