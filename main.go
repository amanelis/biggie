package main

import (
	"fmt"

	"github.com/amanelis/skynet/model"
	"github.com/gosuri/uilive"

	exchange "github.com/preichenberger/go-coinbase-exchange"
)

var prf = fmt.Printf

const (
	OrderTypeBUY       = "buy"
	OrderTypeSELL      = "sell"
	OrderStateDone     = "done"
	OrderStateOpen     = "open"
	OrderStateReceived = "received"
)

func main() {
	c := LoadConfig(ConfigDefaults)
	x := exchange.NewClient(c.GetString(CoinbaseSecret), c.GetString(CoinbaseKey), c.GetString(CoinbasePhrase))
	h := LoadHelper()

	a := &App{
		config: c,
		client: x,
		helper: h,
	}

	conn := a.helper.gdaxConnectWss()
	subs := a.helper.gdaxSubscribeParams("subscribe", "ETH-USD")
	if err := conn.WriteJSON(subs); err != nil {
		prf(err.Error())
	}
	defer conn.Close()

	orders := map[string]map[string]exchange.Message{}
	orders[OrderTypeBUY] = make(map[string]exchange.Message)
	orders[OrderTypeSELL] = make(map[string]exchange.Message)

	writer := uilive.New()
	writer.Start()

	orderBook, err := model.DownloadOrderBook(subs["product_id"], "3")
	if err != nil {
		panic(err)
	}

	fmt.Println(orderBook.AskOrders()[0])
	fmt.Println(orderBook.AskOrders()[1])
	fmt.Println(orderBook.AskOrders()[2])
	fmt.Println(orderBook.AskOrders()[3])
	// fmt.Println(reflect.TypeOf(orderBook))
	// fmt.Println(orderBook.Sequence)
	// fmt.Println(reflect.TypeOf(orderBook.AskOrders()))

	// message := exchange.Message{}
	// for true {
	// 	if err := conn.ReadJSON(&message); err != nil {
	// 		println(err.Error())
	// 		break
	// 	}
	//
	// 	if message.Side == OrderTypeBUY || message.Side == OrderTypeSELL {
	// 		if message.Type == OrderStateOpen {
	// 			_, exists := orders[message.Side][message.OrderId]
	//
	// 			if !exists {
	// 				orders[message.Side][message.OrderId] = message
	// 			}
	// 		}
	//
	// 		if message.Type == OrderStateDone {
	// 			_, exists := orders[message.Side][message.OrderId]
	//
	// 			if exists {
	// 				delete(orders[message.Side], message.OrderId)
	// 			}
	// 		}
	// 	}
	//
	// 	fmt.Fprintf(writer, "B: %d\n", len(orders["buy"]))
	// 	fmt.Fprintf(writer, "S: %d\n", len(orders["sell"]))
	// }

	// if message.Side == "buy" {
	// 	fmt.Printf("[%d][\033[32m%s\033[0m][%s]			%f			%f\n", message.Sequence, message.Side, message.Type, message.Size, message.Price)
	// } else {
	// 	fmt.Printf("[%d][\033[31m%s\033[0m][%s]			%f			%f\n", message.Sequence, message.Side, message.Type, message.Size, message.Price)
	// }

	// if message.Type == "done" || message.Type == "open" {
	// 	fmt.Printf("\033[32m%s\033[0m[%s]	TradeId[%d], OrderId[%s]\n", message.Side, message.Type, message.TradeId, message.OrderId)
	// }

	writer.Stop()
}
