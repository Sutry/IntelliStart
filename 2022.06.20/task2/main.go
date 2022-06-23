/*	Задача 2
На вхід подано стрінг з цілими числа, котри розділені пробілами.
Треба повернути найбільше та найменше число.
Наприклад:
input := "1 2 3 4 5" // повертає "5 1"
input := "1 9 3 4 -5" // повертає "9 -5"
Уточнення:
1. Всі числа є не більше, ніж int32. Використовуйте цей тип даних.
2. В стрінгі завжди буде принаймні одне число.
3. Результатом має бути стрінг, в якому два числа розділені пробілом (або одне, якщо дано було лише одне
число). Найбільше число має бути першим.
4. Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
func main(){
input := "1 9 3 4 -5"
var result string
//...
// тут має бути ваш код
// змінна result в кінці функції має тримати стрінг з правильними результатами, згідно до умови задачі
}
*/

package main

import (
	"strconv"
	"strings"
)

// For obtain string with max & min values (or one only) from variable 'input'
func main() {

	var input string = "1 2 3 4 5"
	arrInput := strings.Split(input, " ")

	var (
		iInt   int
		iInt32 int32
		max    int32
		min    int32
		result string
	)

	iInt, _ = strconv.Atoi(arrInput[0])
	min = int32(iInt)
	max = min

	for i := 0; i < len(arrInput); i++ {
		iInt, _ = strconv.Atoi(arrInput[i])
		iInt32 = int32(iInt)
		println(iInt, " ", iInt32)
		//if arrInput[i] > max {
		//			max = arrInput[i]
		if iInt32 > max {
			max = iInt32
		}
		//		if arrInput[i] < min {
		//			min = arrInput[i]
		if iInt32 < min {
			min = iInt32
		}
	}
	if min == max {
		result = strconv.Itoa(int(min))
	} else {
		result = strconv.Itoa(int(max)) + " " + strconv.Itoa(int(min)) //max + " " + min
	}
	println(result)
}
