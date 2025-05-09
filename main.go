package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// генерирует числа и записывает в новый слайс
func generateRandomElements(size int) []int {

	//проверка на размер слайса
	if size <= 0 || size > SIZE {
		fmt.Println("\nразмер слайса должен быть больше 0 и не превышать", SIZE)
		return []int{}
	}

	//проверка на отрицательное число
	if size < 0 {
		fmt.Println("\nчисло не может быть отрицательным")
		return []int{}
	}

	//генерация чисел
	randNumber := rand.New(rand.NewSource(time.Now().UnixNano()))

	//создаем слайс
	newSlice := make([]int, size)

	for i := 0; i < size; i++ {
		newSlice[i] = randNumber.Intn(SIZE)
	}

	return newSlice

}

// находим максимальное число в слайсе
func maximum(data []int) int {
	//проверка на пустоту слайса
	if len(data) == 0 {
		fmt.Println("\nСлайс не может быть пустым")
		return 0
	}

	if len(data) == 1 {
		fmt.Println("\nСлайс содержит 1 элемент")
		return data[0]
	}

	//проверка на отрицательные числа
	negative := false
	for _, n := range data {
		if n < 0 {
			negative = true
			break
		}
	}
	if negative {
		fmt.Println("\nСлайс содержит отрицательные числа")
	}

	//проверка на одинаковые числа в слайсе
	allNumber := true
	for i := 1; i < len(data); i++ {
		if data[i] != data[i-1] {
			allNumber = false
			break
		}
	}
	if allNumber {
		fmt.Println("\nВсе числа в слайсе одинаковые")
		return data[0]
	}

	//для поиска максимального числа начинаем с первого элемента
	maxNum := data[0]

	//поиск максимального числа
	for _, n := range data {
		if n > maxNum {
			maxNum = n
		}
	}

	return maxNum

}

// возращает максимальное значение
// делим слайс на 8 частей, находим в каждом максимум
// создать 8 горутин
// создать новый слайс для хранения 8 макс чисел
func maxChunks(data []int) int {
	//проверка на пустой слайс
	if len(data) == 0 {
		fmt.Println("\nСлайс не может быть пустым")
		return 0
	}

	//проверка на 1 элемент
	if len(data) == 1 {
		fmt.Println("\nСлайс содержит 1 элемент")
		return data[0]
	}

	//проверка на отрицательные числа
	negative := false
	for _, n := range data {
		if n < 0 {
			negative = true
			break
		}
	}
	if negative {
		fmt.Println("\nСлайс содержит отрицательное число")
	}

	//проверка на одинаковые числа в слайсе
	allNumber := true
	for i := 1; i < len(data); i++ {
		if data[i] != data[i-1] {
			allNumber = false
			break
		}
	}
	if allNumber {
		fmt.Println("\nВсе числа в слайсе одинаковые")
		return data[0]
	}

	//размер слайса
	newSlice := len(data) / CHUNKS

	//слайс для хранения максимумов
	maxNumber := make([]int, CHUNKS)

	var wg sync.WaitGroup

	//запуск 8 горутин
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			firstIndex := i * newSlice
			lastIndex := firstIndex + newSlice
			if i == CHUNKS-1 {
				lastIndex = len(data)
			}

			max := data[firstIndex]

			for _, v := range data[firstIndex:lastIndex] {
				if v > max {
					max = v
				}
			}

			maxNumber[i] = max
		}(i)
	}

	wg.Wait()

	result := maxNumber[0]
	for _, m := range maxNumber[1:] {
		if m > result {
			result = m
		}
	}

	return result

}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	generate := generateRandomElements(SIZE)
	fmt.Println()

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(generate)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	fmt.Println()
	start = time.Now()
	max = maxChunks(generate)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
