package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//AttributeValueEntity
type AttributeValueEntity struct{
	ValueID int64 `json:"value_id"`
	OriginalValueName string `json:"original_value_name"`
	ValueUnit string `json:"value_unit"`
}

//String
func(a AttributeValueEntity)String()string{
	return lib.ObjectToString(a)
}