package logistics

import (
	shopeeConfig "github.com/wjpxxx/letgo/x/api/shopee/config"
	"github.com/wjpxxx/letgo/x/api/shopee/logistics/entity"
	"github.com/wjpxxx/letgo/lib"
)

//Logistics
type Logistics struct{
	Config *shopeeConfig.Config
}

//GetShippingParameter
//@Title Use this api to get shipping parameter.
//@Description https://open.shopee.com/documents?module=95&type=1&id=550&version=2
func (l *Logistics)GetShippingParameter(orderSn string)entity.GetShippingParameterResult{
	method:="logistics/get_shipping_parameter"
	params:=lib.InRow{
		"order_sn":orderSn,
	}
	result:=entity.GetShippingParameterResult{}
	err:=l.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}