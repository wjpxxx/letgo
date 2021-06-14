package mysql

import (
	"github.com/wjpxxx/letgo/lib"
	"fmt"
	"math"
	"strings"
)

//db 全局变量
var db *DB
//Model 模型
type Model struct{
	tableName string
	aliasName string
	otherTableName []joinCond
	dbName string
	fields []string
	where []cond
	groupBy []string
	having []cond
	orderBy string
	offset      int
	limit       int
	lastSql string
	preSql string
	preParams []interface{}
	unionModel Modeler
	db DBer
	SoftDelete bool
}
//cond 操作
type cond struct{
	field string
	symbol string
	logic string
	value interface{}
}
//joinCond 连接条件
type joinCond struct {
	tableName string
	on []cond
}
//Modeler 模型接口
type Modeler interface{
	Fields(fields ...string) Fielder
	Count()int64
	GetLastSql() string
	GetSqlInfo()(string,[]interface{})
	GetModelSql()(string,[]interface{})
	Alias(name string) Aliaser
	Join(tableName string) Joiner
	LeftJoin(tableName string) Joiner
	RightJoin(tableName string) Joiner
	Union(model Modeler) Unioner
	Where(field string, value interface{}) Wherer
	WhereRaw(where string) Wherer
	WhereSymbol(field, symbol string, value interface{}) Wherer
	WhereIn(field string, value []interface{}) Wherer
	GroupBy(field ...string) GroupByer
	Get()lib.SqlRows
	Find()lib.SqlRow
	OrderBy(orderBy string)OrderByer
	Limit(offset,count int)Limiter
	Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow)
	Update(data lib.SqlIn)
	Insert(row lib.SqlIn)int64
	Create(row lib.SqlIn)int64
	Replace(row lib.SqlIn) int64
	InsertOnDuplicate(row lib.SqlIn,updateRow lib.SqlIn) int64
	Drop() int64
	Truncate() int64
	Delete() int64
	DB()DBer
}
//Fielder 字段暴露出去的接口
type Fielder interface{
	Alias(name string) Aliaser
	Join(tableName string) Joiner
	LeftJoin(tableName string) Joiner
	RightJoin(tableName string) Joiner
	Union(model Modeler) Unioner
	Where(field string, value interface{}) Wherer
	WhereRaw(where string) Wherer
	WhereSymbol(field, symbol string, value interface{}) Wherer
	WhereIn(field string, value []interface{}) Wherer
	GroupBy(field ...string) GroupByer
	Count()int64
	Get()lib.SqlRows
	Find()lib.SqlRow
	OrderBy(orderBy string)OrderByer
	Limit(offset,count int)Limiter
	Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow)
}
//Aliaser 另外取名接口
type Aliaser interface{
	Join(tableName string) Joiner
	LeftJoin(tableName string) Joiner
	RightJoin(tableName string) Joiner
	Union(model Modeler) Unioner
	Where(field string, value interface{}) Wherer
	WhereRaw(where string) Wherer
	WhereSymbol(field, symbol string, value interface{}) Wherer
	WhereIn(field string, value []interface{}) Wherer
	GroupBy(field ...string) GroupByer
	Count()int64
	Get()lib.SqlRows
	Find()lib.SqlRow
	OrderBy(orderBy string)OrderByer
	Limit(offset,count int)Limiter
	Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow)
	Update(data lib.SqlIn)
}
//Joiner 连接接口
type Joiner interface{
	On(field string, value interface{}) Oner
	OnRaw(on string) Oner
	OnSymbol(field, symbol string, value interface{}) Oner
	OnIn(field string, value []interface{}) Oner
}
//Oner on连接条件
type Oner interface{
	OrOn(field string, value interface{}) Oner
	OrOnRaw(on string) Oner
	OrOnSymbol(field, symbol string, value interface{}) Oner
	OrOnIn(field string, value []interface{}) Oner
	AndOn(field string, value interface{}) Oner
	AndOnRaw(on string) Oner
	AndOnSymbol(field, symbol string, value interface{}) Oner
	AndOnIn(field string, value []interface{}) Oner
	Where(field string, value interface{}) Wherer
	WhereRaw(where string) Wherer
	WhereSymbol(field, symbol string, value interface{}) Wherer
	WhereIn(field string, value []interface{}) Wherer
	Join(tableName string) Joiner
	LeftJoin(tableName string) Joiner
	RightJoin(tableName string) Joiner
	GroupBy(field ...string) GroupByer
	Count()int64
	Get()lib.SqlRows
	Find()lib.SqlRow
	OrderBy(orderBy string)OrderByer
	Limit(offset,count int)Limiter
	Union(model Modeler) Unioner
	Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow)
	Update(data lib.SqlIn)
}
//Wherer 条件接口
type Wherer interface{
	AndWhere(field string, value interface{}) Wherer
	AndWhereRaw(where string) Wherer
	AndWhereSymbol(field, symbol string, value interface{}) Wherer
	AndWhereIn(field string, value []interface{}) Wherer
	OrWhere(field string, value interface{}) Wherer
	OrWhereRaw(where string) Wherer
	OrWhereSymbol(field, symbol string, value interface{}) Wherer
	OrWhereIn(field string, value []interface{}) Wherer
	GroupBy(field ...string) GroupByer
	Get()lib.SqlRows
	Find()lib.SqlRow
	Count()int64
	OrderBy(orderBy string)OrderByer
	Limit(offset,count int)Limiter
	Union(model Modeler) Unioner
	Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow)
	Update(data lib.SqlIn)
}
//GroupByer 分组接口
type GroupByer interface{
	Having(field string, value interface{})Havinger
	HavingRaw(having string)Havinger
	HavingSymbol(field, symbol string, value interface{})Havinger
	HavingIn(field string, value []interface{})Havinger
	Get()lib.SqlRows
	Find()lib.SqlRow
	Count()int64
	OrderBy(orderBy string)OrderByer
	Limit(offset,count int)Limiter
	Union(model Modeler) Unioner
	Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow)
}
//Havinger having条件
type Havinger interface{
	AndHaving(field string, value interface{})Havinger
	AndHavingRaw(having string)Havinger
	AndHavingSymbol(field, symbol string, value interface{})Havinger
	AndHavingIn(field string, value []interface{})Havinger
	OrHaving(field string, value interface{})Havinger
	OrHavingRaw(having string)Havinger
	OrHavingSymbol(field, symbol string, value interface{})Havinger
	OrHavingIn(field string, value []interface{})Havinger
	Get()lib.SqlRows
	Find()lib.SqlRow
	Count()int64
	OrderBy(orderBy string)OrderByer
	Limit(offset,count int)Limiter
	Union(model Modeler) Unioner
	Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow)
}
//OrderByer 排序
type OrderByer interface{
	Get()lib.SqlRows
	Find()lib.SqlRow
	Limit(offset,count int)Limiter
	Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow)
}
//Limiter
type Limiter interface{
	Get()lib.SqlRows
	Find()lib.SqlRow
}
//Unioner
type Unioner interface{
	Count()int64
	Get()lib.SqlRows
	Find()lib.SqlRow
	OrderBy(orderBy string)OrderByer
	Limit(offset,count int)Limiter
	Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow)
}
//Init 初始化
func (m *Model) Init(dbName,tableName string) Modeler{
	m.tableName=tableName
	m.dbName=dbName
	m.db=db.SetDB(m.dbName,m.dbName)
	return m
}
//Init 初始化
func (m *Model) InitByConnectName(connectName,dbName,tableName string) Modeler{
	m.tableName=tableName
	m.dbName=dbName
	m.db=db.SetDB(connectName,m.dbName)
	return m
}
//DBer 获得数据库接口
func (m *Model) DB()DBer{
	return m.db
}
//Fields 查询字段
func (m *Model)Fields(fields ...string) Fielder{
	m.fields=fields
	return m
}
//GetLastSql 获得最后执行的sql
func (m *Model)GetLastSql() string{
	return m.lastSql
}
//GetSqlInfo 获得最后执行的sql
func (m *Model)GetSqlInfo()(string,[]interface{}){
	return m.preSql,m.preParams
}
//Alias 命名
func (m *Model)Alias(name string) Aliaser{
	m.aliasName=name
	m.tableName=fmt.Sprintf("%s as %s", m.tableName, name)
	return m
}
//Join 连接查询
func (m *Model)Join(tableName string) Joiner {
	m.otherTableName=append(
		m.otherTableName,
		joinCond{
			tableName:fmt.Sprintf("INNER JOIN %s",tableName),
		},
	)
	return m
}
//LeftJoin 连接查询
func (m *Model)LeftJoin(tableName string) Joiner {
	m.otherTableName=append(
		m.otherTableName,
		joinCond{
			tableName:fmt.Sprintf("LEFT JOIN %s",tableName),
		},
	)
	return m
}
//RightJoin 连接查询
func (m *Model)RightJoin(tableName string) Joiner {
	m.otherTableName=append(
		m.otherTableName,
		joinCond{
			tableName:fmt.Sprintf("RIGHT JOIN %s",tableName),
		},
	)
	return m
}
//setJoinCond 设置连接条件
func (m *Model) setJoinCond(field,symbol,logic string,value interface{}){
	index:=len(m.otherTableName)-1
	if index>=0 {
		var icond cond
		if len(m.otherTableName[index].on)==0 {
			icond=cond{
				field: field,
				symbol: symbol,
				logic:"",
				value:value,
			}
		}else{
			icond=cond{
				field: field,
				symbol: symbol,
				logic:logic,
				value:value,
			}
		}
		m.otherTableName[index].on=append(
			m.otherTableName[index].on,
			icond,
		)
	}
}
//setWhereCond 设置条件
func (m *Model)setWhereCond(field,symbol,logic string,value interface{}){
	var icond cond
	if len(m.where)==0 {
		icond=cond{
			field: field,
			symbol: symbol,
			logic:"",
			value:value,
		}
	}else{
		icond=cond{
			field: field,
			symbol: symbol,
			logic:logic,
			value:value,
		}
	}
	m.where=append(m.where, icond)
}

