package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func operaciya(x string) string {
	// + Проверяем наличие знака плюс. Если он есть, то удаляем плюс и проверям на наличие других знаков.
	if strings.Contains(x, "+") {
		a := strings.Replace(x, "+", "", 1)
		if strings.Contains(a, "+") || strings.Contains(a, "-") || strings.Contains(a, "*") || strings.Contains(a, "/") {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		} else {
			return "+"
		}
		// - Проверяем на знак минус.
	} else if strings.Contains(x, "-") {
		b := strings.Replace(x, "-", "", 1)
		if strings.Contains(b, "+") || strings.Contains(b, "-") || strings.Contains(b, "*") || strings.Contains(b, "/") {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		} else {
			return "-"
		}
		// * Проверяем на знак умножить.
	} else if strings.Contains(x, "*") {
		c := strings.Replace(x, "*", "", 1)
		if strings.Contains(c, "+") || strings.Contains(c, "-") || strings.Contains(c, "*") || strings.Contains(c, "/") {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		} else {
			return "*"
		}
		// / Проверяем на знак делить.
	} else if strings.Contains(x, "/") {
		d := strings.Replace(x, "/", "", 1)
		if strings.Contains(d, "+") || strings.Contains(d, "-") || strings.Contains(d, "*") || strings.Contains(d, "/") {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		} else {
			return "/"
		}
	}

	panic("Выдача паники, так как строка не является математической операцией.")
}

// Переворачиваем карту
func reverseMap(originalMap map[string]int) map[int]string {
	reversedMap := make(map[int]string)
	for key, value := range originalMap {
		reversedMap[value] = key
	}
	return reversedMap
}

