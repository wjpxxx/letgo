package mysql

import (
	"github.com/wjpxxx/letgo/lib"
    "fmt"
    "testing"
	"time"
)
func TestDB(t *testing.T){
	var db DB
	var configs []MysqlConnect
	var slaves []SlaveDB
	master:=SlaveDB{
			Name:"xingtool_base",
			DatabaseName:"xingtool_base",
			UserName:"wjp",
			Password:"wjp",
			Host:"127.0.0.1",
			Port:"3306",
			Charset:"utf8mb4",
			MaxOpenConns:20,
			MaxIdleConns:10,
	}
	slaves=append(slaves,SlaveDB{
			Name:"xingtool_base",
			DatabaseName:"xingtool_base",
			UserName:"wjp",
			Password:"wjp",
			Host:"127.0.0.1",
			Port:"3306",
			Charset:"utf8mb4",
			MaxOpenConns:20,
			MaxIdleConns:10,
	})
	configs=append(configs, MysqlConnect{
		Master:master,
		Slave:slaves,
	})
	db.SetPool(NewPools(configs))
	db.SetDB("xingtool_base", "xingtool_base")
	db.BeginTransaction()
	table:=NewTable(&db,"sys_user_master")
	list:=table.Select("*","id=?", 2)
	fmt.Println(list[0]["nick_name"].String())
	table.Update(lib.SqlIn{
		"db_code":"001",
		"table_code":"004",
	},nil,"id=?",2)
	i:=table.Delete(nil,"id=?",10)
	fmt.Println(i)
	time.Sleep(1*time.Second)
	db.Commit()
}

func TestNewDB(t *testing.T) {
	db:=NewDB()
	db.SetDB("xingtool_base", "xingtool_base")
	db.BeginTransaction()
	table:=NewTable(db,"sys_user_master")
	list:=table.Select("*","id=?", 2)
	fmt.Println(list[0]["nick_name"].String())
	table.Update(lib.SqlIn{
		"db_code":"001",
		"table_code":"004",
	},nil,"id=?",2)
	i:=table.Delete(nil,"id=?",10)
	fmt.Println(i)
	time.Sleep(1*time.Second)
	db.Commit()
}