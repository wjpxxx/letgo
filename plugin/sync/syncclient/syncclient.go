package syncclient
import (
	"core/file"
	"core/lib"
	"core/plugin/sync/syncconfig"
	"fmt"
)

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