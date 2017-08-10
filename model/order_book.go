package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OrderBook struct {
	Sequence int64           `json:"sequence"`
	Bids     [][]interface{} `json:"bids"`
	Asks     [][]interface{} `json:"asks"`
}

func (ob *OrderBook) AskOrders() []*Order {
	orders := make([]*Order, len(ob.Asks))

	for i, s := range ob.Asks {
		orders[i] = ParseOrder(s)
	}
	return orders
}

func (ob *OrderBook) BidOrders() []*Order {
	orders := make([]*Order, len(ob.Bids))

	for i, s := range ob.Bids {
		orders[i] = ParseOrder(s)
	}
	return orders
}

func DownloadOrderBook(p string, l string) (*OrderBook, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.gdax.com/products/%s/book?level=%s", p, l))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ob := OrderBook{}
	json.Unmarshal(data, &ob)

	return &ob, nil
}
