package lib

import (
	"encoding/xml"
	"reflect"
)

//SqlRows 查询多行
type SqlRows []SqlRow

//ToOutput
func (s SqlRows)ToOutput()[]InRow{
	var list []InRow
	for _,data:=range s{
		d:=make(InRow)
		for k,v:=range data{
			typeOf :=reflect.TypeOf(v.Value)
			switch typeOf.String() {
				case "int64":
					d[k]=v.Int64()
				case "int":
					d[k]=v.Int()
				case "float64":
					d[k]=v.Float64()
				case "float32":
					d[k]=v.Float32()
				case "bool":
					d[k]=v.Value
				default:
					d[k]=v.String()
			}
			
		}
		list=append(list, d)
	}
	return list
}

//SqlRow 查询单行
type SqlRow Row

//ToOutput
func (s SqlRow)ToOutput()InRow{
	d:=make(InRow)
	for k,v:=range s{
		typeOf :=reflect.TypeOf(v.Value)
		switch typeOf.String() {
			case "int64":
				d[k]=v.Int64()
			case "int":
				d[k]=v.Int()
			case "float64":
				d[k]=v.Float64()
			case "float32":
				d[k]=v.Float32()
			case "bool":
				d[k]=v.Value
			default:
				d[k]=v.String()
		}
	}
	return d
}

//SqlIn sql插入更新数据格式
type SqlIn InRow
//Row 数据
type Row map[string] *Data
//InRow 数据
type InRow map[string]interface{}
//IntRow 整型数据
type IntRow map[int]interface{}
//IntStringMap 
type IntStringMap map[int]string
//StringMap 
type StringMap map[string]string

//MergeInRow 合并InRow
func MergeInRow(values ...InRow)InRow{
	result:=make(InRow)
	for _,row:=range values{
		for k,v:=range row{
			result[k]=v
		}
	}
	return result
}

//MergeInRow 合并InRow
func MergeIntRow(values ...IntRow)IntRow{
	result:=make(IntRow)
	for _,row:=range values{
		for k,v:=range row{
			result[k]=v
		}
	}
	return result
}
//MergeRow 合并Row
func MergeRow(values ...Row)Row{
	result:=make(Row)
	for _,row:=range values{
		for k,v:=range row{
			result[k]=v
		}
	}
	return result
}
//MarshalXML
func (i InRow)MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t:=xml.ProcInst{
		Target:"xml",
		Inst:[]byte(`version="1.0" encoding="UTF-8"`),
	}
	e.EncodeToken(t)
	start.Name=xml.Name{
		Space: "",
		Local: "map",
	}
	if err:=e.EncodeToken(start);err!=nil{
		return err
	}
	for key,value:=range i{
		elem:=xml.StartElement{
			Name: xml.Name{
				Space: "",
				Local: key,
			},
			Attr: []xml.Attr{},
		}
		if err:=e.EncodeElement(value,elem);err!=nil{
			return err
		}
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}