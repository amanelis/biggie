package model

import (
	"fmt"
	"reflect"
	"strconv"
)

type Order struct {
	Price float64
	Size  float64
	Id    string
}

func ParseOrder(parts []interface{}) *Order {
	val := reflect.ValueOf(parts)

	p, _ := strconv.ParseFloat(val.Index(0).Elem().String(), 64)
	s, _ := strconv.ParseFloat(val.Index(1).Elem().String(), 64)

	return &Order{
		Price: p,
		Size:  s,
		Id:    val.Index(2).Elem().String(),
	}
}

func (o *Order) String() string {
	return fmt.Sprintf("{Order Id: %v, Price: %v, Size: %v}", o.Id, o.Price, o.Size)
}
