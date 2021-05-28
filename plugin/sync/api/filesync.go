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
func (f *FileSync)Sync(message syncconfig.FileSyncMessage, out *syncconfig.MessageResult) error{
	if !filepath.IsAbs(message.RemotePath) {
		message.RemotePath,_=filepath.Abs(message.RemotePath)
	}else{
		message.RemotePath=filepath.FromSlash(message.RemotePath)
	}
	f.saveFile(message)
	err:=f.sendSlave(message)
	if err!=nil{
		return err
	}
	out.Success=true
	out.Code=200
	out.Err=""
	out.Msg="成功"
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
func (f *FileSync)sendSlave(message syncconfig.FileSyncMessage)error{
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
		client,err:=rpc.NewClient().WithAddress(slave.IP,slave.Port)
		if err!=nil{
			return err
		}
		for{
			var result syncconfig.MessageResult=syncconfig.MessageResult{}
			_,err=client.Call("FileSync.Sync",msg, &result)
			if err!=nil{
				return err
			}
			if result.Success {
				break
			}
			//重发
		}
		client.Close()
	}
	return nil
}