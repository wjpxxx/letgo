package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//StockInfoEntity
type StockInfoEntity struct{
	StockType int `json:"stock_type"`
	StockLocationID string `json:"stock_location_id"`
	CurrentStock int `json:"current_stock"`
	NormalStock int `json:"normal_stock"`
	ReservedStock int `json:"reserved_stock"`
}

//String
func(s StockInfoEntity)String()string{
	return lib.ObjectToString(s)
}