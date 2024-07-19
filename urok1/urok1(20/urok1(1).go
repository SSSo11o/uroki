//   Constant  -- znachenie ne izminaetsa

package main
import "fmt"
const (
	n1 = iota + 4    // iota -- iota   samo daet znachenie 
	n2  
	n3
	n5

)
func main (){
	// const num  = 7
	// const num2 = 14 
	// const num3 = 23
	const (
		num  =7
		num1 = 1 
		num4 = 2
	)
	fmt.Println(n5)
}