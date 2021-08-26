package context

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/web/input"
	"github.com/wjpxxx/letgo/web/output"
	"github.com/wjpxxx/letgo/web/tmpl"
	"github.com/wjpxxx/letgo/web/session"
	"net/http/httputil"
	"net/http"
	"net/url"
	"strings"
)

//HandlerFunc 请求处理函数
type HandlerFunc func(*Context)

//Context 上下文对象
type Context struct {
	Request *http.Request
	Writer http.ResponseWriter
	Input *input.Input
	Output *output.Output
	Session session.Sessioner
	RouterPath string
	sameSite http.SameSite
	Tmpl *tmpl.Tmpl
}


//重置
func (c *Context)Reset(){
	c.Input=input.NewInput()
	c.Output=output.NewOutput()
	c.RouterPath=""
}
//Init 初始化
func (c *Context)Init(){
	c.Session.Start()	//启动session
	c.Input.Init(c.Request)
	c.Output.Init(c.Writer,c.Input,c.Tmpl.Template)
}
//FullPath
func (c *Context)FullPath()string{
	return c.Request.URL.String()
}
//Router
func (c *Context)Router()string{
	requestPath:=strings.ToLower(c.Request.URL.Path)
	if requestPath==""{
		return "/"
	}
	return requestPath
}
//SetCookie 设置cookie
func (c *Context)SetCookies(name,value string,maxAge int,path,domain string,secure,httpOnly bool) {
	if path=="" {
		path="/"
	}
	http.SetCookie(c.Writer,&http.Cookie{
		Name: name,
		Value: url.QueryEscape(value),
		MaxAge: maxAge,
		Path: path,
		Domain: domain,
		SameSite: c.sameSite,
		Secure: secure,
		HttpOnly: httpOnly,
	})
}
//SetCookie 设置cookie
func (c *Context)SetCookie(name,value string) {
	c.SetCookies(name,value,3600,"",c.Host(),false,false)
}
//SetCookieByExpire
func (c *Context)SetCookieByExpire(name,value string,expire int){
	c.SetCookies(name,value,expire,"",c.Host(),false,false)
}
//Host
func (c *Context)Host()string{
	hostArray:=strings.Split(c.Request.Host,":")
	return hostArray[0]
}
//HttpOrigin
func (c *Context)HttpOrigin()string{
	return c.Request.Header.Get("Origin");
}

//Cookie 获得cookie
func (c *Context)Cookie(name string)*lib.Data{
	cookie,err:=c.Request.Cookie(name)
	if err!=nil{
		return nil
	}
	val,_:=url.QueryUnescape(cookie.Value)
	return &(lib.Data{Value: val})
}

//SetSameSite
func (c *Context)SetSameSite(sameSite http.SameSite) {
	c.sameSite=sameSite
}
//ContentType
func (c *Context)ContentType()string{
	return c.Input.ContentType()
}

//DumpRequest
func (c *Context)DumpRequest()string{
	dump, _ := httputil.DumpRequest(c.Request, true)
	return string(dump);
}
//GetHeader
func (c *Context)GetHeader(key string)string{
	return c.Request.Header.Get(key)
}

//NewContext 新建一个上下文
func NewContext()*Context{
	ctx:= &Context{
		Input: input.NewInput(),
		Output: output.NewOutput(),
		Tmpl:tmpl.GetTmpl(),
	}
	ctx.Session=session.GetSession(ctx)
	return ctx
}