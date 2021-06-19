package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//SuccessEntity
type SuccessEntity struct{
	ItemID int64 `json:"item_id"`
	Unlist bool `json:"unlist"`
}

//String
func(d SuccessEntity)String()string{
	return lib.ObjectToString(d)
}

//BoostItemSuccessEntity
type BoostItemSuccessEntity struct{
	ItemIdList []int64 `json:"item_id_list"`
}

//String
func(d BoostItemSuccessEntity)String()string{
	return lib.ObjectToString(d)
}

//UpdatePriceSuccessEntity
type UpdatePriceSuccessEntity struct{
	ModelID int64 `json:"model_id"`
	OriginalPrice float32 `json:"original_price"`
}

//String
func(d UpdatePriceSuccessEntity)String()string{
	return lib.ObjectToString(d)
}