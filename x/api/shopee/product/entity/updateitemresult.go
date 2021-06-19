package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)


//UpdateItemResult
type UpdateItemResult struct{
	AddItemResult
}

//String
func(g UpdateItemResult)String()string{
	return lib.ObjectToString(g)
}