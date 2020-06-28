package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//sortme := []int{60, 56, 38, 42, 50439, 87, 31, 5, 20005, 54, 3, 10, 88, 6987, 9875, 1234569, 97, 95, 6, 13, 52, 12, 20005, 54, 3, 10, 88, 97, 5003, 98, 10, 90, 12, 42, 46, 3053, 87, 29, 22, 98, 8, 66, 60, 56, 38, 42, 87, 31, 5, 54, 10, 88, 97, 95, 6, 13, 52, 12, 98, 10, 90, 12, 42, 46, 87, 29, 22, 98, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 5, 54, 10, 88, 97, 95, 6, 13, 52, 12, 98, 10, 90, 12, 42, 46, 87, 29, 22, 98, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 83, 38, 34, 77, 28, 95, 6, 13, 52, 12, 98, 10, 90, 12, 42, 46, 87, 29, 22, 98, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 5, 54, 10, 88, 97, 95, 6, 13, 52, 12, 98, 10, 90, 12, 42, 46, 87, 29, 22, 98, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 83, 46, 28, 68, 11, 44, 43, 95, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 83, 38, 34, 77, 28, 46, 28, 68, 11, 44, 43, 95, 43, 22, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 86, 72, 53, 72, 53, 71, 29, 81, 82, 52, 85, 88, 12, 47, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 83, 38, 34, 2, 77, 28, 46, 28, 68, 11, 44, 43, 95, 17, 4, 20, 53, 65, 10, 98, 58, 42, 11, 1, 12, 6, 82, 5}
	sortme := []int{95, 17, 1, 12, 252543, 6, 82, 5, 50449, 3334333, 5, 21, 54, 645, 140, 948, 58, 42, 60, 56, 38, 42, 50439, 3333333, 4, 20, 53, 65, 10, 98, 58, 42, 11, 87, 31, 5, 20005, 54, 3, 10, 88, 6987, 9875, 1234569, 97, 95, 6, 13, 52, 12, 20005, 54, 3, 10, 88, 97, 5003, 98, 10, 90, 12, 42, 46, 3053, 87, 29, 22, 98, 8, 66, 60, 56, 38, 42, 87, 31, 5, 54, 10, 88, 97, 95, 6, 13, 52, 12, 98, 10, 90, 12, 42, 46, 87, 29, 22, 98, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 5, 54, 10, 88, 97, 95, 6, 13, 52, 12, 98, 10, 90, 12, 42, 46, 87, 29, 22, 98, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 83, 38, 34, 77, 28, 95, 6, 13, 52, 12, 98, 10, 90, 12, 42, 46, 87, 29, 22, 98, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 5, 54, 10, 88, 97, 95, 6, 13, 52, 12, 98, 10, 90, 12, 42, 46, 87, 29, 22, 98, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 83, 46, 28, 68, 11, 44, 43, 95, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 83, 38, 34, 77, 28, 46, 28, 68, 11, 44, 43, 95, 43, 22, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 86, 72, 53, 72, 53, 71, 29, 81, 82, 52, 85, 88, 12, 47, 8, 66, 64, 12, 66, 17, 57, 40, 46, 76, 16, 16, 43, 43, 22, 80, 50, 45, 55, 48, 24, 91, 78, 75, 47, 95, 28, 55, 12, 57, 83, 38, 34, 2, 77, 28, 46, 28, 68, 11, 44, 43}
	all(sortme)
}

