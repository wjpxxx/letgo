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
			//rpc.NewClient().WithAddress("127.0.0.1","8080").Call("Hello.Say","nihao",&reply).Close()
			rpc.NewClient().Start().Call("Hello.Say","nihao",&reply).Close()
			fmt.Println(reply)
			rm:=rpc.RpcMessage{
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

## Command

```go
import "github.com/wjpxxx/letgo/command/command"

func main(){
	cmd:=command.New().Cd("D:\\Development\\go\\web\\src").SetCMD("dir")
	//cmd.AddPipe(New().SetCMD("find","'\\c'","'80'"))
	cmd.Run()
}
```

## Synchronize files

1.server config in config/sync_server.config

```go
{"ip":"127.0.0.1","port":"5566"}
```
start server

```sh
./main server
```

2.client config in config/sync_client.config

```go
{
    "locationPath":"D:/Development/go/web/src/sync-v2", //Local directory to be synchronized
    "remotePath":"/root/new/buyer2",	//Directory stored on the server
    "filter":[			//Filter out files that are not synchronized to the server
        "main.go",
        "runtime/cache/sync/*"
    ],
    "server":{
        "ip":"127.0.0.1",	//Server IP
        "port":"5566",
        "slave":[
            {
                "ip":"127.0.0.2",	//For other servers, it is recommended to be in the same LAN as the server
                "port":"5566"
            }
        ]
    }
}
```
start client

```sh
./main client
```

The client sends the file to be synchronized to the server, and then the server synchronizes the file to the slave

3.go code

```go
package main

import (
	"core/plugin"
	"core/plugin/sync/syncconfig"
	"fmt"
	"os"
)


func main() {
	args:=os.Args
	if len(args)>1{
		if args[1]=="server"{
			plugin.Plugin("sync-server").Run()
		} else if args[1]=="file"{
			plugin.Plugin("sync-file").Run()
		} else if args[1]=="cmd"{
			rs:=plugin.Plugin("sync-cmd").Run("/www/","ls -al")
			result:=rs.(map[string]syncconfig.CmdResult)
			fmt.Println(result)
		}
	}
```

