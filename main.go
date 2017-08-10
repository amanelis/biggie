package main

import (
	"fmt"
	"time"

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

	subs := a.helper.gdaxSubscribeParams("subscribe", "BTC-USD")
	orders := map[string]map[string]exchange.Message{}
	orders[OrderTypeBUY] = make(map[string]exchange.Message)
	orders[OrderTypeSELL] = make(map[string]exchange.Message)

	writer := uilive.New()
	writer.Start()

	index := 0

	for true {
		orderBook, err := model.DownloadOrderBook(subs["product_id"], "2")
		if err != nil {
			panic(err)
		}
		aOrders := orderBook.AskOrders()
		bOrders := orderBook.BidOrders()

		avgMinMaxAorders := a.helper.avrMinMax(aOrders)
		avgMinMaxBorders := a.helper.avrMinMax(bOrders)

		prf("Sequence[%d], Index[%d]\n", orderBook.Sequence, index)
		prf("\033[32mAsks\033[0m[%d] 	avg[%.6f]	min[%.6f]	max[%.6f]\n", len(aOrders), avgMinMaxAorders["avg"], avgMinMaxAorders["min"], avgMinMaxAorders["max"])

		prf("\033[31mBids\033[0m[%d]	avg[%.6f]	min[%.6f]	max[%.6f]\n", len(bOrders), avgMinMaxBorders["avg"], avgMinMaxAorders["min"], avgMinMaxAorders["max"])

		index++
		fmt.Println()
		time.Sleep(50 * time.Millisecond)
	}

	// conn := a.helper.gdaxConnectWss()
	// if err := conn.WriteJSON(subs); err != nil {
	// 	prf(err.Error())
	// }
	// defer conn.Close()
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
	// 	fmt.Println(message)
	//
	// 	// fmt.Fprintf(writer, "B: %d\n", len(orders["buy"]))
	// 	// fmt.Fprintf(writer, "S: %d\n", len(orders["sell"]))
	// }

	writer.Stop()
}
