package main
import ("fmt" 
"strconv"
 "os")
func main(){
	num1, _:=strconv.ParseFloat(os.Args[1],64)
	num2, _:=strconv.ParseFloat(os.Args[2],64)
	fmt.Println(num1+num2)
}