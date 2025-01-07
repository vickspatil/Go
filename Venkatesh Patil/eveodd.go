package main
import (
	"os"
	"fmt"
	"strconv"
)
func main(){
	if len(os.Args) < 2 {
		fmt.Println("Poora likh na")
		return 
	}
	num1,_:=strconv.Atoi(os.Args[1])
	if num1%2!=0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}
}