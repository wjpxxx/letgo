package context

import (
	"core/lib"
	"core/web/input"
	"core/web/output"
	"core/web/tmpl"
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
	c.Input.Init(c.Request)
	c.Output.Init(c.Writer,c.Input,c.Tmpl.Template)
}
//FullPath
func (c *Context)FullPath()string{
	return c.Request.URL.String()
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
	hostArray:=strings.Split(c.Request.Host,":")
	c.SetCookies(name,value,3600,"",hostArray[0],false,false)
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
//GetHeader
func (c *Context)GetHeader(key string)string{
	return c.Request.Header.Get(key)
}

//NewContext 新建一个上下文
func NewContext()*Context{
	return &Context{
		Input: input.NewInput(),
		Output: output.NewOutput(),
		Tmpl:tmpl.GetTmpl(),
	}
}