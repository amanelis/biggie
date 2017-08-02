package main

import ws "github.com/gorilla/websocket"

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

func LoadHelper() Helper {
	return Helper{}
}
