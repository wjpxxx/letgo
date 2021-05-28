package plugin

import (
	"core/plugin/iplugin"
	"core/plugin/sync/syncclient"
	"core/plugin/sync/syncserver"
)

//pluginList
var pluginList map[string] iplugin.Pluginer

//Register 注册插件
func Register(name string,plg iplugin.Pluginer){
	if pluginList==nil{
		pluginList=make(map[string]iplugin.Pluginer)
	}
	pluginList[name]=plg
}
//Plugin
func Plugin(name string)iplugin.Pluginer{
	return pluginList[name]
}
//init 注册插件
func init(){
	Register("sync-server", syncserver.New())
	Register("sync-client", syncclient.New())
}