//setHavingCond 设置条件
func (m *Model)setHavingCond(field,symbol,logic string,value interface{}){
	var icond cond
	if len(m.having)==0 {
		icond=cond{
			field: field,
			symbol: symbol,
			logic:"",
			value:value,
		}
	}else{
		icond=cond{
			field: field,
			symbol: symbol,
			logic:logic,
			value:value,
		}
	}
	m.having=append(m.having, icond)
}
//On Join 条件on
func (m *Model)On(field string, value interface{}) Oner{
	m.setJoinCond(field,"=", "and",value)
	return m
}
//OnRaw Join 条件on
func (m *Model)OnRaw(on string) Oner{
	m.setJoinCond(fmt.Sprintf("(%s)",on),"", "and",nil)
	return m
}
//OnSymbol Join 条件on
func (m *Model)OnSymbol(field, symbol string, value interface{}) Oner{
	m.setJoinCond(field,symbol, "and",value)
	return m
}
//OnIn Join 条件on
func (m *Model)OnIn(field string, value []interface{}) Oner{
	var values []string
	for _,v:=range value{
		values=append(values, lib.InterfaceToString(v))
	}
	valuesStr:=fmt.Sprintf("(%s)",strings.Join(values,",")) 
	m.setJoinCond(field,"in", "and", valuesStr)
	return m
}
//OrOn Join 条件on
func (m *Model)OrOn(field string, value interface{}) Oner{
	m.setJoinCond(field,"=", "or", value)
	return m
}
//OrOnRaw Join 条件on
func (m *Model)OrOnRaw(on string) Oner{
	m.setJoinCond(fmt.Sprintf("(%s)",on),"", "or", nil)
	return m
}
//OrOnSymbol Join 条件on
func (m *Model)OrOnSymbol(field, symbol string, value interface{}) Oner{
	m.setJoinCond(field,symbol, "or", value)
	return m
}
//OrOnIn Join 条件on
func (m *Model)OrOnIn(field string, value []interface{}) Oner{
	var values []string
	for _,v:=range value{
		values=append(values, lib.InterfaceToString(v))
	}
	valuesStr:=fmt.Sprintf("(%s)",strings.Join(values,",")) 
	m.setJoinCond(field,"in", "or", valuesStr)
	return m
}
//AndOn Join 条件on
func (m *Model)AndOn(field string, value interface{}) Oner{
	m.setJoinCond(field,"=", "and", value)
	return m
}
//AndOnRaw Join 条件on
func (m *Model)AndOnRaw(on string) Oner{
	m.setJoinCond(fmt.Sprintf("(%s)",on),"", "and", nil)
	return m
}
//AndOnSymbol Join 条件on
func (m *Model)AndOnSymbol(field, symbol string, value interface{}) Oner{
	m.setJoinCond(field,symbol, "and", value)
	return m
}
//AndOnIn Join 条件on
func (m *Model)AndOnIn(field string, value []interface{}) Oner{
	var values []string
	for _,v:=range value{
		values=append(values, lib.InterfaceToString(v))
	}
	valuesStr:=fmt.Sprintf("(%s)",strings.Join(values,",")) 
	m.setJoinCond(field,"in", "and", valuesStr)
	return m
}
//Union 联合查询
func (m *Model)Union(model Modeler) Unioner{
	m.unionModel=model
	return m
}
//Where where条件
func (m *Model)Where(field string, value interface{}) Wherer{
	m.setWhereCond(field,"=", "and",value)
	return m
}
//WhereRaw where条件
func (m *Model)WhereRaw(where string) Wherer{
	m.setWhereCond(fmt.Sprintf("(%s)",where),"", "and",nil)
	return m
}
//WhereSymbol where条件
func (m *Model)WhereSymbol(field, symbol string, value interface{}) Wherer{
	m.setWhereCond(field,symbol, "and",value)
	return m
}
//WhereIn where条件
func (m *Model)WhereIn(field string, value []interface{}) Wherer{
	var values []string
	for _,v:=range value{
		values=append(values, lib.InterfaceToString(v))
	}
	valuesStr:=fmt.Sprintf("(%s)",strings.Join(values,",")) 
	m.setWhereCond(field,"in", "and",valuesStr)
	return m
}
//AndWhere where条件
func (m *Model)AndWhere(field string, value interface{}) Wherer{
	m.setWhereCond(field,"=", "and",value)
	return m
}
//AndWhereRaw where条件
func (m *Model)AndWhereRaw(where string) Wherer{
	m.setWhereCond(fmt.Sprintf("(%s)",where),"", "and",nil)
	return m
}
//AndWhereSymbol where条件
func (m *Model)AndWhereSymbol(field, symbol string, value interface{}) Wherer{
	m.setWhereCond(field,symbol, "and",value)
	return m
}
//AndWhereIn where条件
func (m *Model)AndWhereIn(field string, value []interface{}) Wherer{
	var values []string
	for _,v:=range value{
		values=append(values, lib.InterfaceToString(v))
	}
	valuesStr:=fmt.Sprintf("(%s)",strings.Join(values,",")) 
	m.setWhereCond(field,"in", "and",valuesStr)
	return m
}
//OrWhere where条件
func (m *Model)OrWhere(field string, value interface{}) Wherer{
	m.setWhereCond(field,"=", "or",value)
	return m
}
//OrWhereRaw where条件
func (m *Model)OrWhereRaw(where string) Wherer{
	m.setWhereCond(fmt.Sprintf("(%s)",where),"", "or",nil)
	return m
}
//OrWhereSymbol where条件
func (m *Model)OrWhereSymbol(field, symbol string, value interface{}) Wherer{
	m.setWhereCond(field,symbol, "or",value)
	return m
}
//OrWhereIn where条件
func (m *Model)OrWhereIn(field string, value []interface{}) Wherer{
	var values []string

	for _,v:=range value{
		values=append(values, lib.InterfaceToString(v))
	}
	valuesStr:=fmt.Sprintf("(%s)",strings.Join(values,",")) 
	m.setWhereCond(field,"in", "or",valuesStr)
	return m
}

