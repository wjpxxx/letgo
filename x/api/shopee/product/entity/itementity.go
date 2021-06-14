package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//ItemEntity
type ItemEntity struct{
	ItemID int64 `json:"item_id"`
	CategoryID int64 `json:"category_id"`
	ItemName string `json:"item_name"`
	Description string `json:"description"`
	ItemSku string `json:"item_sku"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
	AttributeList []AttributeEntity `json:"attribute_list"`
	PriceInfo []PriceInfoEntity `json:"price_info"`
	StockInfo []StockInfoEntity `json:"stock_info"`
	Image ImageEntity `json:"image"`
	Weight string `json:"weight"`
	Dimension DimensionEntity `json:"dimension"`
	LogisticInfo []LogisticInfoEntity `json:"logistic_info"`
	PreOrder PreOrderEntity `json:"pre_order"`
	Wholesales []WholesalesEntity `json:"wholesales"`
	Condition string `json:"condition"`
	SizeChart string `json:"size_chart"`
	ItemStatus string `json:"item_status"`
	HasModel bool `json:"has_model"`
	PromotionID int64 `json:"promotion_id"`
	VideoInfo []VideoInfoEntity `json:"video_info"`
	Brand BrandEntity `json:"brand"`
	ItemDangerous int `json:"item_dangerous"`
}

//String
func(i ItemEntity)String()string{
	return lib.ObjectToString(i)
}

//ItemEntity
type ItemExtraEntity struct{
	ItemID int64 `json:"item_id"`
	Sale int `json:"sale"`
	Views int `json:"views"`
	Likes int `json:"likes"`
	RatingStar float32 `json:"rating_star"`
	CommentCount int `json:"comment_count"`
}

//String
func(i ItemExtraEntity)String()string{
	return lib.ObjectToString(i)
}


//ItemListEntity
type ItemListEntity struct{
	ItemID int64 `json:"item_id"`
	ItemStatus string `json:"item_status"`
	UpdateTime int `json:"update_time"`
}

//String
func(i ItemListEntity)String()string{
	return lib.ObjectToString(i)
}