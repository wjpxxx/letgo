package lib
import(
	"testing"
	"fmt"
)
func TestTime(t *testing.T){
	fmt.Println(HumpName(UnderLineName("wjpNameGee")))
	fmt.Println(Time(),1625189194-1625189188)
}