package main

type OrderBook struct {
	buys     []*Order
	sells    []*Order
	orderMap map[int64]*Order
	nextID   int64
}

func CreateEmptyOB() *OrderBook {
	ob := OrderBook{
		buys:     make([]*Order, 0),
		sells:    make([]*Order, 0),
		orderMap: make(map[int64]*Order, 1),
		nextID:   1,
	}

	return &ob
}

func (*OrderBook) PlaceOrder(ticker string, side OrderSide, price float64, quantity int32) *Order {
	return nil
}

func (*OrderBook) CancelOrder(ID int32) bool {
	return false
}

func (*OrderBook) MatchBuy(order *Order) {

}

func (*OrderBook) MatchSell(order *Order) {

}

func InsertSorted(book []*Order, O *Order, bettter func(order1, order2 *Order) []*Order) {

}

// true if order 1 comes before order 2
func IsBetterBid(order1, order2 *Order) bool {
	return false
}

// true if order 1 comes before order 2
func IsBetterAsk(order1, order2 *Order) bool {
	return false
}
