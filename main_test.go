package main

import (
	"fmt"
	"testing"
)

//тест для функции generateRandomElements

func TestGenerateRandomElement(t *testing.T) {

	//тестовые данные
	size1 := 100_000_001
	size2 := -10
	size3 := 0

	//проверка на 0
	fmt.Printf("Тест №1: если аргумент равен %d, ожидаем пустой слайс\n", size3)
	data := generateRandomElements(size3)
	if len(data) != 0 {
		t.Errorf("Ожидаемый результат: 0, получен размер %d", len(data))
	}

	//проверка на отрицательное число и
	fmt.Printf("Тест №2: если аргумент является отрицательным числом, ожидаем пустой слайс\n")
	data = generateRandomElements(size2)
	if len(data) != 0 {
		t.Errorf("Ожидаемый результат: 0, Фактический результат: %d", len(data))
	}

	//проверка на выход из диапазона размера
	fmt.Printf("Тест №3: если аргумент больше 100_000_000, ожидаем пустой слайс")
	data = generateRandomElements(size1)
	if len(data) > 0 {
		t.Errorf("Ожидается пустой слайс при выходе из диапазона размера, получено %d эелементов", size1)
	}

}

func TestMaximum(t *testing.T) {

	//тестовые данные
	data1 := []int{}
	data2 := []int{55}
	data3 := []int{10, -10}
	data4 := []int{10, 10, 10}

	//проверка на пустой слайс
	fmt.Println("ТЕСТ №1: если передан пустой слайс, ожидаем результат 0")
	result := maximum(data1)
	if result != 0 {
		t.Errorf("Ожидаемый результат: 0, Фактический результат: %d", result)
	}

	//проверка на 1 элемент в слайсе
	fmt.Println("ТЕСТ №2: если в слайсе только 1 элемент, ожидаем вывод этого элемента")
	result = maximum(data2)
	if result != 55 {
		t.Errorf("Ожидаемый результат: %d, Фактический результат: %d", data2, result)
	}

	//проверка на отрицательное число
	fmt.Println("ТЕСТ №3: проверка на отрицательное число")
	result = maximum(data3)
	if result != 10 {
		t.Errorf("Ожидаемый результат: 10, Фактический реузльтат: %d", result)
	}

	//проверка на одинаковые числа в слайсе
	fmt.Println("ТЕСТ №4: если в слайсе одинаковые числа, выводим число")
	result = maximum(data4)
	if result != 10 {
		t.Errorf("Ожидаемый результат 10, Фактический результат: %d", result)
	}

}

func TestMaximumChunks(t *testing.T) {
	data1 := []int{}
	data2 := []int{55}
	data3 := []int{10, -10}
	data4 := []int{10, 10, 10}

	//проверка на пустой слайс
	fmt.Println("ТЕСТ №1: если передан пустой слайс, ожидаем результат 0")
	result := maxChunks(data1)
	if result != 0 {
		t.Errorf("Ожидаемый результат: 0, Фактический результат: %d", result)
	}

	//проверка на 1 элемент в слайсе
	fmt.Println("ТЕСТ №2: если в слайсе только 1 элемент, ожидаем вывод этого элемента")
	result = maxChunks(data2)
	if result != 55 {
		t.Errorf("Ожидаемый результат: %d, Фактический результат: %d", data2, result)
	}

	//проверка на отрицательное число
	fmt.Println("ТЕСТ №3: проверка на отрицательное число")
	result = maxChunks(data3)
	if result != 10 {
		t.Errorf("Ожидаемый результат: 10, Фактический реузльтат: %d", result)
	}

	//проверка на одинаковые числа в слайсе
	fmt.Println("ТЕСТ №4: если в слайсе одинаковые числа, выводим число")
	result = maxChunks(data4)
	if result != 10 {
		t.Errorf("Ожидаемый результат 10, Фактический результат: %d", result)
	}

}
