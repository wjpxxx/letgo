package httpclient

import (
	"github.com/wjpxxx/letgo/lib"
	"net/http"
	"fmt"
	"testing"
)

func TestHttp(t *testing.T){
	c:=&HttpClient{}
	fmt.Println(c.WithRequestBefore(func(req *http.Request){
		fmt.Println("dddd")
	}).Post("http://api-www.yutang.cn/api/Login/getSiteInfo",lib.InRow{
		"@a":"httpclient.go",
		"c":2,
	}).Body())
}
