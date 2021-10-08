package lib

import (
	"fmt"
	"strings"
	"regexp"
)

//SubString 字符串截取
//参数str：输入字符串
//参数start：起始位置
//参数end：结束位置
//返回值：截完后的字符串
func SubString(str string, start int, end int) string {
	arr := []rune(str)
	if end == -1 {
		return string(arr[start:])
	} else {
		return string(arr[start:end])
	}

}

//FirstToUpper 将首字母转化为大写
//参数str：输入字符串
//返回值：首字母大写字符串
func FirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

//InStringArray 是否包含字符
func InStringArray(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

//ResolveAddress 解析地址
func ResolveAddress(addr []string)string{
	switch len(addr) {
	case 0:
		return ":1122"
	case 1:
		return fmt.Sprintf("%s:1122",addr[0])
	case 2:
		return fmt.Sprintf("%s:%s",addr[0],addr[1])
	default:
		panic("too many parameters")
	}
}
//ReplaceIndex 替换指定第n个处
func ReplaceIndex(s,old,new string,n int)string{
	arr:=strings.Split(s, old)
	r:=""
	for i,v:=range arr{
		if v!="" {
			if i==n{
				r+=v+new
			}else{
				r+=v+old
			}
		}
	}
	return r
}
//IsFloat 判断字符串是否是一个小数
func IsFloat(s string) bool{
	match1,_:=regexp.MatchString(`^[\+-]?\d*\.\d+$`,s)
	match2,_:=regexp.MatchString(`^[\+-]?\d+\.\d*$`,s)
	return match1||match2
}