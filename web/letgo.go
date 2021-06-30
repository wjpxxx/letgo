package web

import (
	"github.com/wjpxxx/letgo/web/context"
	"github.com/wjpxxx/letgo/web/server"
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/log"
	"github.com/wjpxxx/letgo/web/filter"
	"html/template"
	"net/http"
	"sync"
	"fmt"
	"reflect"
	"strings"
	"os"
	"os/signal"
	"syscall"
	syscontext "context"
	"time"
	"os/exec"
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
func Run(addr ...string) {
	go func(){
		pid:=os.Getpid()
		log.DebugPrint("Start web server Pid:%d",pid)
		if err:=httpServer().Run(addr...);err!=nil{
			//log.DebugPrint("letgo stop :%v", err)
		}
	}()
	waitSignal()
}
//waitSignal 监控信号
func waitSignal(){
	quit:=make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	pid:=os.Getpid()
	for {
		sig:= <-quit
		switch sig {
		case syscall.SIGINT,syscall.SIGTERM:
			//启动新进程
			startNewProcess()
			//准备关闭旧进程
			ctx,cancel:=syscontext.WithTimeout(syscontext.Background(), 10*time.Second)
			defer cancel()
			if err:=httpServer().Shutdown(ctx); err!=nil{
				log.DebugPrint("Shutdown Fail %v", err)
				return
			}
			log.DebugPrint("Shutdown Pid:%d", pid)
			return
		default:
			return
		}
	}
}

//startNewProcess 启动新进程
func startNewProcess(){
	path := os.Args[0]
	env:=os.Environ()
	var args []string
	if len(os.Args)>1{
		args=os.Args[1:]
	}
	cmd:=exec.Command(path, args...)
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr
	cmd.Env=env
	err:=cmd.Start()
	if err!=nil{
		log.DebugPrint("Restart fail %v",err)
		return
	}
	log.DebugPrint("Restart success")
}

//Run 启动
func RunTLS(certFile, keyFile string, addr ...string) {
	go func ()  {
		pid:=os.Getpid()
		log.DebugPrint("Start web server Pid:%d",pid)
		if err:=httpServer().RunTLS(certFile, keyFile,addr...);err!=nil{
			//log.DebugPrint("Letgo Start fail :%v", err)
		}
	}()
	waitSignal()
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
	httpServer().RegisterRouter("ANY", rootPath, fun)
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

//RegisterController 注册控制器
func RegisterController(controller interface{},mapMethods ...string){
	name:=getControllerName(controller)
	methods:=getControllerMethod(controller,mapMethods...)
	for _,v:=range methods{
		path:=strings.ToLower(fmt.Sprintf("/%s/%s",name,v.name))
		httpServer().RegisterRouter(v.method,path, v.fun.Interface().(func(*context.Context)))
	}
	
}
//getControllerName 获得控制器名称
func getControllerName(controller interface{})string{
	getType:=reflect.TypeOf(controller)
	name:=getType.Name()
	if name==""{
		name=getType.Elem().Name()
	}
	i:=strings.Index(name,"Controller")
	if i==-1{
		panic("The controller name must end with controller")
	}
	name=name[0:i]
	return name
}
type controllerMethod struct {
	name string
	fun reflect.Value
	method string

}
//getControllerMethod 获得控制器方法
func getControllerMethod(controller interface{},mapMethods ...string)[]controllerMethod{
	getType:=reflect.TypeOf(controller)
	getValue:=reflect.ValueOf(controller)
	var funs []controllerMethod
	mapMethod:=getMapMethods(mapMethods...)
	for i:=0;i<getType.NumMethod();i++{
		argName:=getType.Method(i).Type.In(1).Name()
		if argName==""{
			argName=getType.Method(i).Type.In(1).Elem().Name()
		}
		if argName!="Context"{
			continue
			//panic("The first parameter of the method must be *context.Context")
		}
		if (getType.Method(i).Type.NumOut()>0){
			continue
		}
		methodName:=getType.Method(i).Name
		httpMethod:="ANY"
		if _,ok:=mapMethod[strings.ToLower(methodName)];ok{
			httpMethod=mapMethod[strings.ToLower(methodName)]
		}
		fun:=controllerMethod{
			name: methodName,
			fun: getValue.Method(i),
			method: httpMethod,
		}
		funs=append(funs, fun)
	}
	return funs
}

//getMapMethods 获得方法映射
func getMapMethods(mapMethods ...string)lib.StringMap{
	mp:=make(lib.StringMap)
	for _,s:=range mapMethods{
		mpArray:=strings.Split(s,":")
		if len(mpArray)!=2{
			panic("mapMethods error")
		}
		mp[strings.ToLower(mpArray[1])]=strings.ToUpper(mpArray[0])
	}
	return mp
}

//AddFilter 添加过滤
func AddFilter(pattern string, pos int, filterFunc context.HandlerFunc){
	filter.AddFilter(pattern,pos,filterFunc)
}