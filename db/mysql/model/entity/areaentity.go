package entity

import (
    "fmt"
    "github.com/wjpxxx/letgo/lib"
)

//AreaEntity
type AreaEntity struct{
    AreaId    int  `json:"area_id"`
    Title    string  `json:"title"`
    Pid    int  `json:"pid"`
    Initials    string  `json:"initials"`
    Sort    int  `json:"sort"`
    IsCommoned    int  `json:"is_commoned"`

}
//String
func (e *AreaEntity)String()string{
    return lib.ObjectToString(e)
}
