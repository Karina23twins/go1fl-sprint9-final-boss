package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		log.Println("size of the slice must be greater than zero")
		return nil
	}

	src := rand.NewSource(time.Now().Unix())
	rng := rand.New(src)
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		randomNumber := rng.Int()

		slice[i] = randomNumber
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {

	if len(data) == 0 {
		log.Println("slice is empty")
		return 0
	}

	maxNumber := data[0]
	for _, num := range data {
		if num > maxNumber {
			maxNumber = num
		}
	}

	if maxNumber < 0 {
		log.Println("slice of negative numbers was generated")
		return 0
	}
	return maxNumber
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup

	var chunkSize int
	var chunks int

	sliceOfMax := make([]int, CHUNKS)

	if len(data) == 0 {
		log.Println("slice is empty")
		return 0
	}

	if len(data) == 1 {
		return maximum(data)
	}

	if len(data)%CHUNKS == 0 {
		chunkSize = len(data) / CHUNKS
		chunks = CHUNKS
	} else {
		for i := 2; i < SIZE; i++ {
			if len(data)%i == 0 {
				chunkSize = len(data) / i
				chunks = i
				break
			}
		}
	}

	for i := 0; i < chunks; i++ {
		index := i * chunkSize
		chunk := data[index:(index + chunkSize)]

		wg.Add(1)

		go func(chunk []int) {
			defer wg.Done()

			max := maximum(chunk)
			sliceOfMax = append(sliceOfMax, max)
		}(chunk)
	}
	wg.Wait()

	maxOfMax := maximum(sliceOfMax)
	return maxOfMax
}

func main() {
	if SIZE <= 0 {
		log.Println("slice size must be greater than zero")
		return
	}

	if CHUNKS <= 0 {
		log.Println("value of CHUNKS must be greater than zero")
		return
	}

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	slice := generateRandomElements(SIZE)
	if len(slice) == 0 {
		log.Println("slice is empty")
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(slice)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(slice)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
