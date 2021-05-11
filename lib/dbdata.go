package lib

//SqlRows 查询多行
type SqlRows []SqlRow

//SqlRow 查询单行
type SqlRow Row

//SqlIn sql插入更新数据格式
type SqlIn InRow
//Row 数据
type Row map[string] *Data
//InRow 数据
type InRow map[string]interface{}
//IntRow 整型数据
type IntRow map[int]interface{}

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