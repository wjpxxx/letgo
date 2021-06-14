package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//AttributeEntity
type AttributeEntity struct{
	AttributeID int64 `json:"attribute_id"`
	OriginalAttributeName string `json:"original_attribute_name"`
	IsMandatory bool `json:"is_mandatory"`
	AttributeType int `json:"attribute_type"`
	AttributeValueList []AttributeValueEntity `json:"attribute_value_list"`
}

//String
func(a AttributeEntity)String()string{
	return lib.ObjectToString(a)
}