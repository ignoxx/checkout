package storage

type Storage interface {
    CreateOrder(order Order) error
}
