package main

import (
	"time"
)

type OrderSide string

const Buy OrderSide = "BUY"
const Sell OrderSide = "SELL"

type Order struct {
	ID        int64
	Ticker    string
	Side      OrderSide
	Price     float64
	Qty       int32
	Timestamp time.Time
}
