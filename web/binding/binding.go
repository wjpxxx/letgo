package binding

import (
	"net/http"
)

//Binding
type Binding interface {
	Name() string
	Bind(*http.Request,interface{})error
}
//Render
type Rendering interface{
	Render(int,http.ResponseWriter,interface{})error
}
const (
	MIMEJSON              = "application/json"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEYAML              = "application/x-yaml"
)
var (
	JSON =jsonBinding{}
	XML=xmlBinding{}
	YAML=yamlBinding{}
	JSONP=jsonpBinding{}
)
func NewBind(contentType string)Binding{
	switch contentType {
	case MIMEJSON:
		return JSON
	case MIMEXML,MIMEXML2:
		return XML
	case MIMEYAML:
		return YAML
	default:
		panic("unknown type")
	}
}

func NewRender(contentType string)Rendering{
	switch contentType {
	case MIMEJSON:
		return JSON
	case MIMEXML,MIMEXML2:
		return XML
	case MIMEYAML:
		return YAML
	default:
		panic("unknown type")
	}
}
//writeContentType
func writeContentType(w http.ResponseWriter, value []string) {
	header:=w.Header()
	if v:=header["Content-Type"];len(v)==0{
		header["Content-Type"] = value
	}
}