package plugin

import (
	"github.com/wjpxxx/letgo/plugin/iplugin"
	"github.com/wjpxxx/letgo/plugin/sync/syncclient"
	"github.com/wjpxxx/letgo/plugin/sync/syncserver"
	"github.com/wjpxxx/letgo/plugin/sync/syncconfig"
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
//SyncFile 获得文件同步对象
func SyncFile()*syncclient.FileSync{
	return syncclient.NewFileSync()
}
//SyncFile 获得文件同步对象
func SyncFileByConfig(server syncconfig.SyncServer)*syncclient.FileSync{
	return syncclient.NewFileSyncByConfig(server)
}
//init 注册插件
func init(){
	Register("sync-server", syncserver.New())
	Register("sync-file", syncclient.NewFileSync())
	Register("sync-cmd",syncclient.NewCommandSync())
}