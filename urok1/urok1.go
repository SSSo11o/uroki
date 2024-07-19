/* package main
// import "fmt"    // fmt -- eto paket formata
// func main ()  {

// 	fmt.Println("hello")
*/ 
             // Peremenie -- eto echeka pameti kotorie imet v sebe znachenie
			 // int, int8, int16, int32, int64
			 // int8   -128 do 127
			 // int16   -256 do 255
			 // uint64: представляет целое число от 0 до 18 446 744 073 709 551 615
             // byte: синоним типа uint8, представляет целое число от 0 до 255 и занимает 
			 // rune: синоним типа int32,
			 /* be znakovie peremenie
			 uint, uint8, uint16, uint32, uint64  -- dly polozhitelnix chisel
			 */  
			 //float32, float64  -- eto ne seloe chislo
			 // bool -- eto True ili False
			 // complex 

/*  package main
import "fmt"
func main (){
	var num int 
	num = 8   // var -- eto peremenae i znachenie minaetsa
	fmt.Println(num)
	var num2 int = 6
	fmt.Println(num2)
	num3  := 10
	fmt.Println(num3)
*/

package main
import "fmt"
func main (){
	var num int 
	num = 4   
	var num2 int8 = 6
	num3 := int8(num) + num2
	fmt.Println(num3)
} 

