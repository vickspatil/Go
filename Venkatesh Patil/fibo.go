package main
import (
	"os"
	"fmt"
	"strconv"
)
func main(){
	num1,_:=strconv.Atoi(os.Args[1])
	j:=1
	x:=0
	s:=0
	for i:=0;i<num1;i++ {
		s=x+j
		x=j
		j=s
		fmt.Println(x)
	}
	
}