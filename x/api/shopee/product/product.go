package product

import (
	shopeeConfig "github.com/wjpxxx/letgo/x/api/shopee/config"
	"github.com/wjpxxx/letgo/x/api/shopee/product/entity"
	"github.com/wjpxxx/letgo/lib"
	"strings"
)

const (
	INT_TYPE AttributeType="INT_TYPE"
	STRING_TYPE AttributeType="STRING_TYPE"
	ENUM_TYPE AttributeType="ENUM_TYPE"
	FLOAT_TYPE AttributeType="FLOAT_TYPE"
	DATE_TYPE AttributeType="DATE_TYPE"
	TIMESTAMP_TYPE AttributeType="TIMESTAMP_TYPE"
	NORMAL ItemStatus="NORMAL"
	BANNED ItemStatus="BANNED"
	DELETED ItemStatus="DELETED"
	UNLIST ItemStatus="UNLIST"
)
//AttributeType
type AttributeType string

type ItemStatus string

//Product
type Product struct{
	Config *shopeeConfig.Config
}

//GetComment
//@Title Use this api to get comment by shop_id, item_id, or comment_id.
//@Description https://open.shopee.com/documents?module=89&type=1&id=562&version=2
func (p *Product)GetComment(itemID,commentID int64,cursor string,pageSize int)entity.GetCommentResult{
	method:="product/get_comment"
	params:=lib.InRow{
		"item_id":itemID,
		"comment_id":commentID,
		"cursor":cursor,
		"page_size":pageSize,
	}
	result:=entity.GetCommentResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//ReplyComment
//@Title Use this api to reply comments from buyers in batch.
//@Description https://open.shopee.com/documents?module=89&type=1&id=563&version=2
func (p *Product)ReplyComment(commentList []entity.ReplyCommentRequestCommentEntity)entity.ReplyCommentResult{
	method:="product/reply_comment"
	result:=entity.ReplyCommentResult{}
	err:=p.Config.HttpPost(method,commentList,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetItemBaseInfo
//@Title Use this api to get basic info of item by item_id list.
//@Description https://open.shopee.com/documents?module=89&type=1&id=612&version=2
func (p *Product)GetItemBaseInfo(itemIdList []int64)entity.GetItemBaseInfoResult{
	method:="product/get_item_base_info"
	params:=lib.InRow{
		"item_id_list":strings.Join(lib.Int64ArrayToArrayString(itemIdList),","),
	}
	result:=entity.GetItemBaseInfoResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetItemExtraInfo
//@Title Use this api to get extra info of item by item_id list.
//@Description https://open.shopee.com/documents?module=89&type=1&id=613&version=2
func (p *Product)GetItemExtraInfo(itemIdList []int64)entity.GetItemExtraInfoResult{
	method:="product/get_item_extra_info"
	params:=lib.InRow{
		"item_id_list":strings.Join(lib.Int64ArrayToArrayString(itemIdList),","),
	}
	result:=entity.GetItemExtraInfoResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetItemList
//@Title Use this call to get a list of items.
//@Description https://open.shopee.com/documents?module=89&type=1&id=614&version=2
func (p *Product)GetItemList(offset,pageSize,updateTimeFrom,updateTimeTo int,itemStatus ItemStatus)entity.GetItemListResult{
	method:="product/get_item_list"
	params:=lib.InRow{
		"offset":offset,
		"page_size":pageSize,
		"update_time_from":updateTimeFrom,
		"update_time_to":updateTimeTo,
		"item_status":itemStatus,
	}
	result:=entity.GetItemListResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//DeleteItem
//@Title Use this call to delete a product item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=615&version=2
func (p *Product)DeleteItem(itemID int64)entity.DeleteItemResult{
	method:="product/delete_item"
	params:=lib.InRow{
		"item_id":itemID,
	}
	result:=entity.DeleteItemResult{}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//AddItem
//@Title Add a new item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=616&version=2
func (p *Product)AddItem(item entity.AddItemRequestItemEntity)entity.AddItemResult{
	method:="product/add_item"
	result:=entity.AddItemResult{}
	err:=p.Config.HttpPost(method,item,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}