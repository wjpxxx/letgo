package shopee

import(
	shopeeConfig "github.com/wjpxxx/letgo/x/api/shopee/config"
	"github.com/wjpxxx/letgo/x/api/shopee/auth"
	authEntity "github.com/wjpxxx/letgo/x/api/shopee/auth/entity"
	orderEntity "github.com/wjpxxx/letgo/x/api/shopee/order/entity"
	logisticsEntity "github.com/wjpxxx/letgo/x/api/shopee/logistics/entity"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"github.com/wjpxxx/letgo/x/api/shopee/order"
	"github.com/wjpxxx/letgo/x/api/shopee/logistics"
)

//Shopeer
type Shopeer interface{
	//auth
	AuthorizationURL()string
	GetAccesstoken(code string,shopID int64) authEntity.GetAccessTokenResult
	RefreshAccessToken(shop commonentity.ShopInfo)authEntity.RefreshAccessTokenResult
	//order
	GetOrderList(
	timeRangeField order.TimeRangeField,
	timeFrom,timeTo,pageSize int,
	cursor string,
	orderStatus order.OrderStatus,
	responseOptionalFields string) orderEntity.GetOrderListResult
	GetShipmentList(cursor string,pageSize int) orderEntity.GetShipmentListResult
	GetOrderDetail(orderSnList []string,responseOptionalFields ...string) orderEntity.GetOrderDetailResult
	SplitOrder(orderSn string,packageList []orderEntity.PackageListRequestEntity) orderEntity.SplitOrderResult
	UnSplitOrder(orderSn string) orderEntity.UnSplitOrderResult
	CancelOrder(orderSn string,cancelReason order.CancelReason,itemList []orderEntity.CancelOrderRequestEntity) orderEntity.CancelOrderResult
	HandleBuyerCancellation(orderSn string,operation order.Operation) orderEntity.HandleBuyerCancellationResult
	SetNote(orderSn,note string) orderEntity.SetNoteResult
	AddInvoiceData(orderSn string,invoiceData orderEntity.InvoiceDataEntity) orderEntity.AddInvoiceDataResult
	//logistics
	GetShippingParameter(orderSn string)logisticsEntity.GetShippingParameterResult
}
//Shopee
type Shopee struct{
	auth.Auth
	order.Order
	logistics.Logistics
}

//shopeeList 接口列表
var shopeeList map[string]Shopeer

//Register
func Register(name string,cfg *shopeeConfig.Config){
	shopeeList[name]=&Shopee{
		auth.Auth{Config:cfg},
		order.Order{Config:cfg},
		logistics.Logistics{Config:cfg},
	}
}
//GetApi
func GetApi(name string)Shopeer{
	return shopeeList[name];
}

//init
func init(){
	shopeeList=make(map[string]Shopeer)
}