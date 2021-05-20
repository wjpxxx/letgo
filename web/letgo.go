package web

import (
	"core/web/context"
	"core/web/server"
	"html/template"
	"net/http"
	"sync"
)



var initserver *server.Server

var onceDo sync.Once

func httpServer() *server.Server{
	onceDo.Do(func(){
		initserver=server.NewServer()
	})
	return initserver
}
//Run 启动
func Run() {
	httpServer().Run()
}
//Get 请求
func Get(rootPath string,fun context.HandlerFunc){
	httpServer().RegisterRouter(http.MethodGet, rootPath, fun)
}
//Post 请求
func Post(rootPath string,fun context.HandlerFunc){
	httpServer().RegisterRouter(http.MethodPost, rootPath, fun)
}

//Any 任何请求
func Any(rootPath string,fun context.HandlerFunc){
	httpServer().RegisterRouter("Any", rootPath, fun)
}

//Put 请求
func Put(rootPath string,fun context.HandlerFunc){
	httpServer().RegisterRouter(http.MethodPut, rootPath, fun)
}

//Patch 请求
func Patch(rootPath string,fun context.HandlerFunc){
	httpServer().RegisterRouter(http.MethodPatch, rootPath, fun)
}

//Head 请求
func Head(rootPath string,fun context.HandlerFunc){
	httpServer().RegisterRouter(http.MethodHead, rootPath, fun)
}

//Options 请求
func Options(rootPath string,fun context.HandlerFunc){
	httpServer().RegisterRouter(http.MethodOptions, rootPath, fun)
}

//Delete 请求
func Delete(rootPath string,fun context.HandlerFunc){
	httpServer().RegisterRouter(http.MethodDelete, rootPath, fun)
}
//Static 静态目录
func Static(relativePath, root string) {
	httpServer().Router().Static(relativePath, root)
}
//StaticFile 静态文件
func StaticFile(relativePath, filePath string){
	httpServer().Router().StaticFile(relativePath, filePath)
}
//LoadHTMLGlob
func LoadHTMLGlob(pattern string){
	httpServer().Tmpl().LoadHTMLGlob(pattern)
}
//LoadHTMLFiles
func LoadHTMLFiles(files ...string){
	httpServer().Tmpl().LoadHTMLFiles(files...)
}
//Delims
func Delims(left,right string){
	httpServer().Tmpl().SetDelims(left,right)
}

//SetFuncMap
func SetFuncMap(funcMap template.FuncMap){
	httpServer().Tmpl().SetFuncMap(funcMap)
}