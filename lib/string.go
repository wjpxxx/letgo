package lib

//字符串截取
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

//将首字母转化为大写
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

//是否包含字符
func InStringArray(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}
