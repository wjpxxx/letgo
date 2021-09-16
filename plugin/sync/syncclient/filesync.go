package syncclient

import (
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/net/rpc"
	"github.com/wjpxxx/letgo/encry"
	"github.com/wjpxxx/letgo/cache/filecache"
	"github.com/wjpxxx/letgo/plugin/sync/syncconfig"
	"github.com/wjpxxx/letgo/plugin/sync/walkdir"
	"path/filepath"
	"github.com/wjpxxx/letgo/log"
)
//FileSync
type FileSync struct {}

//Run
func (f *FileSync)Run(values ...interface{})interface{}{
	client,err:=rpc.NewClient().WithAddress(config.Server.IP,config.Server.Port)
	if err!=nil{
		return false
	}
	defer client.Close()
	log.DebugPrint("连接到服务器: %s 成功",config.Server.IP)
	//同步文件
	f.SyncFile(client)
	return true
}

//SyncFile 同步文件
func (f *FileSync) SyncFile(client *rpc.Client){
	for _,c:=range config.Paths{
		walkdir.Walk(c.LocationPath,&walkdir.Options{
			Callback: func(pathName, fileName, fullName,LocationPath,RemotePath string) {
				filer:=file.NewFile(fullName)
				if f.getFileModifyTime(fullName)!=filer.ModifyTime(){
					//文件变化了
					fsize:=filer.Size()
					var size int64=1024*1024
					var success bool =false
					for {
						buf,seek:=filer.ReadBlock(size)
						if seek>=0{
							//文件有内容并存在
							message:=f.packedFileSync(buf,seek,filer,LocationPath,RemotePath)
							//log.DebugPrint("%v",message)
							f.rpcCall(client,message,seek,filer)
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
						f.saveFileModifyTime(fullName,filer)
					}
				}
			},
			Filter: c.Filter,
			LocationPath:c.LocationPath,
			RemotePath:c.RemotePath,
		})
	}
}


//getFileModifyTime 获得文件修改时间
func (f *FileSync) getFileModifyTime(fullName string) int{
	path:="runtime/cache/sync/"
	icache:=filecache.NewFileCacheByPath(path)
	var t int
	icache.Get(encry.Md5(fullName),&t)
	return t
}

//saveFileModifyTime
func (f *FileSync) saveFileModifyTime(fullName string, filer file.Filer) {
	path:="runtime/cache/sync/"
	icache:=filecache.NewFileCacheByPath(path)
	icache.Set(encry.Md5(fullName),filer.ModifyTime(),-1)
}

//rpcCall 发送
func (f *FileSync) rpcCall(client *rpc.Client,message syncconfig.FileSyncMessage,seek int64,filer file.Filer){
	for {
		var result syncconfig.MessageResult=syncconfig.MessageResult{}
		//fmt.Println(fmt.Sprintf("message:%s",message))
		client.Call("FileSync.Sync",message, &result)
		if result.Success {
			f.showProccess(message,seek,filer)
			break
		}else{
			//发送失败
		}
		//不成功则重发
	}
}

//showProccess
func (f *FileSync) showProccess(message syncconfig.FileSyncMessage,seek int64,filer file.Filer){
	var sended float32
	if filer.Size()>0{
		sended=float32(seek)/float32(filer.Size())*100
	}else{
		sended=100
	}
	log.DebugPrint("正在发送文件%s,已发送%.2f%s",filer.FullPath(),sended,"%")
}

//packed 打包
func (f *FileSync) packedFileSync(data []byte,seek int64,filer file.Filer,LocationPath,RemotePath string) syncconfig.FileSyncMessage{
	var locationPath string=LocationPath
	if !filepath.IsAbs(LocationPath) {
		locationPath,_=filepath.Abs(LocationPath)
	}else{
		locationPath=filepath.FromSlash(LocationPath)
	}
	relPath,_:=filepath.Rel(locationPath,filer.Path())
	return syncconfig.FileSyncMessage{
		LocationPath: locationPath,
		RemotePath: RemotePath,
		RelPath:relPath,
		File: syncconfig.FileData{
			Name: filer.Name(),
			Path: filer.Path(),
			Seek: seek,
			Size: filer.Size(),
			Data: data,
		},
		Slave: config.Server.Slave,
	}
}

//NewFileSync
func NewFileSync()*FileSync{
	return &FileSync{}
}