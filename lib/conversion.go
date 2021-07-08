package lib

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"bytes"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

//字符串转float32
func StrToFloat32(str string) float32 {
	vv, err := strconv.ParseFloat(str, 32)
	if err == nil {
		return float32(vv)
	}
	return 0
}

//字符串转float64
func StrToFloat64(str string) float64 {
	vv, err := strconv.ParseFloat(str, 32)
	if err == nil {
		return vv
	}
	return 0
}

//字符串转int
func StrToInt(str string) int {
	vv, err := strconv.Atoi(str)
	if err == nil {
		return vv
	}
	return 0
}
//字符串转uint
func StrToUInt(str string) uint {
	vv, err := strconv.Atoi(str)
	if err == nil {
		return uint(vv)
	}
	return 0
}
//字符串转int64
func StrToInt64(str string) int64 {
	vv, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return vv
	}
	return 0
}

//将float64转int
func Float64ToInt(f float64) int {
	return int(f)
}

//将Interface转int
func InterfaceToInt(data interface{}) int {
	str := fmt.Sprintf("%v", data)
	return StrToInt(str)
}

//将Interface转String
func InterfaceToString(data interface{}) string {
	str := fmt.Sprintf("%v", data)
	return str
}

//将Interface转int
func InterfaceToInt64(data interface{}) int64 {
	str := fmt.Sprintf("%v", data)
	return StrToInt64(str)
}

//float64转int64
func Float64ToInt64(data float64) int64 {
	return int64(data)
}

//RowsToSqlRows sql.Rows转 SqlRows
func RowsToSqlRows(rows *sql.Rows) SqlRows{
	cols, err := rows.Columns()
	if err != nil {
		return nil
	}
	scanArgs := make([]interface{}, len(cols))
	values := make([]interface{}, len(cols))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	var list SqlRows
	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			break
		}
		record := make(SqlRow)
		for i, col := range values {
			row := &Data{}
			row.Set(col)
			record[cols[i]] = row
		}
		list = append(list, record)
	}
	return list
}

//ObjectToString 将对象转成json字符串
func ObjectToString(data interface{}) string {
	js, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(js)
}

//StringToObject json字符串转对象
func StringToObject(str string, data interface{}) bool {
	err := json.Unmarshal([]byte(str), data)
	if err == nil {
		return true
	}
	return false
}
//Int64ArrayToInterfaceArray int64转[]interface{}
func Int64ArrayToInterfaceArray(data []int64)[]interface{}{
	var it []interface{}
	for _,v:=range data{
		it=append(it, v)
	}
	return it
}

//StringArrayToInterfaceArray string转[]interface{}
func StringArrayToInterfaceArray(data []string)[]interface{}{
	var it []interface{}
	for _,v:=range data{
		it=append(it, v)
	}
	return it
}

//interface转ArrayString
func InterfaceArrayToArrayString(list []interface{}) []string {
	var rp []string
	for _, v := range list {
		data := &Data{Value: v}
		rp = append(rp, data.String())
	}
	return rp
}

//interface转ArrayString
func Int64ArrayToArrayString(list []int64) []string {
	var rp []string
	for _, v := range list {
		rp = append(rp, fmt.Sprintf("%d",v))
	}
	return rp
}

//Utf8ToGb2312 UTF8转GBK2312
func Utf8ToGb2312(src string) string {
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GB18030.NewEncoder()))
	return string(data)
}

//Gb2312ToUtf8 utf8转gbk
func Gb2312ToUtf8(src string) string {
	data, _ := ioutil.ReadAll(simplifiedchinese.GB18030.NewDecoder().Reader(bytes.NewReader([]byte(src))))
	rts := string(data)
	return rts
}