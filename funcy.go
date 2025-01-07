package main
import (
	"os"
	"fmt"
	"strconv"
)
func do(a int, b int) int {
	s:=(a+b)*(a-b)
	return s
}
type Number struct{
	N int
	N2 int
}
func main(){
	num1,_:=strconv.Atoi(os.Args[1])
	num2,_:=strconv.Atoi(os.Args[2])
	fmt.Println(do(num1,num2))
}