package httpclient

import (
	"github.com/wjpxxx/letgo/lib"
	"fmt"
	"testing"
)

func TestHttp(t *testing.T){
	c:=&HttpClient{}
	fmt.Println(c.Post("http://api-www.yutang.cn/api/Login/getSiteInfo",lib.InRow{
		"@a":"httpclient.go",
		"c":2,
	}).Body())
}