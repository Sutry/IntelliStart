// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
// Ми маємо 23 грн.
// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?
//
// Задача:
// Опишіть вирішення всіх пунктів задачі використовуючи необхідні змінні чи/та константи.
// Під опишіть, я маю на увазі наступне:
// Я маю збілдити ваш код і запустити програму. В результаті цього, я маю побачити роздруковані
// відповіді на поставлені вище запитання. Перед відповідю треба роздрукувати саме питання.
// Правильно обирайте типи даних та назви змінних чи констант.
// Публікація:
// Створити папку в своєму репозиторії в github та завантажити туди main.go файл, в котрому буде зроблено дане завдання.

package main

// We use all prices in kopiikas (to be an integer)
// 1 hryvnia is 100 kopiikas - https://bank.gov.ua/en/uah/obig-coin#10_2016
// but in Wikipwdia kopek or kopeck. I used the term of the National Bank of Ukraine
const (
	priceApple int = 599  // Ціна 1 яблука, копійок
	pricePear  int = 700  // Ціна 1 груші, копійок
	money      int = 2300 // Наші гроші, в копійках
)

// main - main func
func main() {
	// 1 пункт завдання
	var (
		amountApple int = 9
		amountPear  int = 8
	)
	println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	println("\tЩоб купити 9 яблук нам треба мати:", amountApple*priceApple/100, "грн")
	println("\tЩоб купити 8 груш нам треба мати:", amountPear*pricePear/100, "грн")
	println("\tщоб купити 9 яблук та 8 груш разом:", (amountApple*priceApple+amountPear*pricePear)/100, "грн")
	// 2 пункт завдання
	println("\n2. Скільки груш ми можемо купити?")
	println("\tМи можемо купити", money/pricePear, "груш(и)")
	// 3 пункт завдання
	println("\n3. Скільки яблук ми можемо купити?")
	println("\tМи можемо купити", money/priceApple, "яблук(а)")
	// 4 пункт завдання
	amountApple = 2
	amountPear = 2
	println("\n4. Чи ми можемо купити 2 груші та 2 яблука?")
	if priceApple*amountApple+pricePear*amountPear <= money {
		println("\tТак, зможемо :)")
	} else {
		println("\tНі, не зможемо :(")
	}
	// println("\n4. Чи ми можемо купити 2 груші та 2 яблука?", priceApple*amountApple+pricePear*amountPear <= money)
}
