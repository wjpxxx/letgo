package command

import (
	"testing")


func TestCommand(t *testing.T) {
	cmd:=New().Cd("D:\\Development\\go\\web\\src").SetCMD("notepad",".gitignore")
	//cmd.AddPipe(New().SetCMD("find","'\\c'","'80'"))
	cmd.Run()
}