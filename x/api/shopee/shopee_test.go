package shopee

import(
	"testing"
	shopeeConfig "github.com/wjpxxx/letgo/x/api/shopee/config"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"github.com/wjpxxx/letgo/x/api/shopee/order"
	orderEntity "github.com/wjpxxx/letgo/x/api/shopee/order/entity"
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/lib"
	"fmt"
)
func TestShopee(t *testing.T){
	Register("shopee-api",shopeeConfig.New("https://partner.test-stable.shopeemobile.com","/api/v2/",1001219,"cea778f3b36d99bda5d16a4e511fa55f9032464940163fe4acfee13c48658f42","/shopee_callback"))
	fmt.Println(GetApi("shopee-api").AuthorizationURL())
	file.PutContent("json",fmt.Sprintf("%v",GetApi("shopee-api").GetAccesstoken("69958af305efe832865ff5eb67a9c2e3",9714)))
	fmt.Println(GetApi("shopee-api").RefreshAccessToken(*commonentity.NewShop(9714,14377,"8d15a63559e5efcf75f64e096d60071e","67decdc2a76c641ea8ba0fd2cf4cc014")))
	Register("shopee-api-v2",shopeeConfig.New("https://partner.test-stable.shopeemobile.com","/api/v2/",1001219,"cea778f3b36d99bda5d16a4e511fa55f9032464940163fe4acfee13c48658f42","/shopee_callback").SetShopInfo(&commonentity.ShopInfo{
		RefreshToken:"7b9e401bedce79d51d34ffeeee47c713",
		AccessToken:"4fe3d5896fa522574e8dc50eb265c0f0",
		ExpireIn:14376,
		ShopID:9714,
	}))
	fmt.Println(GetApi("shopee-api-v2").GetOrderList("create_time",lib.Time()-3600*24*10,lib.Time(),20,"",order.UNPAID,"order_status"))
	fmt.Println(GetApi("shopee-api-v2").GetOrderDetail([]string{"210606JQ3AFK4A"}))
	fmt.Println(GetApi("shopee-api-v2").SplitOrder("210606JQ3AFK4A",[]orderEntity.PackageListRequestEntity{
		orderEntity.PackageListRequestEntity{
			ItemList:[]orderEntity.PackageListRequestItemListEntity{
				orderEntity.PackageListRequestItemListEntity{
					ItemID:100015844,
				},
			},
		},
	}))
	fmt.Println(GetApi("shopee-api-v2").CancelOrder("210606JQ3AFK4A",order.OUT_OF_STOCK,[]orderEntity.CancelOrderRequestEntity{
		orderEntity.CancelOrderRequestEntity{
			ItemID:100015844,
			ModelID: 10000083295,
		},
	}))
	fmt.Println(GetApi("shopee-api-v2").GetShippingParameter("210606JQ3AFK4A"))
}