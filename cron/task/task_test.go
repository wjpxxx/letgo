package task
import (
	"testing"
	"time"
	"fmt"
)
func TestTask(t *testing.T) {
	RegisterTask("add",3,func(ctx *Context){
		fmt.Println(ctx.Name,ctx.TaskNo)
		//fmt.Println()
		//time.Sleep(1*time.Second)
	})
	go func(){
		time.Sleep(105*time.Second)
		Stop()
	}()
	StartAndWait()
}