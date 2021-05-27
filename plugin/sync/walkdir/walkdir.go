package walkdir

import (
	"core/log"
	"io/ioutil"
	"path/filepath"
)

//Walk 遍历目录
func Walk(dirname string,options *Options){
	if !filepath.IsAbs(dirname) {
		dirname,_=filepath.Abs(dirname)
	}
	//dirname=filepath.FromSlash(dirname)
	dir,err:=ioutil.ReadDir(dirname)
	if err!=nil{
		log.DebugPrint("walk dir error %s", err)
		return 
	}
	for _,p:=range dir{
		if p.IsDir() {
			tmpDirName:=filepath.Join(dirname,p.Name())
			Walk(tmpDirName, options)
		}else{
			fullPath:=filepath.Join(dirname,p.Name())
			if options.Callback!=nil&&filter(fullPath,options.Filter){
				options.Callback(dirname,p.Name(),fullPath)
			}
		}
	}
}
type WalkFunc func(pathName,fileName,fullName string)
//Options
type Options struct{
	Callback WalkFunc
	Filter []string //Filter files
}
//filter
func filter(fullPath string,filterArray []string)bool{
	for _,f:=range filterArray{
		if !filepath.IsAbs(f) {
			f,_=filepath.Abs(f)
		}
		r,err:=filepath.Match(f,fullPath)
		if err!=nil{
			log.DebugPrint("filter error %v",err)
			continue
		}
		if r {
			return false
		}
	}
	return true
}