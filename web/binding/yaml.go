package binding

import (
	"gopkg.in/yaml.v2"
	"errors"
	"net/http"
)
type yamlBinding struct{}

//Name
func (yamlBinding) Name() string {
	return "yaml"
}
//Bind
func (yamlBinding) Bind(req *http.Request, value interface{}) error {
	if req==nil||req.Body==nil{
		return errors.New("error request")
	}
	decoder:=yaml.NewDecoder(req.Body)
	if err:=decoder.Decode(value);err!=nil{
		return err
	}
	return nil
}


//Render
func(yamlBinding)Render(code int,w http.ResponseWriter,value interface{})error{
	writeContentType(w,[]string{"application/x-yaml; charset=utf-8"})
	w.WriteHeader(code)
	yamlData,err:=yaml.Marshal(value)
	if err!=nil{
		return err
	}
	_,err=w.Write(yamlData)
	return err
}