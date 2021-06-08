package binding


import (
	"net/http"
)

//textBinding
type textBinding struct{}

//Render
func(textBinding)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"application/html; charset=utf-8"})
	w.WriteHeader(code)
	str:=value.(string)
	_,err:=w.Write([]byte(str))
	return err
}