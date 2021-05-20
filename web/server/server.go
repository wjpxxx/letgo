package server

import (
	"core/lib"
	"core/web/context"
	"core/web/router"
	"core/web/tmpl"
	"net/http"
	"sync"
)

//Server httpServer类
type Server struct {
	pool sync.Pool
	route *router.Router
}
//NewServer
func NewServer()*Server{
	sr:= &Server{
		route: router.HttpRouter(),
	}
	sr.pool.New=func()interface{}{
		return context.NewContext()
	}
	return sr
}
//ServeHTTP http请求
func (s *Server) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	c:=s.pool.Get().(*context.Context)
	c.Request=r
	c.Writer=w
	c.Reset()
	s.handleHttpRequest(c)
	s.pool.Put(c)
}
//handleHttpRequest 处理http请求
func (s *Server)handleHttpRequest(c *context.Context){
	s.route.HandleHttpRequest(c)
}
//Run 启动服务
func (s *Server)Run(addr ...string)error {
	address:=lib.ResolveAddress(addr)
	//fmt.Println(address)
	http.ListenAndServe(address, s)
	return nil
}

//RunTLS 启动服务
func(s *Server)RunTLS(certFile, keyFile string, addr ...string)error{
	address:=lib.ResolveAddress(addr)
	http.ListenAndServeTLS(address,certFile,keyFile,s)
	return nil
}
//RegisterRouter 注册路由
func(s *Server)RegisterRouter(method,relativePath string, handler context.HandlerFunc){
	s.route.RegisterRouter(method,relativePath,handler)
}
//Router 获得路由
func (s *Server)Router()*router.Router{
	return s.route
}
//Tmpl 获得模板对象
func (s *Server)Tmpl()*tmpl.Tmpl{
	return tmpl.GetTmpl()
}