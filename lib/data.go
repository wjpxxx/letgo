package lib

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"
)

//Data 数据
type Data struct{
	Value interface{}
}
//Set 设置值
func (d *Data)Set(value interface{})*Data{
	d.Value=value
	return d
}
//Get 获得值
func (d *Data)Get()interface{}{
	return d.Value
}
//String 字符串输出
func (d *Data) String() string {
	typeOf :=reflect.TypeOf(d.Value)
	if typeOf!=nil{
		switch typeOf.String() {
		case "[]byte":
			b:=d.Value.([]byte)
			if len(b)>=6{
				flg:=[]byte{0,1,1,1,1,1}
				if string(b[:6])==string(flg){
					var s string
					UnSerialize(b,&s)
					return s
				}
			}
			return string(b)
		case "string":
			return d.Value.(string)
		case "[]string":
			return "[" + strings.Join(d.Value.([]string), " ") + "]"
		case "int64":
			return strconv.FormatInt(d.Int64(), 10)
		case "int":
			return strconv.Itoa(d.Value.(int))
		case "[]uint8":
			b:=d.Value.([]uint8)
			if len(b)>=6{
				flg:=[]byte{0,1,1,1,1,1}
				if string(b[:6])==string(flg){
					var s string
					UnSerialize(b,&s)
					return s
				}
			}
			return string(b)
		case "float64":
			return strconv.FormatFloat(d.Float64(), 'f', -1, 64)
		case "float32":
			return strconv.FormatFloat(float64(d.Float32()), 'f', -1, 32)
		case "*multipart.FileHeader":
			return d.Value.(*multipart.FileHeader).Filename
		case "[]*multipart.FileHeader":
			fh := d.Value.([]*multipart.FileHeader)
			fname := make([]string, len(fh))
			for k, v := range fh {
				fname[k] = v.Filename
			}
			return "[" + strings.Join(fname, " ") + "]"
		default:
			return fmt.Sprintf("%v", d.Value)
		}
	}
	return ""
}

//SqlRows 查询数据多行
func (d *Data)SqlRows()SqlRows{
	if v, ok := d.Value.(SqlRows); ok {
		return v
	}
	return nil
}

//SqlRow 查询数据单行
func (d *Data) SqlRow() SqlRow {
	if v, ok := d.Value.(SqlRow); ok {
		return v
	}
	return nil
}


//WhetherFloat64 输出浮点型
func (d *Data) WhetherFloat64() (float64, bool) {
	if v, ok := d.Value.(float64); ok {
		return v, true
	} else if v2, ok2 := d.Value.([]byte); ok2 {
		vv, err := strconv.ParseFloat(string(v2), 32)
		if err == nil {
			return vv, true
		}
		var i float64
		UnSerialize(v2,&i)
		return i,true
	}
	return 0, false

}

//Float64 输出浮点型
func (d *Data) Float64() float64 {
	r, _ := d.WhetherFloat64()
	return r
}
//WhetherArrayByte 输出[]byte字节数组
func (d *Data) WhetherArrayByte() ([]byte, bool) {
	if v, ok := d.Value.([]byte); ok {
		return v, true
	}
	return nil, false
}
//ArrayByte 输出[]byte字节数组
func (d *Data)ArrayByte()[]byte{
	r, _ := d.WhetherArrayByte()
	return r
}
//WhetherFloat32 输出浮点型
func (d *Data) WhetherFloat32() (float32, bool) {
	if v, ok := d.Value.(float32); ok {
		return v, true
	} else if v2, ok2 := d.Value.([]byte); ok2 {
		vv, err := strconv.ParseFloat(string(v2), 32)
		if err == nil {
			return float32(vv), true
		}
		var i float32
		UnSerialize(v2,&i)
		return i,true
	} else if v2, ok2 := d.Value.(string); ok2 {
		vv, err := strconv.ParseFloat(v2, 32)
		if err == nil {
			return float32(vv), true
		}
	} else if v2, ok2 := d.Value.(float64); ok2 {
		return float32(v2), true
	} else if v2, ok2 := d.Value.(int); ok2 {
		return float32(v2), true
	} else if v2, ok2 := d.Value.(int64); ok2 {
		return float32(v2), true
	}
	return 0, false

}

//Float32 输出浮点型
func (d *Data) Float32() float32 {
	r, _ := d.WhetherFloat32()
	return r
}

//WhetherInt 输出int类型
func (d *Data) WhetherInt() (int, bool) {
	if v, ok := d.Value.(int); ok {
		return v, true
	} else if v2, ok2 := d.Value.([]byte); ok2 {
		vv, err := strconv.Atoi(string(v2))
		if err == nil {
			return vv, true
		}
		var i int
		UnSerialize(v2,&i)
		return i,true
	} else if v2, ok2 := d.Value.(string); ok2 {
		vv, err := strconv.Atoi(v2)
		if err == nil {
			return vv, true
		}
	} else if v2, ok2 := d.Value.(int64); ok2 {
		return int(v2), true
	} else if v2, ok2 := d.Value.(bool); ok2 {
		if v2 {
			return 1, true
		} else {
			return 0, true
		}
	} else if v2, ok2 := d.Value.(float64); ok2 {
		return Float64ToInt(v2), true
	}
	return 0, false

}

//Int 输出int类型
func (d *Data) Int() int {
	r, _ := d.WhetherInt()
	return r
}

//WhetherInt64 输出int64类型
func (d *Data) WhetherInt64() (int64, bool) {
	if v, ok := d.Value.(int64); ok {
		return v, true
	} else if v2, ok2 := d.Value.([]byte); ok2 {
		vv, err := strconv.ParseInt(string(v2), 10, 64)
		if err == nil {
			return vv, true
		}
		var i int64
		UnSerialize(v2,&i)
		return i,true
	} else if v2, ok2 := d.Value.(string); ok2 {
		vv, err := strconv.ParseInt(v2, 10, 64)
		if err == nil {
			return vv, true
		}
	} else if v2, ok2 := d.Value.(float64); ok2 {
		return Float64ToInt64(v2), true
	}
	return 0, false

}

//Int64 输出int64类型
func (d *Data) Int64() int64 {
	r, _ := d.WhetherInt64()
	return r
}

//WhetherArrayString 输出[]string字符串数组
func (d *Data) WhetherArrayString() ([]string, bool) {
	if v, ok := d.Value.([]string); ok {
		return v, true
	} else if v, ok := d.Value.([]interface{}); ok {
		return InterfaceArrayToArrayString(v), true
	}
	return nil, false
}

//ArrayString 输出[]string字符串数组
func (d *Data) ArrayString() ([]string) {
	r, _ := d.WhetherArrayString()
	return r
}
var fix []byte=[]byte{0,1,1,1,1,1}
//序列化数据
//参数data:待序列化数据
//返回值:序列化后的数据
func Serialize(data interface{}) []byte {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	enc.Encode(data)
	return append(fix,result.Bytes()...) 
}

//反序列化数据
//参数data:序列化数据
//参数rdata:反序列化后的数据
func UnSerialize(data []byte, rdata interface{}) {
	decoder := gob.NewDecoder(bytes.NewReader(data[len(fix):]))
	decoder.Decode(rdata)
}
