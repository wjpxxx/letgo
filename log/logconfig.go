package log

import (
	"core/lib"
	"os"
)

type LogConfig struct{
	Debug bool `json:"debug"`
	Writer string `json:"writer"`
	LogFilePath string `json:"logFilePath"`
	File *os.File
}

func (l LogConfig)String()string{
	return lib.ObjectToString(l)
}