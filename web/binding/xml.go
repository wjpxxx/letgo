package binding

import (
	"encoding/xml"
	"errors"
	"net/http"
	"bytes"
)

//xmlBinding
type xmlBinding struct{}
//Name
func (xmlBinding) Name() string {
	return "xml"
}
//Bind
func (xmlBinding) Bind(req *http.Request,body []byte, value interface{}) error {
	if req==nil||body==nil{
		return errors.New("error request")
	}
	decoder:=xml.NewDecoder(bytes.NewReader(body))
	if err:=decoder.Decode(value);err!=nil{
		return err
	}
	return nil
}

//Render
func(xmlBinding)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"application/xml; charset=utf-8"})
	w.WriteHeader(code)
	return xml.NewEncoder(w).Encode(value)
}