func main() {
	// Сканируем выражение
L00P:
	fmt.Println("Введите два числа от 1 до 10 или от I до X, и знак между ними")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	txt := sc.Text()

	// Посчитать кол-во математических операций в строке (max 1)
	// Удаляем пробелы из введеного выражения
	txt_bez := strings.ReplaceAll(txt, " ", "")
	// Если длина строки 1, то начинаем заного
	if len(txt_bez) == 1 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	// Узнаем наш знак
	znak := operaciya(txt_bez)
	// Если знак один, то ОК. Если нет, то выдаем ошибку.
	if len(znak) == 1 {

	} else if znak != "+" || znak != "-" || znak != "*" || znak != "/" {
		fmt.Println(znak)
		goto L00P
	}

	// Посчитать сколько операндов (max 2)
	// Посчитать сколько арабских цифр в строке (max 2)
	num := regexp.MustCompile("[0-9]+")
	numbers_a := num.FindAllString(txt, -1)

	// Если цифр больше двух, то выдаем ошибку
	if len(numbers_a) >= 3 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	// Посчитать сколько римских цифр в строке
	num1 := regexp.MustCompile("[A-Z]+")
	numbers_r := num1.FindAllString(txt, -1)

	// Если римских цифр больше двух, то выдаем ошибку
	if len(numbers_r) >= 3 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	// Если римских цифр и арабских больше двух, или их вообще нету, то выдаем ошибку
	if len(numbers_a)+len(numbers_r) >= 3 || len(numbers_a)+len(numbers_r) == 0 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
	// Если римских цифр и арабских по одному, то выдаем ошибку
	if len(numbers_a) == 1 && len(numbers_r) == 1 {
		panic("Используются одновременно разные системы счисления.")
	}
	// Если римских цифр и арабских в сумме равно одному, то выдаем ошибку
	if len(numbers_a)+len(numbers_r) == 1 {
		panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	//Если у нас 2 арабских числа, то делаем проверки, и считаем. (Если у нас 2 римских числа, и не было ошибки)
	if len(numbers_a) == 2 {
		// Если арабские цифры не равны от 1 до 10, то выдаем ошибку
		// Наши арабские цифры в int
		Anum1, _ := strconv.Atoi(numbers_a[0])
		Anum2, _ := strconv.Atoi(numbers_a[1])
		if Anum1 < 1 || Anum1 > 10 || Anum2 < 1 || Anum2 > 10 {
			panic("Введенные цифры не от 1 до 10")
		}
		// Если строка содержит что-то кроме 1-10 или I-X или (+, -, /, *) , то выдаем ошибку
		// Удалить из начального "выражения без пробелов" первую цифру и вторую и знак операции, если оно будет больше нуля то выдаем ошибку
		txt_bez_vsego1 := strings.Replace(txt_bez, numbers_a[0], "", 1)
		txt_bez_vsego2 := strings.Replace(txt_bez_vsego1, numbers_a[1], "", 1)
		txt_bez_vsego3 := strings.Replace(txt_bez_vsego2, znak, "", 1)
		if len(txt_bez_vsego3) > 0 {
			panic("Введены не подходящие знаки")
		}
		// Если ввели 2 арабские цифры и знак, но не в правильном порядке, то выдаем ошибку
		correct_txt := numbers_a[0] + znak + numbers_a[1]
		if txt_bez != correct_txt {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор между ними (+, -, /, *)")
		}
		// Посчитать если обе цифры арабские
		if len(numbers_a) == 2 {
			if znak == "+" {
				fmt.Println("Решение", Anum1+Anum2)
			}
			if znak == "-" {
				fmt.Println("Решение", Anum1-Anum2)
			}
			if znak == "*" {
				fmt.Println("Решение", Anum1*Anum2)
			}
			if znak == "/" {
				fmt.Println("Решение", Anum1/Anum2)
			}

		}

	}

	// Создаем мапу римских цифр и обратно
	romanMap := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
		"XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30,
		"XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40,
		"XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50,
		"LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60,
		"LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70,
		"LXX1": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80,
		"LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90,
		"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100}

	// Перевернутая мапа
	romanMapConverted := reverseMap(romanMap)

	// Нужно проверить два римских числа x +-/y и чтобы ничего лишнего

	//Если у нас 2 римских числа, то делаем проверки, и считаем. (Если у нас 2 арабских числа, и не было ошибки)
	if len(numbers_r) == 2 {
		Rnum1 := numbers_r[0] // Первая римская в строке
		Rnum2 := numbers_r[1] // Вторая римская в строке
		// Конвертируем в int
		convert1 := (romanMap[Rnum1]) // Первая римская в int
		convert2 := (romanMap[Rnum2]) // Вторая римская в int
		// Проверить правильность введеных римских цифр, от 1 до 10
		if convert1 < 1 || convert1 > 10 || convert2 < 1 || convert2 > 10 {
			panic("Введенные римские цифры не от I до X")
		}
		// Если ввели 2 римские цифры и знак, но не в правильном порядке, то выдаем ошибку
		correct_txtR := numbers_r[0] + znak + numbers_r[1]
		if txt_bez != correct_txtR {
			panic("Формат математической операции не удовлетворяет заданию — два операнда и один оператор между ними (+, -, /, *)")
		}
		// Посчитать если обе цифры римские и результат сконвертировать в римское число
		if len(numbers_r) == 2 {
			if znak == "+" {
				result := (romanMapConverted[convert1+convert2])
				fmt.Println("Решение", result)
			}
			if znak == "-" {
				result := (romanMapConverted[convert1-convert2])
				fmt.Println("Решение", result)
				// Если результат меньше единицы, то выдаем ошибку
				if result < "1" {
					panic("В римской системе нет отрицательных чисел и нуля")
				}
			}
			if znak == "*" {
				result := (romanMapConverted[convert1*convert2])
				fmt.Println("Решение", result)
			}
			if znak == "/" {
				result := (romanMapConverted[convert1/convert2])
				fmt.Println("Решение", result)
				if result < "1" {
					panic("В римской системе нет отрицательных чисел и нуля")
				}

			}

		}

	}
	goto L00P

}