//GroupBy 分组
func (m *Model)GroupBy(field ...string) GroupByer{
	m.groupBy=field
	return m
}

//Having having条件
func (m *Model)Having(field string, value interface{}) Havinger{
	m.setHavingCond(field,"=", "and",value)
	return m
}
//HavingRaw having条件
func (m *Model)HavingRaw(having string) Havinger{
	m.setHavingCond(fmt.Sprintf("(%s)",having),"", "and",nil)
	return m
}
//HavingSymbol having条件
func (m *Model)HavingSymbol(field, symbol string, value interface{}) Havinger{
	m.setHavingCond(field,symbol, "and",value)
	return m
}
//HavingIn having条件
func (m *Model)HavingIn(field string, value []interface{}) Havinger{
	var values []string
	for _,v:=range value{
		values=append(values, lib.InterfaceToString(v))
	}
	valuesStr:=fmt.Sprintf("(%s)",strings.Join(values,",")) 
	m.setHavingCond(field,"in", "and",valuesStr)
	return m
}

//AndHaving having条件
func (m *Model)AndHaving(field string, value interface{}) Havinger{
	m.setHavingCond(field,"=", "and",value)
	return m
}
//AndHavingRaw having条件
func (m *Model)AndHavingRaw(having string) Havinger{
	m.setHavingCond(fmt.Sprintf("(%s)",having),"", "and",nil)
	return m
}
//AndHavingSymbol having条件
func (m *Model)AndHavingSymbol(field, symbol string, value interface{}) Havinger{
	m.setHavingCond(field,symbol, "and",value)
	return m
}
//AndHavingIn having条件
func (m *Model)AndHavingIn(field string, value []interface{}) Havinger{
	var values []string
	for _,v:=range value{
		values=append(values, lib.InterfaceToString(v))
	}
	valuesStr:=fmt.Sprintf("(%s)",strings.Join(values,",")) 
	m.setHavingCond(field,"in", "and",valuesStr)
	return m
}

