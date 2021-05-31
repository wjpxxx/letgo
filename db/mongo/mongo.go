package mongo

import (
	"context"
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/lib"
	"fmt"
	"time"
	"github.com/wjpxxx/letgo/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoDB
type MongoDB struct {
	databaseName string
	connectName string
	dbPool MongoPooler
}

//SetPool 设置连接池
func (db *MongoDB)SetPool(pool MongoPooler)*MongoDB{
	db.dbPool=pool
	return db
}

//SetDB 设置连接名和数据库名称
func (db *MongoDB)SetDB(connectName,databaseName string)*MongoDB{
	db.connectName=connectName
	db.databaseName=databaseName
	return db
}

//NewDB 
func NewDB()*MongoDB{
	dbFile:="config/mongo_db.config"
	cfgFile:=file.GetContent(dbFile)
	var configs []MongoConnect
	if cfgFile==""{
		db:=MongoConnect{
			Name:"connectName",
			Database:"databaseName",
			UserName:"userName",
			Password:"password",
			Hosts:[]Host{
				Host{
					Hst:"127.0.0.1",
					Port:"27017",
				},
			},
			ConnectTimeout: 10,
			ExecuteTimeout: 10,
		}
		configs=append(configs,db)
		file.PutContent(dbFile,fmt.Sprintf("%v",configs))
		panic("please setting mongo database config in config/mongo_db.config file")
	}
	lib.StringToObject(cfgFile, &configs)
	var db MongoDB
	return db.SetPool(NewPools(configs))
}

//Tabler
type Tabler interface{
	SetDB(db *MongoDB)
	InsertOne(document interface{})primitive.ObjectID
	InsertMany(documents []interface{})[]primitive.ObjectID
	UpdateOne(filter interface{}, update interface{}) *mongo.UpdateResult
	UpdateMany(filter interface{}, update interface{}) *mongo.UpdateResult
	UpdateByID(id interface{}, update interface{}) *mongo.UpdateResult
	ReplaceOne(filter interface{}, update interface{}) *mongo.UpdateResult
	DeleteOne(filter interface{}) *mongo.DeleteResult
	DeleteMany(filter interface{}) *mongo.DeleteResult
	FindOne(filter interface{},result interface{})
	Find(filter interface{},result interface{})
	Pager(filter interface{},result interface{},page,pageSize int64)
	Aggregate(pipeline interface{},result interface{})
}

//Table 操作表
type Table struct{
	tableName string
	db *MongoDB
}

//NewTable 初始化表
func NewTable(db *MongoDB,tableName string) Tabler{
	var table *Table=&Table{}
	table.tableName=tableName
	table.SetDB(db)
	return table
}

//SetDB 设置数据库
func (t *Table) SetDB(db *MongoDB){
	t.db=db
}

//getDB
func (t *Table)getDB() *DBInfo{
	return t.db.dbPool.GetDB(t.db.connectName)
}

//InsertOne
func (t *Table)InsertOne(document interface{})primitive.ObjectID{
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).InsertOne(ctx,document)
	if err!=nil{
		log.DebugPrint("mongo InsertOne error %v", err)
		panic(err)
	}
	return res.InsertedID.(primitive.ObjectID)
}

//InsertMany
func (t *Table)InsertMany(documents []interface{})[]primitive.ObjectID{
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).InsertMany(ctx,documents)
	if err!=nil{
		log.DebugPrint("mongo InsertMany error %v", err)
		panic(err)
	}
	var result []primitive.ObjectID
	for _,v:=range res.InsertedIDs{
		result=append(result, v.(primitive.ObjectID))
	}
	return result
}

//UpdateOne
func (t *Table)UpdateOne(filter interface{}, update interface{}) *mongo.UpdateResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).UpdateOne(ctx,filter,update)
	if err!=nil{
		log.DebugPrint("mongo UpdateOne error %v", err)
		panic(err)
	}
	return res
}

//UpdateMany
func (t *Table)UpdateMany(filter interface{}, update interface{}) *mongo.UpdateResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).UpdateMany(ctx,filter,update)
	if err!=nil{
		log.DebugPrint("mongo UpdateMany error %v", err)
		panic(err)
	}
	return res
}

//UpdateByID
func (t *Table)UpdateByID(id interface{}, update interface{}) *mongo.UpdateResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).UpdateByID(ctx,id,update)
	if err!=nil{
		log.DebugPrint("mongo UpdateByID error %v", err)
		panic(err)
	}
	return res
}

//ReplaceOne
func (t *Table)ReplaceOne(filter interface{}, update interface{}) *mongo.UpdateResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).ReplaceOne(ctx,filter,update)
	if err!=nil{
		log.DebugPrint("mongo ReplaceOne error %v", err)
		panic(err)
	}
	return res
}

//DeleteOne
func (t *Table)DeleteOne(filter interface{}) *mongo.DeleteResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).DeleteOne(ctx,filter)
	if err!=nil{
		log.DebugPrint("mongo DeleteOne error %v", err)
		panic(err)
	}
	return res
}

//DeleteMany
func (t *Table)DeleteMany(filter interface{}) *mongo.DeleteResult {
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	res,err:=db.Database.Collection(t.tableName).DeleteMany(ctx,filter)
	if err!=nil{
		log.DebugPrint("mongo DeleteMany error %v", err)
		panic(err)
	}
	return res
}

//FindOne
func (t *Table)FindOne(filter interface{},result interface{}){
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	err:=db.Database.Collection(t.tableName).FindOne(ctx,filter).Decode(result)
	if err!=nil{
		log.DebugPrint("mongo FindOne error %v", err)
		panic(err)
	}
}

//Find
func (t *Table)Find(filter interface{},result interface{}){
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	cur,err:=db.Database.Collection(t.tableName).Find(ctx,filter)
	if err!=nil{
		log.DebugPrint("mongo Find error %v", err)
		panic(err)
	}
	defer cur.Close(context.Background())
	err=cur.All(context.Background(), result)
	if err!=nil{
		log.DebugPrint("mongo Find cur error %v", err)
		panic(err)
	}
}

//Pager
func (t *Table)Pager(filter interface{},result interface{},page,pageSize int64){
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	option:=options.Find()
	option.SetLimit(pageSize)
	option.SetSkip(pageSize * (page-1))
	cur,err:=db.Database.Collection(t.tableName).Find(ctx,filter,option)
	if err!=nil{
		log.DebugPrint("mongo Pager error %v", err)
		panic(err)
	}
	defer cur.Close(context.Background())
	err=cur.All(context.Background(), result)
	if err!=nil{
		log.DebugPrint("mongo Pager cur error %v", err)
		panic(err)
	}
}

//Aggregate
func (t *Table)Aggregate(pipeline interface{},result interface{}){
	db:=t.getDB()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Duration(db.Config.ExecuteTimeout)*time.Second)
	defer func(){
		cancel()
		db.Close()
	}()
	cur,err:=db.Database.Collection(t.tableName).Aggregate(ctx,pipeline)
	if err!=nil{
		log.DebugPrint("mongo Aggregate error %v", err)
		panic(err)
	}
	defer cur.Close(context.Background())
	err=cur.All(context.Background(), result)
	if err!=nil{
		log.DebugPrint("mongo Aggregate cur error %v", err)
		panic(err)
	}
}