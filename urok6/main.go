/*              UKAZATELI
var y *int 
* -- oboznachenie ukazatel
var m = new(int) --- eto tozhe ukazatel 
var x int = 5
b := &x --znachenie budet adres 
fmt.println(b) ---- budet adres
fmt.println(*b) ------ 5
*/
package main
import "fmt"
func main (){
	var x int = 10
	var ptr *int 
	ptr = &x
	fmt.Println(x)  // vivod znachenie "x"
	fmt.Println(&x) // adres "x"
	fmt.Println(ptr) // adres "ptr"
	fmt.Println(*ptr) // vivod znachenie "ptr"
	
	n := 5
	fmt.Println("before:", n)
	updateVallue(&n)
	fmt.Println("after:", n)

	s := 3
	d := 5
	fmt.Println("before:", s, d)
	swap(&s, &d)
	fmt.Println("after:", s, d)


	
}
// Напишите функцию updateValue, которая принимает указатель на целое число и увеличивает его значение на 10.
func updateVallue(p *int) { 
	*p = *p + 10
}

// Напишите функцию swap, которая меняет местами значения двух переменных с использованием указателей.
func swap (l, m *int) {
	*l, *m = *m, *l
}

// Напишите программу для управления персональными данными. Реализуйте функции для хранения и вывода имени и возраста человека с использованием указателей.
// func storename( namePtr *string, name string) {
// 	*namePtr = name
// }


 // Напишите программу для учета голосов на выборах. Реализуйте функции для подсчета голосов за каждого кандидата и определения победителя.
func vote (candidator *string, votesPtr *int) {
	*votesPtr++ 
}
func winner(candidate1VotesPtr *int, candidate2VotesPtr *int ) string {
	if *candidate1VotesPtr > *candidate2VotesPtr {
		return "Кандидат 1"
	} else if *candidate1VotesPtr < *candidate2VotesPtr {
		return "Кандидат 2"
	} else {
		return "Ничья"
	}
}

// Напишите программу для учета голосов на выборах. Реализуйте функции для подсчета голосов за каждого кандидата и определения победителя.
func addExpense(totalPtr *float64, amount float64) {
	*totalPtr += amount
}
func printTotalExpenses(totalPtr *float64) {
	fmt.Printf("Общая сумма расходов : %.2f\n ", *totalPtr)
}

