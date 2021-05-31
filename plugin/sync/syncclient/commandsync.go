package syncclient

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/net/rpc"
	"github.com/wjpxxx/letgo/plugin/sync/syncconfig"
)

//CommandSync
type CommandSync struct{}

//Run
func (c *CommandSync)Run(values ...interface{})interface{}{
	client,err:=rpc.NewClient().WithAddress(config.Server.IP,config.Server.Port)
	if err!=nil{
		return nil
	}
	defer client.Close()
	return c.do(client,values...)
}

//NewCommandSync
func NewCommandSync()*CommandSync{
	return &CommandSync{}
}
//do
func (c *CommandSync) do(client *rpc.Client,values ...interface{})map[string]syncconfig.CmdResult{
	var dir string
	var cmd string
	if len(values)==1{
		cmd=values[0].(string)
	}else if len(values)==2{
		dir=values[0].(string)
		cmd=values[1].(string)
	}
	message:=c.packedCmd(dir,cmd)
	return c.rpcCall(client,message)
}

//rpcCall
func (c *CommandSync)rpcCall(client *rpc.Client,message syncconfig.CmdMessage)map[string]syncconfig.CmdResult{
	for {
		var result syncconfig.MessageResult=syncconfig.MessageResult{}
		client.Call("Command.Run", message,&result)
		if result.Success {
			var rs map[string]syncconfig.CmdResult
			lib.UnSerialize(result.Data, &rs)
			return rs
		}
	}
	return nil
}

//packed 打包
func (c *CommandSync) packedCmd(dir,cmd string)syncconfig.CmdMessage{
	return syncconfig.CmdMessage{
		Server: syncconfig.Server{
			IP: config.Server.IP,
			Port: config.Server.Port,
		},
		Dir: dir,
		Cmd: cmd,
		Slave:c.cmdSlave(dir,cmd),
	}
}
//cmdSlave
func (c *CommandSync) cmdSlave(dir,cmd string)[]syncconfig.CmdSlave{
	var slaves []syncconfig.CmdSlave
	for _,slave:=range config.Server.Slave{
		s:=syncconfig.CmdSlave{
			Server: syncconfig.Server{
				IP: slave.IP,
				Port: slave.Port,
			},
			Dir: dir,
			Cmd: cmd,
		}
		slaves=append(slaves, s)
	}
	return slaves
}

