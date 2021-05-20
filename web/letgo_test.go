package web

import (
	"core/lib"
	"core/web/context"
	"fmt"
	"testing"
)


func TestLetgo(t *testing.T) {
	Static("/assets/", "./assets")
	StaticFile("/1.png", "./assets/b3.jpg")
	LoadHTMLGlob("config/*")
	Get("/", func(ctx *context.Context){
		ctx.Output.HTML(200,"index.tmpl",lib.InRow{
			"title":"wjp",
		})
	})
	
	Post("/user/:id([0-9]+)", func(ctx *context.Context){
		type A struct{
			Data string `json:"data"`
			ShopID int64 `json:"shop_id"`
		}
		a:=A{}
		ctx.SetCookie("a", "234234")
		ctx.Input.BindJSON(&a)
		fmt.Println(a)
		ctx.Output.YAML(500,lib.InRow{
			"message":"123123",
			"b":2,
		})
	})

	Get("/user/:id([0-9]+)/:id3([0-9]+)", func(ctx *context.Context){
		ctx.Output.XML(200,lib.InRow{
			"message":"123123",
			"b":2,
		})
	})
	Run()
}