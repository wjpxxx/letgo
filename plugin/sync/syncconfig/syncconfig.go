package syncconfig

import "core/lib"

//ClientConfig 客户端配置文件
type ClientConfig struct {
	LocationPath string `json:"locationPath"`
	RemotePath string `json:"remotePath"`
	Filter []string `json:"filter"`
	Server SyncServer `json:"server"`
}

//String
func (c ClientConfig)String()string{
	return lib.ObjectToString(c)
}
//SyncServer
type SyncServer struct{
	Server
	Slave []Server	`json:"slave"`
}

//String
func (s SyncServer)String()string{
	return lib.ObjectToString(s)
}

//Server 服务器信息
type Server struct{
	IP string `json:"ip"`
	Port string `json:"port"`
}
//String
func (s Server)String()string{
	return lib.ObjectToString(s)
}

//FileSyncMessage 同步文件信息
type FileSyncMessage struct {
	LocationPath string `json:"locationPath"`
	RemotePath string `json:"remotePath"`
	RelPath string `json:"relPath"`
	File FileData `json:"file"`
	Slave []Server `json:"slave"`
}
//String
func (f FileSyncMessage)String()string{
	return lib.ObjectToString(f)
}
//FileData
type FileData struct{
	Name string `json:"name"`
	Path string `json:"path"`
	Seek int64 `json:"seek"`
	Size int64 `json:"size"`
	Data []byte `json:"data"`
}