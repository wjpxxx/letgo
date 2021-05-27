package syncserver

import (
	"core/file"
	"core/lib"
	"core/net/rpc"
	"core/plugin/sync/api"
	"core/plugin/sync/syncconfig"
	"fmt"
)

//Run
func Run(){
	rpc.NewServer().Register(new(api.FileSync)).Register(new(api.Command)).Run(config.IP,config.Port)
}

//sync server config
var config syncconfig.Server
//init
func init(){
	serverFile:="config/sync_server.config"
	cfgFile:=file.GetContent(serverFile)
	if cfgFile==""{
		serverConfig:=syncconfig.Server{
			IP: "127.0.0.1",
			Port: "5566",
		}
		file.PutContent(serverFile,fmt.Sprintf("%v",serverConfig))
		panic("please setting sync server config in config/sync_server.config file")
	}
	lib.StringToObject(cfgFile, &config)
}