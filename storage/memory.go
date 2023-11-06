package storage

import (
	"errors"

	"github.com/ignoxx/sl3/checkout/types"
)

type MemoryStorage struct {
	orders map[int]types.Order
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		orders: make(map[int]types.Order),
	}
}

func (s *MemoryStorage) CreateOrder(order types.Order) error {
	if _, ok := s.orders[order.ID]; ok {
		return errors.New("create order: order already exists")
	}

	s.orders[order.ID] = order

	return nil
}

func (s *MemoryStorage) GetOrder(id int) (types.Order, error) {
	order, ok := s.orders[id]
	if !ok {
		return types.Order{}, errors.New("get order: order not found")
	}

	return order, nil
}

func (s *MemoryStorage) UpdateOrder(order types.Order) error {
	if _, ok := s.orders[order.ID]; !ok {
		return errors.New("update order: order not found")
	}

	s.orders[order.ID] = order

	return nil
}

// Ensure MemoryStorage implements Storage interface
var _ Storage = (*MemoryStorage)(nil)
