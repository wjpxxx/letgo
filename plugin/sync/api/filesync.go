package api

import (
	"core/file"
	"core/net/rpc"
	"core/plugin/sync/syncconfig"
	"path/filepath"
)

//FileSync 文件同步
type FileSync struct{
}

//Sync 同步文件
func (f *FileSync)Sync(message syncconfig.FileSyncMessage, out *bool) error{
	if !filepath.IsAbs(message.RemotePath) {
		message.RemotePath,_=filepath.Abs(message.RemotePath)
	}else{
		message.RemotePath=filepath.FromSlash(message.RemotePath)
	}
	f.saveFile(message)
	f.sendSlave(message)
	*out=true
	return nil
}
//saveFile
func (f *FileSync)saveFile(message syncconfig.FileSyncMessage){
	var fullName string
	if message.RelPath=="." {
		//当前目录
		fullName=filepath.Join(message.RemotePath,message.File.Name)
	}else{
		path:=filepath.Join(message.RemotePath,message.RelPath)
		file.Mkdir(path)
		fullName=filepath.Join(path,message.File.Name)
	}
	fn:=file.NewFile(fullName)
	fn.WriteAt(message.File.Data,message.File.Size-message.File.Seek)
}
//sendSlave
func (f *FileSync)sendSlave(message syncconfig.FileSyncMessage){
	for _,slave:=range message.Slave{
		msg:=syncconfig.FileSyncMessage{
			LocationPath: message.LocationPath,
			RemotePath: message.RemotePath,
			RelPath: message.RelPath,
			File: syncconfig.FileData{
				Name: message.File.Name,
				Path: message.File.Path,
				Seek: message.File.Seek,
				Size: message.File.Size,
				Data: message.File.Data,
			},
			Slave:nil,
		}
		client:=rpc.NewClient().WithAddress(slave.IP,slave.Port)
		for{
			var success bool
			client.Call("FileSync.Sync",msg, &success)
			if success {
				break
			}
			//重发
		}
		client.Close()
	}
}