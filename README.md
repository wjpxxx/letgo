# letgo Web Framework

letgo is an open-source, high-performance web framework for the Go programming language.

## Contents

- [letgo Web Framework](#letgo-web-framework)
  - [Installation](#installation)
  - [Quick start](#quick-start)

## Installation

To install letgo package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.12+ is required**), then you can use the below Go command to install letgo.

```sh
$ go get -u github.com/wjpxxx/letgo
```

2. Import it in your code:

```go
import "github.com/wjpxxx/letgo"
```
## Quick start

## Web

```go
package main

import (
    "github.com/wjpxxx/letgo/web"
    "github.com/wjpxxx/letgo/web/context"
)

func main() {
	web.LoadHTMLGlob("templates/*")
	web.Get("/", func(ctx *context.Context){
		//ctx.Output.Redirect(301,"http://www.baidu.com")
		x:=ctx.Input.Param("a")
		if x!=nil{
			ctx.Session.Set("a",x.Int())
			var a int
			ctx.Session.Get("a",&a)
			fmt.Println("a:",a,"x:",x.Int())
		}
		ctx.Output.HTML(200,"index.tmpl",lib.InRow{
			"title":"wjp",
		})
	})
	
	web.Post("/user/:id([0-9]+)", func(ctx *context.Context){
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

	web.Get("/user/:id([0-9]+)/:id3([0-9]+)", func(ctx *context.Context){
		ctx.Output.XML(200,lib.InRow{
			"message":"123123",
			"b":2,
		})
	})

	web.Run()
}
```

## Serving static files

```go
func main() {
	web.Static("/assets/", "./assets")
	web.StaticFile("/1.png", "./assets/b3.jpg")
	web.Run()
}
```

## Register controller

```go
type UserController struct{}
func (c *UserController)Add(ctx *context.Context){
	ctx.Output.JSON(200,lib.InRow{
		"a":1,
		"b":"wjp",
	})
}
func main() {
	c:=&UserController{}
	web.RegisterController(c)
	c2:=&UserController{}
	web.RegisterController(c2,"get:add")
	web.Run()
}
```

## Model operation

```go
package main

import "github.com/wjpxxx/letgo/db/mysql"
func main() {
    model:=NewModel("dbname","tablename")
        model.Fields("*").
                Alias("m").
                Join("sys_shopee_shop as s").
                On("m.id","s.master_id").
                OrOn("s.master_id",1).
                AndOnRaw("m.id=1 or m.id=2").
                LeftJoin("sys_lazada_shop as l").
                On("m.id", "l.master_id").
                WhereRaw("m.id=1").
                AndWhere("m.id",2).
                OrWhereIn("m.id",lib.Int64ArrayToInterfaceArray(ids)).
                GroupBy("m.id").
                Having("m.id",1).
                AndHaving("m.id",1).
                OrderBy("m.id desc").Find()
        fmt.Println(model.GetLastSql())
}
```

## RPC

```go
import "github.com/wjpxxx/letgo/net/rpc"

func main(){
	s:=rpc.NewServer()
	//s.RegisterName("Hello",new(Hello))
	s.Register(new(Hello))
	go func(){
		for{
			time.Sleep(10*time.Second)
			var reply string
			rpc.NewClient().Start().Call("Hello.Say","nihao",&reply).Close()
			fmt.Println(reply)
			rm:=RpcMessage{
				Method: "Hello.Say",
				Args: "rpc message",
				Callback: func(a interface{}){
					fmt.Println(a.(string))
				},
			}
			rpc.NewClient().Start().CallByMessage(rm).Close()
		}
	}()
	s.Run()
}
```