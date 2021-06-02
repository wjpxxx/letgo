package output

import (
	"github.com/wjpxxx/letgo/web/binding"
	"github.com/wjpxxx/letgo/web/input"
	"html/template"
	"net/http"
)

//Output
type Output struct{
	writer http.ResponseWriter
	in *input.Input
	templ *template.Template
	status int
}

//Init 初始化
func (o *Output)Init(writer http.ResponseWriter,in *input.Input,templ *template.Template) {
	o.writer=writer
	o.in=in
	o.templ=templ
	o.status=0
}
//Header 设置头
func (o *Output)Header(key,value string){
	o.writer.Header().Set(key,value)
}
//JSON
func (o *Output)JSON(code int, value interface{})error{
	return o.Render(code,value, binding.JSON)
}

//JSONP
func (o *Output)JSONP(code int, value interface{})error{
	return o.Render(code,value, binding.JSONP)
}
//Render
func (o *Output)Render(code int,value interface{},bind binding.Rendering)error{
	if o.status==0&&code>0{
		o.status=code
		err:= bind.Render(code,o.writer,value)
		return err
	}
	return nil
}

//HTML
func (o *Output)HTML(code int,name string, value interface{})error{
	bind:=binding.NewHTML(name,o.templ)
	return o.Render(code,value, bind)
}

//XML
func (o *Output)XML(code int, value interface{})error{
	return o.Render(code,value, binding.XML)
}

//YAML
func (o *Output)YAML(code int, value interface{})error{
	return o.Render(code,value, binding.YAML)
}
//Redirect 跳转
func(o *Output)Redirect(code int,location string){
	if o.status==0&&code>0{
		o.status=code
		http.Redirect(o.writer,o.in.R(),location,code)
	}
}

//NotFound 404
func(o *Output)NotFound(){
	if o.status==0{
		o.status=404
		http.NotFound(o.writer, o.in.R())
	}
}
//HasOutput 是否输出了 true已经输出 false 未输出
func(o *Output)HasOutput()bool{
	if o.status==0{
		return false
	}
	return true
}
//NewInput 新建一个input
func NewOutput()*Output{
	return &Output{
	}
}