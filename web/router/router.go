package router

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/web/context"
	"github.com/wjpxxx/letgo/web/filter"
	"net/http"
	"path"
	"regexp"
	"strings"
	"sync"
	"fmt"
)


var initRouter *Router

var onceDo sync.Once
//Router 路由
type Router struct {
	ctx *context.Context
	routerInfo []*RouterInfo
}
//RouterInfo 路由信息
type RouterInfo struct{
	regex *regexp.Regexp
	params lib.IntStringMap
	handler context.HandlerFunc
	isStatic bool  //是否是静态资源
	method string
	path string
}

//HttpRouter 获得路由
func HttpRouter() *Router{
	onceDo.Do(func(){
		initRouter=&Router{}
	})
	return initRouter
}
//HandleHttpRequest 处理http请求
func (r *Router)HandleHttpRequest(ctx *context.Context){
	r.ctx=ctx
	r.ctx.Init()
	//过滤
	filter.ExecFilter(filter.BEFORE_ROUTER,r.ctx)
	requestPath:=strings.ToLower(r.ctx.Request.URL.Path)
	found:=false
	for _,router:=range r.routerInfo{
		if router.method!="ANY"&&r.ctx.Request.Method!=router.method {
			continue
		}
		if !router.regex.MatchString(requestPath){
			continue
		}
		matches:=router.regex.FindStringSubmatch(requestPath)
		//fmt.Println(requestPath, matches,router.regex)
		if len(matches[0])!=len(requestPath) {
			continue
		}
		if len(router.params)>0 {
			for i,match:=range matches[1:]{
				r.ctx.Input.SetParam(router.params[i], match)
			}
		}
		filter.ExecFilter(filter.BEFORE_EXEC,r.ctx)
		//初始化参数
		r.ctx.RouterPath=router.path
		if !r.ctx.Output.HasOutput() {
			router.handler(r.ctx)
		}
		filter.ExecFilter(filter.AFTER_EXEC,r.ctx)
		found=true
	}
	//no found
	if !found {
		ctx.Output.NotFound()
	}else{
		filter.ExecFilter(filter.AFTER_ROUTER,r.ctx)
	}
	
}

//File 输出文件
func (r *Router)File(filePath string) {
	http.ServeFile(r.ctx.Writer,r.ctx.Request,filePath)
}
//StaticFile 静态文件
func (r *Router)StaticFile(relativePath, filePath string) {
	if strings.Contains(relativePath,":")|| strings.Contains(relativePath,"*") {
		panic("parameters can not be used when serving a static folder")
	}
	handler := func(c *context.Context) {
		filter.ExecFilter(filter.BEFORE_STATIC,c)
		r.File(filePath)
	}
	r.RegisterRouter(http.MethodGet,relativePath,handler)
	r.RegisterRouter(http.MethodHead,relativePath,handler)
}
//Static 静态资源
func (r *Router)Static(relativePath, root string) {
	r.StaticFS(relativePath, http.Dir(root))
}
//StaticFS 静态资源
func (r *Router)StaticFS(relativePath string, fs http.FileSystem){
	if strings.Contains(relativePath,":")|| strings.Contains(relativePath,"*") {
		panic("parameters can not be used when serving a static folder")
	}
	handler:=r.createStaticHandle(relativePath, fs)
	urlPattern:=path.Join(relativePath, "/*filepath")
	r.RegisterRouter(http.MethodGet,urlPattern,handler)
	r.RegisterRouter(http.MethodHead,urlPattern,handler)
}
//createStaticHandle 创建静态处理
func (r *Router) createStaticHandle(relativePath string, fs http.FileSystem)context.HandlerFunc{
	fileServer:=http.StripPrefix(relativePath, http.FileServer(fs))
	return func(ctx *context.Context){
		filter.ExecFilter(filter.BEFORE_STATIC,ctx)
		fileServer.ServeHTTP(ctx.Writer,ctx.Request)
	}
	
}
//RegisterRouter 注册路由
func (r *Router)RegisterRouter(method,relativePath string, handler context.HandlerFunc){
	var isStatic bool=false
	oldRelativePath:=relativePath
	parts:=strings.Split(relativePath, "/")
	paramsIndex:=0
	params:=make(lib.IntStringMap)
	for i,part:=range parts {
		if strings.HasPrefix(part, ":") {
			expr:="([^/]+)"
			if index:=strings.Index(part, "(");index!=-1{
				expr=part[index:]
				part=part[:index]
			}
			params[paramsIndex]=part
			parts[i]=expr
			paramsIndex++
		}
	}
	relativePath=strings.Join(parts, "/")
	//静态资源
	if strings.Index(relativePath, "*filepath")!=-1 {
		relativePath=strings.Replace(relativePath, "*filepath", ".*",1)
		isStatic=true
	}
	lasti:=strings.LastIndex(relativePath,"*")
	if lasti!=-1{
		relativePath=fmt.Sprintf("%s%s",relativePath[:lasti],".*")
	}
	regex,regexErr:=regexp.Compile(relativePath)
	if regexErr!=nil{
		panic(regexErr)
	}
	router:=&RouterInfo{}
	router.method=method
	router.params=params
	router.regex=regex
	router.path=oldRelativePath
	router.handler=handler
	router.isStatic=isStatic
	r.routerInfo=append(r.routerInfo, router)
}
