package types

import (
	"time"
)

type OrderStatus string
type PaymentStatus string
type CryptoCurrency string

const (
	OrderStatusCreated  OrderStatus = "created"
	OrderStatusPending  OrderStatus = "pending"
	OrderStatusComplete OrderStatus = "complete"
	OrderStatusFailed   OrderStatus = "failed"

	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusComplete PaymentStatus = "complete"
	PaymentStatusFailed   PaymentStatus = "failed"

	CryptoCurrencyETH CryptoCurrency = "eth"
	CryptoCurrencyBTC CryptoCurrency = "btc"
)

type Order struct {
	ID            int
	UserID        int
	Status        OrderStatus
	StatusDetails string
	Notes         string
	Payment       CryptoPayment
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CryptoPayment struct {
	ID              int
	Status          PaymentStatus
	StatusDetails   string
	Amount          float64
	AmountReceived  float64
	ReceiverAddress string
	SenderAddress   string
	TransactionHash string
	Currency        CryptoCurrency
}

