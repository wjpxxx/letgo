package output

import (
	"core/web/binding"
	"core/web/input"
	"html/template"
	"net/http"
)

//Output
type Output struct{
	writer http.ResponseWriter
	in *input.Input
	templ *template.Template
}

//Init 初始化
func (o *Output)Init(writer http.ResponseWriter,in *input.Input,templ *template.Template) {
	o.writer=writer
	o.in=in
	o.templ=templ
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
	err:= bind.Render(code,o.writer,value)
	return err
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
	http.Redirect(o.writer,o.in.R(),location,code)
}

//NewInput 新建一个input
func NewOutput()*Output{
	return &Output{
	}
}