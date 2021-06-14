package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//ImageEntity
type ImageEntity struct{
	ImageUrlList []string `json:"image_url_list"`
	ImageIdList []string `json:"image_id_list"`
}

//String
func(i ImageEntity)String()string{
	return lib.ObjectToString(i)
}