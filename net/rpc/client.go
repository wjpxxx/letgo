package rpc

import (
	"core/lib"
	"core/log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//Client
type Client struct{
	address string
	conn net.Conn
}
//WithAddress 设置地址
func (c *Client)WithAddress(addr ...string)*Client{
	c.address=lib.ResolveAddress(addr)
	var err error
	c.conn,err=net.Dial("tcp", c.address)
	if err!=nil{
		log.DebugPrint("RPC Dial error %v", err)
		return nil
	}
	return c
}
//Start 启动
func (c *Client)Start()*Client{
	return c.WithAddress()
}
//Close 关闭连接
func (c *Client)Close(){
	c.conn.Close()
}
//Call 调用
func (c *Client)Call(serviceMethod string, args interface{}, reply interface{})*Client{
	var err error
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c.conn))
	defer client.Close()
	err=client.Call(serviceMethod,args,reply)
	if err!=nil{
		log.DebugPrint("RPC Call error %v",err)
		return nil
	}
	return c
}
//CallByMessage
func (c *Client)CallByMessage(message RpcMessage)*Client{
	var err error
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c.conn))
	defer client.Close()
	var reply interface{}
	err=client.Call(message.Method,message.Args,&reply)
	if err!=nil{
		log.DebugPrint("RPC Call error %v",err)
		return nil
	}
	if message.Callback!=nil{
		message.Callback(reply)
	}
	return c
}
//NewClient
func NewClient()*Client{
	return &Client{}
}
//RpcMessage
type RpcMessage struct {
	Method string
	Args interface{}
	Callback func(reply interface{})
}