package syncclient
import (
	"core/file"
	"core/lib"
	"core/net/rpc"
	"core/encry"
	"core/cache/filecache"
	"core/plugin/sync/syncconfig"
	"core/plugin/sync/walkdir"
	"fmt"
	"path/filepath"
	"core/log"
)

//Run
func Run(){
	client:=rpc.NewClient().WithAddress(config.Server.IP,config.Server.Port)
	defer client.Close()
	walkdir.Walk(config.LocationPath,&walkdir.Options{
		Callback: func(pathName, fileName, fullName string) {
			f:=file.NewFile(fullName)
			if getFileModifyTime(fullName)!=f.ModifyTime(){
				//文件变化了
				fsize:=f.Size()
				var size int64=1024*1024
				var success bool =false
				for {
					buf,seek:=f.ReadBlock(size)
					if seek>=0{
						//文件有内容并存在
						message:=packed(buf,seek,f)
						rpcCall(client,message,seek,f)
						if fsize==seek{
							success=true
							break
						}
					}else{
						break
					}
				}
				if success{
					//发送成功
					saveFileModifyTime(fullName,f)
				}
			}
		},
		Filter: config.Filter,
	})
}

//getFileModifyTime 获得文件修改时间
func getFileModifyTime(fullName string) int{
	path:="runtime/cache/sync/"
	icache:=filecache.NewFileCacheByPath(path)
	var t int
	icache.Get(encry.Md5(fullName),&t)
	return t
}

//saveFileModifyTime
func saveFileModifyTime(fullName string, f file.Filer) {
	path:="runtime/cache/sync/"
	icache:=filecache.NewFileCacheByPath(path)
	icache.Set(encry.Md5(fullName),f.ModifyTime(),-1)
}

//rpcCall 发送
func rpcCall(client *rpc.Client,message syncconfig.FileSyncMessage,seek int64,f file.Filer){
	for {
		var success bool
		//fmt.Println(fmt.Sprintf("message:%s",message))
		client.Call("FileSync.Sync",message, &success)
		if success {
			showProccess(seek,f)
			break
		}
		//不成功则重发
	}
}

//showProccess
func showProccess(seek int64,f file.Filer){
	var sended float32
	if f.Size()>0{
		sended=float32(seek)/float32(f.Size())*100
	}else{
		sended=100
	}
	log.DebugPrint("正在发送文件%s,已发送%.2f%s",f.Name(),sended,"%")
}

//packed 打包
func packed(data []byte,seek int64,f file.Filer) syncconfig.FileSyncMessage{
	var locationPath string=config.LocationPath
	if !filepath.IsAbs(config.LocationPath) {
		locationPath,_=filepath.Abs(config.LocationPath)
	}else{
		locationPath=filepath.FromSlash(config.LocationPath)
	}
	relPath,_:=filepath.Rel(locationPath,f.Path())
	return syncconfig.FileSyncMessage{
		LocationPath: locationPath,
		RemotePath: config.RemotePath,
		RelPath:relPath,
		File: syncconfig.FileData{
			Name: f.Name(),
			Path: f.Path(),
			Seek: seek,
			Size: f.Size(),
			Data: data,
		},
		Slave: config.Server.Slave,
	}
}

//config
var config syncconfig.ClientConfig
//init
func init() {
	clientFile:="config/sync_client.config"
	cfgFile:=file.GetContent(clientFile)
	if cfgFile==""{
		clientConfig:=syncconfig.ClientConfig{
			LocationPath: "./",
			RemotePath: "./",
			Filter:[]string{},
			Server: syncconfig.SyncServer{
				Server: syncconfig.Server{
					IP: "127.0.0.1",
					Port: "5566",
				},
				Slave: []syncconfig.Server{},
			},
		}
		file.PutContent(clientFile,fmt.Sprintf("%v",clientConfig))
		panic("please setting sync client config in config/sync_client.config file")
	}
	if !lib.StringToObject(cfgFile, &config) {
		panic("config/sync_client.config file format error, Please check carefully")
	}
}