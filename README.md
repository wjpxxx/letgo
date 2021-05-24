letgo is an open-source, high-performance web framework for the Go programming language.

# Model operation

```go
package main

import "github.com/wjpxxx/letgo/db/mysql"
func main() {
model:=NewModel("dbname","tablename")
        model.Fields("*").
                Alias("m").
                Join("sys_shopee_shop as s").
                On("m.id","s.master_id").
                OrOn("s.master_id",1).
                AndOnRaw("m.id=1 or m.id=2").
                LeftJoin("sys_lazada_shop as l").
                On("m.id", "l.master_id").
                WhereRaw("m.id=1").
                AndWhere("m.id",2).
                OrWhereIn("m.id",lib.Int64ArrayToInterfaceArray(ids)).
                GroupBy("m.id").
                Having("m.id",1).
                AndHaving("m.id",1).
                OrderBy("m.id desc").Find()
        fmt.Println(model.GetLastSql())
}
```