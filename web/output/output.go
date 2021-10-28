package output

import (
	"github.com/wjpxxx/letgo/web/binding"
	"github.com/wjpxxx/letgo/web/input"
	"github.com/wjpxxx/letgo/lib"
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
	if o!=nil&&o.writer!=nil&&o.writer.Header()!=nil{
		o.writer.Header().Set(key,value)
	}
	
}
//JSON
func (o *Output)JSON(code int, value interface{})error{
	return o.Render(code,value, binding.JSON)
}

//JSONOK
func (o *Output)JSONOK(code int, message string)error{
	return o.JSON(code,lib.InRow{
		"code":1,
		"success":true,
		"msg":message,
		"err":"",
		"sub_code":"success",
	})
}

//JSONOK
func (o *Output)JSONERROR(code int, message,subCode string)error{
	return o.JSON(code,lib.InRow{
		"code":0,
		"success":false,
		"msg":"",
		"err":message,
		"sub_code":subCode,
	})
}

//JSONOK
func (o *Output)JSONFail(code int, message string)error{
	return o.JSON(code,lib.InRow{
		"code":0,
		"success":false,
		"msg":"",
		"err":message,
		"sub_code":"fail",
	})
}

//JSONObject
func(o *Output)JSONObject(code int,info interface{})error{
	return o.JSON(code,lib.InRow{
		"code":1,
		"success":true,
		"msg":"获取成功",
		"err":"",
		"info":info,
		"sub_code":"info.success",
	})
}

//JSONList
func(o *Output)JSONList(code int,list interface{})error{
	return o.JSON(code,lib.InRow{
		"code":1,
		"success":true,
		"msg":"获取成功",
		"err":"",
		"list":list,
		"sub_code":"list.success",
	})
}

//JSONPager
func(o *Output)JSONPager(code int,list interface{},pager interface{})error{
	return o.JSON(code,lib.InRow{
		"code":1,
		"success":true,
		"msg":"获取成功",
		"err":"",
		"list":list,
		"pager":pager,
		"sub_code":"pager.success",
	})
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

//Text
func (o *Output)Text(code int, value interface{})error{
	return o.Render(code,value, binding.TEXT)
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