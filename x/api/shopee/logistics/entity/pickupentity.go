package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//PickupEntity
type PickupEntity struct{
	AddressList []AddressEntity `json:"address_list"`
}

//String
func(p PickupEntity)String()string{
	return lib.ObjectToString(p)
}