func all(sortme []int) []int {
	c1 := make(chan []int)
	c2 := make(chan []int)
	c3 := make(chan []int)
	c4 := make(chan []int)
	start := time.Now()
	timeMap := make(map[string]int64, 4)

	go doQuick(sortme, start, timeMap, c1)
	go doBubble(sortme, start, timeMap, c2)
	go doMerge(sortme, start, timeMap, c3)
	go doInsertion(sortme, start, timeMap, c4)

	quickSorted := <-c1
	bubbleSorted := <-c2
	mergeSorted := <-c3
	insertionSorted := <-c4

	if quickSorted != nil {
		fmt.Println(quickSorted)
		fmt.Println(timeMap)
		return quickSorted
	} else if bubbleSorted != nil {
		fmt.Println(bubbleSorted)
		fmt.Println(timeMap)
		return bubbleSorted
	} else if mergeSorted != nil {
		fmt.Println(mergeSorted)
		fmt.Println(timeMap)
		return mergeSorted
	} else if insertionSorted != nil {
		fmt.Println(insertionSorted)
		fmt.Println(timeMap)
		return insertionSorted
	} else {
		fmt.Println("Error: Unsorted")
		fmt.Println(timeMap)
		return sortme
	}
}

func timeMeasure(t time.Time, functionName string, timeMap map[string]int64) {
	// var mutex = &sync.Mutex{}
	// mutex.Lock()
	elapsed := time.Since(t).Nanoseconds()
	// mutex.Unlock()
	timeMap[functionName] = elapsed

	//timeMap.Store(functionName, elapsed)
	//timeMap.Set(functionName, elapsed)
	//fmt.Println(timeMap)
}

func doQuick(sortme []int, startTime time.Time, timeMap map[string]int64, c1 chan []int) {
	// var mutex = &sync.Mutex{}
	// mutex.Lock()
	defer timeMeasure(time.Now(), "Quick Sort", timeMap)
	// mutex.Unlock()
	quicksort(sortme)
	c1 <- sortme
}

func quicksort(sortme []int) []int {

	if len(sortme) < 2 {
		return sortme
	}

	left, right := 0, len(sortme)-1

	pivot := rand.Int() % len(sortme)

	sortme[pivot], sortme[right] = sortme[right], sortme[pivot]

	for i := range sortme {
		if sortme[i] < sortme[right] {
			sortme[left], sortme[i] = sortme[i], sortme[left]
			left++
		}
	}

	sortme[left], sortme[right] = sortme[right], sortme[left]

	quicksort(sortme[:left])
	quicksort(sortme[left+1:])

	return sortme
}

func doBubble(sortme []int, startTime time.Time, timeMap map[string]int64, c2 chan []int) {
	// var mutex = &sync.Mutex{}
	// mutex.Lock()
	defer timeMeasure(time.Now(), "Bubble Sort", timeMap)
	// mutex.Unlock()
	bubblesort(sortme)
	c2 <- sortme
}

func bubblesort(sortme []int) []int {
	for i := len(sortme); i > 0; i-- {
		for j := 1; j < i; j++ {
			if sortme[j-1] > sortme[j] {
				intermediate := sortme[j]
				sortme[j] = sortme[j-1]
				sortme[j-1] = intermediate
			}
		}
	}
	return sortme
}

func doMerge(sortme []int, startTime time.Time, timeMap map[string]int64, c3 chan []int) {
	// var mutex = &sync.Mutex{}
	// mutex.Lock()
	defer timeMeasure(time.Now(), "Merge Sort", timeMap)
	// mutex.Unlock()
	mergeSort(sortme)
	c3 <- sortme
}

func mergeSort(sortme []int) []int {

	var num = len(sortme)

	if num == 1 {

		return sortme
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = sortme[i]
		} else {
			right[i-middle] = sortme[i]
		}
	}
	var sorted []int = merge(mergeSort(left), mergeSort(right))
	return sorted
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}

func doInsertion(sortme []int, startTime time.Time, timeMap map[string]int64, c4 chan []int) {
	// var mutex = &sync.Mutex{}
	// mutex.Lock()
	defer timeMeasure(time.Now(), "Insertion Sort", timeMap)
	// mutex.Unlock()
	insertionSort(sortme)
	c4 <- sortme
}

func insertionSort(sortme []int) {
	var n = len(sortme)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if sortme[j-1] > sortme[j] {
				sortme[j-1], sortme[j] = sortme[j], sortme[j-1]
			}
			j = j - 1
		}
	}

}
