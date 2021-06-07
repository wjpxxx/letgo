package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//DropoffEntity
type DropoffEntity struct{
	BranchList []BranchEntity `json:"branch_list"`
}

//String
func(d DropoffEntity)String()string{
	return lib.ObjectToString(d)
}
