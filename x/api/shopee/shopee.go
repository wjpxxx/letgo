package shopee

import(
	shopeeConfig "github.com/wjpxxx/letgo/x/api/shopee/config"
	"github.com/wjpxxx/letgo/x/api/shopee/auth"
	authEntity "github.com/wjpxxx/letgo/x/api/shopee/auth/entity"
	orderEntity "github.com/wjpxxx/letgo/x/api/shopee/order/entity"
	logisticsEntity "github.com/wjpxxx/letgo/x/api/shopee/logistics/entity"
	productEntity "github.com/wjpxxx/letgo/x/api/shopee/product/entity"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"github.com/wjpxxx/letgo/x/api/shopee/order"
	"github.com/wjpxxx/letgo/x/api/shopee/logistics"
	"github.com/wjpxxx/letgo/x/api/shopee/product"
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
	GetTrackingNumber(orderSn,packageNumber string,responseOptionalFields ...string)logisticsEntity.GetTrackingNumberResult
	ShipOrder(orderSn,packageNumber string,pickup *logisticsEntity.ShipOrderRequestPickupEntity,dropoff *logisticsEntity.ShipOrderRequestDropoffEntity,nonIntegrated *logisticsEntity.ShipOrderRequestNonIntegratedEntity)logisticsEntity.ShipOrderResult
	UpdateShippingOrder(orderSn,packageNumber string,pickup *logisticsEntity.UpdateShippingOrderRequestPickupEntity)logisticsEntity.UpdateShippingOrderResult
	GetShippingDocumentParameter(orderList *logisticsEntity.ShippingDocumentParameterRequestOrderListEntity)logisticsEntity.GetShippingDocumentParameterResult
	CreateShippingDocument(orderList *logisticsEntity.CreateShippingDocumentRequestOrderListEntity)logisticsEntity.CreateShippingDocumentResult
	GetShippingDocumentResult(orderList *logisticsEntity.GetShippingDocumentResultRequestOrderListEntity)logisticsEntity.GetShippingDocumentResult
	DownloadShippingDocument(orderList *logisticsEntity.DownloadShippingDocumentRequestOrderListEntity)logisticsEntity.DownloadShippingDocumentResult
	GetShippingDocumentInfo(orderSn,packageNumber string)logisticsEntity.GetShippingDocumentInfoResult
	GetTrackingInfo(orderSn,packageNumber string)logisticsEntity.GetTrackingInfoResult
	GetAddressList()logisticsEntity.GetAddressListResult
	SetAddressConfig(showPickupAddress bool,AddressTypeConfig logisticsEntity.AddressTypeConfigEntity)logisticsEntity.SetAddressConfigResult
	DeleteAddress(addressID int64)logisticsEntity.DeleteAddressResult
	GetChannelList()logisticsEntity.GetChannelListResult
	UpdateChannel(logisticsChannelID int64,enabled,preferred,codEnabled bool)logisticsEntity.UpdateChannelResult
	BatchShipOrder(orderList *logisticsEntity.BatchShipOrderRequestOrderListEntity,pickup *logisticsEntity.BatchShipOrderRequestPickupEntity,dropoff *logisticsEntity.BatchShipOrderRequestDropoffEntity,nonIntegrated *logisticsEntity.BatchShipOrderRequestNonIntegratedEntity)logisticsEntity.BatchShipOrderResult
	//product
	GetComment(itemID,commentID int64,cursor string,pageSize int)productEntity.GetCommentResult
	ReplyComment(commentList []productEntity.ReplyCommentRequestCommentEntity)productEntity.ReplyCommentResult
	GetItemBaseInfo(itemIdList []int64)productEntity.GetItemBaseInfoResult
	GetItemExtraInfo(itemIdList []int64)productEntity.GetItemExtraInfoResult
	GetItemList(offset,pageSize,updateTimeFrom,updateTimeTo int,itemStatus product.ItemStatus)productEntity.GetItemListResult
}
//Shopee
type Shopee struct{
	auth.Auth
	order.Order
	logistics.Logistics
	product.Product
}

//shopeeList 接口列表
var shopeeList map[string]Shopeer

//Register
func Register(name string,cfg *shopeeConfig.Config){
	shopeeList[name]=&Shopee{
		auth.Auth{Config:cfg},
		order.Order{Config:cfg},
		logistics.Logistics{Config:cfg},
		product.Product{Config:cfg},
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