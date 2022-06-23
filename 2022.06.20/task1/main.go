/*	Задача 1
Прибрати всі дублікати з слайсу int.
Приклад даних на вхід: [3, 4, 4, 3, 6, 3]
видаляємо 3 по індексу 0
видаляємо 4 по індексу 1
видаляємо 3 по індексу 3
Правильний результат: [3, 4, 6]
Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
func main(){
arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
var result []int
//...
// тут має бути ваш код
// змінна result в кінці функції має тримати слайс з вже видаленими дублікатами відповідно до правил
}
*/

package main

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8, 4, 3}
	var result []int

	result = make([]int, 0, len(arr))
	m := make(map[int]bool)

	for _, val := range arr {
		if _, ok := m[val]; !ok {
			m[val] = true
			result = append(result, val)
		}
	}

	for _, val := range result {
		println(val)
	}
}
