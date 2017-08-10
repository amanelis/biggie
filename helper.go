package main

import (
	"github.com/amanelis/skynet/model"
	ws "github.com/gorilla/websocket"
)

type Helper struct {
}

func (h Helper) gdaxConnectWss() *ws.Conn {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial("wss://ws-feed.gdax.com", nil)
	if err != nil {
		println(err.Error())
	}

	return wsConn
}

func (h Helper) gdaxSubscribeParams(t string, p string) map[string]string {
	return map[string]string{
		"type":       t,
		"product_id": p,
	}
}

func (h Helper) avrMinMax(l []*model.Order) map[string]float64 {
	var total float64
	var min = l[0].Price
	var max = l[0].Price

	for _, value := range l {
		total += value.Price

		if value.Price > max {
			max = value.Price
		}

		if value.Price < min {
			min = value.Price
		}
	}

	avg := total / float64(len(l))

	return map[string]float64{
		"avg": avg,
		"min": min,
		"max": max,
	}
}

func LoadHelper() Helper {
	return Helper{}
}
