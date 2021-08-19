package model

import (
    "github.com/wjpxxx/letgo/db/mysql"
    "github.com/wjpxxx/letgo/lib"
	"hyinx/model/entity"
)

//AreaModel
type AreaModel struct{
    mysql.Model
}
//GetAreaModel 获得操作模型
func GetAreaModel() *AreaModel{
    model:=&AreaModel{}
    model.Init("bdsy","cd_area")
    //开启软删除
    model.SoftDelete=true
    return model
}
//SaveByEntity
func (m *AreaModel)SaveByEntity(data AreaEntity) int64{
    var inData lib.SqlIn
    lib.StringToObject(data.String(), &inData)
    if data.ID>0{
        delete(inData,"create_time")
        m.Where("id", data.ID).Update(inData)
        return data.ID
    }else{
        delete(inData,"id")
        delete(inData,"update_time")
        return m.Insert(inData)
    }
}
//GetEntityById 通过id获得数据
func (m *AmazonShopModel) GetEntityById(id int64) AreaEntity{
    var out AreaEntity
    data:= m.Where("id", id).Find()
    data.Bind(&out)
    return out
}