//OrHaving having条件
func (m *Model)OrHaving(field string, value interface{}) Havinger{
	m.setHavingCond(field,"=", "or",value)
	return m
}
//OrHavingRaw having条件
func (m *Model)OrHavingRaw(having string) Havinger{
	m.setHavingCond(fmt.Sprintf("(%s)",having),"", "or",nil)
	return m
}
//OrHavingSymbol having条件
func (m *Model)OrHavingSymbol(field, symbol string, value interface{}) Havinger{
	m.setHavingCond(field,symbol, "or",value)
	return m
}
//OrHavingIn having条件
func (m *Model)OrHavingIn(field string, value []interface{}) Havinger{
	var values []string
	for _,v:=range value{
		values=append(values, lib.InterfaceToString(v))
	}
	valuesStr:=fmt.Sprintf("(%s)",strings.Join(values,",")) 
	m.setHavingCond(field,"in", "or",valuesStr)
	return m
}
//getTables 表名
func (m *Model)getTables() (string,[]interface{}){
	var table string=m.tableName
	var values []interface{}
	for _,joinTable:=range m.otherTableName{
		ons,v:=m.getOn(joinTable.on)
		if len(v)>0{
			values=append(values, v...)
		}
		table+=fmt.Sprintf(" %s %s",joinTable.tableName,ons) 
		
	}
	return table,values
}
//getOn 获得on条件
func (m *Model)getOn(ons []cond)(string,[]interface{}){
	on:="on"
	var values []interface{}
	for i,o:=range ons{
		var q string=""
		if o.value!=nil{
			values=append(values, o.value)
			q="?"
		}
		if i==0{
			on+=fmt.Sprintf("%s %s %s "+q,o.logic,o.field,o.symbol)
		}else{
			on+=fmt.Sprintf(" %s %s %s "+q,o.logic,o.field,o.symbol)
		}
	}
	return on,values
}
//getWhere
func (m *Model)getWhere()(string,[]interface{}){
	where:=""
	var values []interface{}
	for i,w:=range m.where{
		var q string=""
		if w.value!=nil{
			values=append(values, w.value)
			q="?"
		}
		if i==0{
			where+=fmt.Sprintf("%s %s %s "+q,w.logic,w.field,w.symbol)
		}else{
			where+=fmt.Sprintf(" %s %s %s "+q,w.logic,w.field,w.symbol)
		}
	}
	if m.SoftDelete {
		logic:="and"
		if len(m.where)==0{
			logic=""
		}
		if m.aliasName!="" {
			where+=fmt.Sprintf(" %s %s %s ",logic,m.aliasName+".delete_time","=?")
			values=append(values,-1)
		}else{
			where+=fmt.Sprintf(" %s %s %s ",logic,"delete_time","=?")
			values=append(values,-1)
		}
		//fmt.Println(where,values);
	}
	return where,values
}
//getGroup
func (m *Model)getGroup()(string,[]interface{}){
	if len(m.groupBy)==0{
		return "",nil
	}
	group:=fmt.Sprintf(" GROUP BY %s ",strings.Join(m.groupBy,","))
	if len(m.having)>0{
		group=group+" HAVING "
	}
	var values []interface{}
	for i,w:=range m.having{
		var q string=""
		if w.value!=nil{
			values=append(values, w.value)
			q="?"
		}
		if i==0{
			group+=fmt.Sprintf("%s %s %s "+q,w.logic,w.field,w.symbol)
		}else{
			group+=fmt.Sprintf(" %s %s %s "+q,w.logic,w.field,w.symbol)
		}
	}
	return group,values
}
//getOrderBy 获得排序
func (m *Model)getOrderBy()string{
	if m.orderBy!=""{
		return fmt.Sprintf(" ORDER BY %s",m.orderBy)
	}
	return ""
}
//getLimit 分页
func (m *Model)getLimit()string{
	if m.limit==0&&m.offset==0{
		return ""
	}
	return fmt.Sprintf(" LIMIT %d,%d",m.offset,m.limit)
}
//getFields 获得查询字段
func (m *Model)getFields() string{
	if len(m.fields)==0{
		return "*"
	}else{
		return strings.Join(m.fields,",")
	}
	
}
//GetModelSql 模型sql
func (m *Model)GetModelSql()(string,[]interface{}){
	tableName,values:=m.getTables()
	where,whereValues:=m.getWhere()
	if len(whereValues)>0{
		values=append(values, whereValues...)
	}
	group,groupValues:=m.getGroup()
	if group!=""{
		where=where+group
		if len(groupValues)>0{
			values=append(values, groupValues...)
		}
	}
	fields:=m.getFields()
	var sql string
	if where!=""{
		sql=fmt.Sprintf("select %s from %s where %s", fields,tableName, where)
	}else{
		sql=fmt.Sprintf("select %s from %s", fields,tableName)
	}
	return sql,values
}
//Count 统计数量
func (m *Model)Count()int64{
	tableName,values:=m.getTables()
	where,whereValues:=m.getWhere()
	if len(whereValues)>0{
		values=append(values, whereValues...)
	}
	group,groupValues:=m.getGroup()
	if group!=""{
		where=where+group
		if len(groupValues)>0{
			values=append(values, groupValues...)
		}
	}
	if m.unionModel!=nil{
		union,unionValues:=m.unionModel.GetModelSql()
		if where!="" {
			where=where+" UNION "+union
		}else{
			where="1=1 UNION "+union
		}
		if len(unionValues)>0{
			values=append(values, unionValues...)
		}
	}
	where=where+" LIMIT 0,1"
	table:=NewTable(db,tableName)
	rows:= table.Select("count(1) as c",where,values...)
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
	if len(rows)==1{
		return rows[0]["c"].Int64()
	}
	return -1
}
//Get 获得列表
func (m *Model)Get()lib.SqlRows{
	tableName,values:=m.getTables()
	where,whereValues:=m.getWhere()
	if len(whereValues)>0{
		values=append(values, whereValues...)
	}
	group,groupValues:=m.getGroup()
	if group!=""{
		where=where+group
		if len(groupValues)>0{
			values=append(values, groupValues...)
		}
	}
	if m.unionModel!=nil{
		union,unionValues:=m.unionModel.GetModelSql()
		if where!="" {
			where=where+" UNION "+union
		}else{
			where="1=1 UNION "+union
		}
		if len(unionValues)>0{
			values=append(values, unionValues...)
		}
	}
	where=where+m.getOrderBy()
	where=where+m.getLimit()
	table:=NewTable(db,tableName)
	fields:=m.getFields()
	rows:= table.Select(fields,where,values...)
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
	return rows
}
//Find 获得单条数据
func (m *Model)Find()lib.SqlRow{
	tableName,values:=m.getTables()
	where,whereValues:=m.getWhere()
	if len(whereValues)>0{
		values=append(values, whereValues...)
	}
	group,groupValues:=m.getGroup()
	if group!=""{
		where=where+group
		if len(groupValues)>0{
			values=append(values, groupValues...)
		}
	}
	if m.unionModel!=nil{
		union,unionValues:=m.unionModel.GetModelSql()
		if where!="" {
			where=where+" UNION "+union
		}else{
			where="1=1 UNION "+union
		}
		if len(unionValues)>0{
			values=append(values, unionValues...)
		}
	}
	where=where+m.getOrderBy()
	where=where+" LIMIT 0,1"
	table:=NewTable(db,tableName)
	fields:=m.getFields()
	rows:= table.Select(fields,where,values...)
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
	if len(rows)==1{
		return rows[0]
	}
	return nil
}
//Pager 分页查询
func (m *Model)Pager(page, pageSize int)(lib.SqlRows,lib.SqlRow){
	offset := pageSize * (page - 1)
	m.offset=offset
	m.limit=pageSize
	var pageInfo lib.SqlRow=lib.SqlRow{}
	list:=m.Get()
	total:=m.Count()
	pageInfo["total"]=(&lib.Data{}).Set(total)
	pageInfo["pageCount"] = (&lib.Data{}).Set(int(math.Ceil(float64(total) / float64(pageSize))))
	pageInfo["page"]=(&lib.Data{}).Set(page)
	pageInfo["pageSize"]=(&lib.Data{}).Set(pageSize)
	return list,pageInfo
}
//OrderBy 排序
func (m *Model)OrderBy(orderBy string) OrderByer{
	m.orderBy=orderBy
	return m
}
//Limit 排序
func (m *Model)Limit(offset,count int)Limiter{
	m.limit=count
	m.offset=offset
	return m
}
//Update 更新操作
func (m *Model)Update(data lib.SqlIn){
	tableName,values:=m.getTables()
	where,whereValues:=m.getWhere()
	table:=NewTable(db,tableName)
	table.Update(data,values,where,whereValues...)
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
}
//Insert 插入操作
func (m *Model)Insert(row lib.SqlIn)int64{
	table:=NewTable(db,m.tableName)
	id:=table.Insert(row)
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
	return id
}
//Create 插入数据
func (m *Model)Create(row lib.SqlIn)int64{
	return m.Insert(row)
}
//Replace 插入操作
func (m *Model)Replace(row lib.SqlIn) int64{
	table:=NewTable(db,m.tableName)
	effects:=table.Replace(row)
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
	return effects
}
//InsertOnDuplicate 如果你插入的记录导致一个UNIQUE索引或者primary key(主键)出现重复，那么就会认为该条记录存在，则执行update语句而不是insert语句，反之，则执行insert语句而不是更新语句。
func (m *Model)InsertOnDuplicate(row lib.SqlIn,updateRow lib.SqlIn) int64{
	table:=NewTable(db,m.tableName)
	effects:=table.InsertOnDuplicate(row,updateRow)
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
	return effects
}
//Drop 删除表
func (m *Model)Drop() int64 {
	table:=NewTable(db,m.tableName)
	effects:=table.Drop()
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
	return effects
}
//Truncate 清空表
func (m *Model)Truncate() int64{
	table:=NewTable(db,m.tableName)
	effects:=table.Truncate()
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
	return effects
}
//Delete 删除
func (m *Model)Delete() int64{
	tableName,values:=m.getTables()
	where,whereValues:=m.getWhere()
	table:=NewTable(db,tableName)
	effects:=table.Delete(values,where,whereValues...)
	m.lastSql=table.GetLastSql()
	m.preSql,m.preParams=table.GetSqlInfo()
	return effects
}
//init 初始化连接池
func init(){
	db=NewDB()
}
//GetDB 获得数据库对象
func GetDB()*DB{
	return db
}
//NewModel 新建一个模型
func NewModel(dbName,tableName string) Modeler{
	model:=Model{}
	return model.Init(dbName,tableName)
}

//NewModelByConnectName 新建一个模型
func NewModelByConnectName(connectName,dbName,tableName string) Modeler{
	model:=Model{}
	return model.InitByConnectName(connectName,dbName,tableName)
}