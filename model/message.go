package model

// import (
// 	"fmt"
// )
//
// type Message struct {
// 	Type          string  `json:"type"`
// 	ProductId     string  `json:"product_id"`
// 	TradeId       int     `json:"trade_id,number"`
// 	OrderId       string  `json:"order_id"`
// 	Sequence      int     `json:"sequence,number"`
// 	MakerOrderId  string  `json:"maker_order_id"`
// 	TakerOrderId  string  `json:"taker_order_id"`
// 	Time          Time    `json:"time,string"`
// 	RemainingSize float64 `json:"remaining_size,string"`
// 	NewSize       float64 `json:"new_size,string"`
// 	OldSize       float64 `json:"old_size,string"`
// 	Size          float64 `json:"size,string"`
// 	Price         float64 `json:"price,string"`
// 	Side          string  `json:"side"`
// 	Reason        string  `json:"reason"`
// 	OrderType     string  `json:"order_type"`
// 	Funds         float64 `json:"funds,string"`
// 	NewFunds      float64 `json:"new_funds,string"`
// 	OldFunds      float64 `json:"old_funds,string"`
// 	Message       string  `json:"message"`
// }
//
// func (m *Message) ParsedSize() float64 {
// 	// s, _ := strconv.ParseFloat(m.Size, 64)
// 	return m.Size
// }
//
// func (m *Message) ParsedPrice() float64 {
// 	// p, _ := strconv.ParseFloat(m.Price, 64)
// 	return m.Price
// }
//
// func (m *Message) IsReceived() bool {
// 	return m.Type == "received"
// }
//
// func (m *Message) IsOpen() bool {
// 	return m.Type == "open"
// }
//
// func (m *Message) IsDone() bool {
// 	return m.Type == "done"
// }
//
// func (m *Message) IsMatch() bool {
// 	return m.Type == "match"
// }
//
// func (m *Message) IsBuy() bool {
// 	return m.Side == "buy"
// }
//
// func (m *Message) IsSell() bool {
// 	return m.Side == "sell"
// }
//
// func (m *Message) IsCanceled() bool {
// 	return m.Reason == "canceled"
// }
//
// func (m *Message) IsFilled() bool {
// 	return m.Reason == "filled"
// }
//
// func (m *Message) Order() *Order {
// 	return &Order{
// 		Id:    m.OrderId,
// 		Size:  m.ParsedSize(),
// 		Price: m.ParsedPrice(),
// 	}
// }
//
// func (m *Message) String() string {
// 	if m.Type == "received" {
// 		return fmt.Sprintf("%v: %v %v %v %v", m.Type, m.Side, m.Price, m.Size, m.OrderId)
// 	} else if m.Type == "open" {
// 		return fmt.Sprintf("%v: %v %v %v %v", m.Type, m.Side, m.Price, m.Size, m.OrderId)
// 	} else if m.Type == "done" {
// 		return fmt.Sprintf("%v: %v %v %v %v %v", m.Type, m.Side, m.Reason, m.Price, m.Size, m.OrderId)
// 	} else if m.Type == "match" {
// 		return fmt.Sprintf("%v: %v %v %v %v %v", m.Type, m.Side, m.Price, m.Size, m.MakerOrderId, m.TakerOrderId)
// 	} else if m.Type == "change" {
// 		return fmt.Sprintf("%v: %v %v %v %v", m.Type, m.Side, m.Price, m.NewSize, m.OldSize)
// 	} else {
// 		return fmt.Sprintf("unknown %v", m.Type)
// 	}
// }
