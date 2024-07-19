//                  Sikli
/* for init; cond; operations {
body 
}      != -- oznachaet neravenstva
for i := 0; i <= 9; i ++ {
println(i)
}
init := 0
for ; cond; operations {
body }
*/
package main
import "fmt"
func main () {
// 	var i int
// 	for i := 0; i <= 9; i ++ {
// 		fmt.Println(i)
// }
// fmt.Println(i)

var num = 546
for num > 0 {
	a := num % 10
	fmt.Println(a)
	num = num / 10
}
// n/= 10 -- eto oznachaet n= n/ 10
for n:= 654; n > 0; n/= 10 {
a := n % 10
fmt.Println(a)
}

/* beskonechnie sikl
for {
	fmt.Println("hi")
}
*/

for m := 0; m <= 100; m++ {
if m % 3 == 0 {
	continue
}
fmt.Println(m)
}

//             vlozhenie sikli -- eto sikli vnutri sikla 
/* %d --- это способ указать, что в строку нужно вставить целое число. 
%s — это спецификатор формата для строки. Он указывает, что на его месте будет выведена строковая переменная (name).
%d — это спецификатор формата для целого числа. Он указывает, что на его месте будет выведена целочисленная переменная (age).
\n — это символ новой строки, который добавляется в конце строки.
 printf --- форматирует строку и выводит её на экран.
 %.2f -- okruglaet chislo */

// for l := 1; l < 10; l++ {
// 	for k := 1; k < 10; k++ {
// 		fmt.Printf("%d * %d = %d\n", l, k, l * k)
// 	} 
// 	fmt. Println("----------------")
// }

var n1 int 
n1 = 5
sum := 0.0
for i := 1; i < n1; i++ {
sum = sum + 1 / float64(i)
fmt.Println(sum)
}

var p float64 = 1
var n float64
for i := 1; i < 10; i++ {
	fmt.Scan(&n)
	p = p * n
}
fmt.Println(p)

}


/*
cd ./.. or cd .. - выход с папки
cd ./name_of_folder - вход в папку
*/