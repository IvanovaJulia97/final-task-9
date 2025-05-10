package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

//тест для функции generateRandomElements

func TestGenerateRandomElement(t *testing.T) {

	//тестовые данные
	test := []struct {
		name     string //название теста
		size     int    //элементы для проверки
		expected int    // ожидаемое значение
	}{
		{"Тест №1: проверка на 0", 0, 0},
		{"Тест №2: проверка на выход из диапазона размера", 100_000_001, 0},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			fmt.Println(v.name)
			data := generateRandomElements(v.size)
			require.Len(t, data, v.expected)
		})
	}

}

func TestMaximum(t *testing.T) {

	//тестовые данные
	test := []struct {
		name     string //название теста
		elements []int  //элементы для проверки
		expected int    //ожидаемый результат
	}{
		{"Тест №1: проверка на пустой слайс", []int{}, 0},
		{"Тест №2: проверка на один элемент в слайсе", []int{55}, 55},
		{"Тест №3: проверка на поиск максимального числа", []int{1, 5, 10, 55}, 55},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			fmt.Println(v.name)
			res := maximum(v.elements)
			require.Equal(t, v.expected, res)
		})
	}

}

func TestMaximumChunks(t *testing.T) {
	//тестовые данные
	test := []struct {
		name     string //название теста
		elements []int  //элементы для проверки
		expected int    //ожидаемый результат
	}{
		{"Тест №1: проверка на пустой слайс", []int{}, 0},
		{"Тест №2: проверка на один элемент в слайсе", []int{55}, 55},
		{"Тест №3: проверка на поиск максимального числа", []int{1, 5, 10, 55}, 55},
	}

	for _, v := range test {
		t.Run(v.name, func(t *testing.T) {
			fmt.Println(v.name)
			res := maxChunks(v.elements)
			require.Equal(t, v.expected, res)
		})
	}

}
