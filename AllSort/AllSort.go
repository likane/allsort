package allsort

import (
	"fmt"
	"math/rand"
	"time"
)

// create goroutines to sort an array or slice
// create a func that will run all goroutines in parallel and return once the first routine is done
//	todo
//  func all
//  func quick --
//  func bubble --
//  func binary
//  func merge --
//  func insertion --
//  func selection
//  func counting
//  func bogo
//  func timsort
//  func cocktail shaker sort
//  func bitonic sort
//  func block
//  func cubesort
//  func library
//

/*
Todo
1) add more sort functions
2) test as a module w/o hard coded slice
3) remove comments
*/

func AllSort(sortme []int) []int {
	if len(sortme) < 1 {
		return sortme
	}
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
