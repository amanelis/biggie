package main

import (
	"fmt"

	"github.com/amanelis/skynet/model"
	exchange "github.com/preichenberger/go-coinbase-exchange"
)

var prf = fmt.Printf
var srf = fmt.Sprintf
var orders = map[string]map[string]exchange.Message{}

const (
	OrderTypeBUY       = "buy"
	OrderTypeSELL      = "sell"
	OrderStateDone     = "done"
	OrderStateOpen     = "open"
	OrderStateReceived = "received"
)

func main() {
	c, err := LoadConfig(ConfigDefaults)
	if err != nil {
		panic(err)
	}
	e := exchange.NewClient(c.GetString(CoinbaseSecret), c.GetString(CoinbaseKey), c.GetString(CoinbasePhrase))

	a := &App{
		config: c,
		client: e,
		helper: LoadHelper(),
		logger: LoadLogger(c),
	}

	subs := a.helper.gdaxSubscribeParams("subscribe", "BTC-USD")

	orders[OrderTypeBUY] = make(map[string]exchange.Message)
	orders[OrderTypeSELL] = make(map[string]exchange.Message)

	index := 0
	for true {
		orderBook, err := model.SyncOrderBook(subs["product_id"], "2")
		if err != nil {
			panic(err)
		}
		aOrders, bOrders := orderBook.AskOrders(), orderBook.BidOrders()
		avgMinMaxAorders, avgMinMaxBorders := a.helper.avrMinMax(aOrders), a.helper.avrMinMax(bOrders)

		prf("Sequence[%d], Index[%d]\n", orderBook.Sequence, index)
		prf("\033[32mAsks\033[0m[%d] 	avg[%.6f]	min[%.6f]	max[%.6f]\n", len(aOrders), avgMinMaxAorders["avg"], avgMinMaxAorders["min"], avgMinMaxAorders["max"])
		prf("\033[31mBids\033[0m[%d]	avg[%.6f]	min[%.6f]	max[%.6f]\n", len(bOrders), avgMinMaxBorders["avg"], avgMinMaxAorders["min"], avgMinMaxAorders["max"])

		index++
		fmt.Println()
	}

	// writer := uilive.New()
	// conn := a.helper.gdaxConnectWss()
	// if err := conn.WriteJSON(subs); err != nil {
	// 	prf(err.Error())
	// }
	// defer conn.Close()
	//
	// writer.Start()
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
	// writer.Stop()
}
