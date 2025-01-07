package main
import ("fmt")
func add (a string, b string) string {
	return a+" "+b
}
type Car struct {
	Name string
	Make string
}
func main(){
	a:=Car{"Audi","A4"}
	b:=Car{"BMW","X7"}
	fmt.Println(a,b)
	fmt.Println(add(a.Name,b.Name))
}