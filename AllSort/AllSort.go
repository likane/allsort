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
//  func quick
//  func bubble
//  func binary
//  func merge
//  func insertion
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

// func all(a []int) ([]int, float64) {
// 	sortMe(a)
// }

func all(a []int) {
	sortMe(a)
}

func sortMe(unsorted []int) {
	//var sorted [len(a)]int
	//c := make(chan []int)
	//start := time.Now()
	go quicksort(unsorted)
	go bubblesort(unsorted)
	go mergeSort(unsorted)
	// fmt.Println(go quicksort(unsorted))
	// fmt.Println(go bubblesort(unsorted))
	// fmt.Println(go mergeSort(unsorted))
	//go insertionsort(unsorted)
	//sorted <- c
	//var timeLapsed float64 = time.Since(start).Seconds()
	//return sorted, timeLapsed

}

// ([]int, float64)
//https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html
func quicksort(a []int) {
	var timeLapsed float64
	start := time.Now()
	if len(a) < 2 {
		timeLapsed = time.Since(start).Seconds()
		fmt.Println(a, timeLapsed)
		//return a, timeLapsed
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])
	timeLapsed = time.Since(start).Seconds()
	fmt.Println(a, timeLapsed)
	//return a, timeLapsed
}

// https://www.golangprograms.com/golang-program-for-implementation-of-bubble-sort.html
// https://medium.com/@houzier.saurav/merge-sort-golang-14a5e4f04f00#:~:text=Bubble%20Sort%20is%20the%20simplest,and%20swaps%20since%205%20%3E%201.
func bubblesort(numbers []int) {
	var timeLapsed float64
	start := time.Now()
	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				intermediate := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = intermediate
			}
		}
	}
	timeLapsed = time.Since(start).Seconds()
	fmt.Println(numbers, timeLapsed)
	//return numbers, timeLapsed
}

func mergeAlgorithm(a []int) {
	var timeLapsed float64
	start := time.Now()
	var sorted = mergeSort(a)
	timeLapsed = time.Since(start).Seconds()
	fmt.Println(sorted, timeLapsed)
	//fmt.Println(sorted, timeLapsed)
	//return sorted, timeLapsed
}

func mergeSort(items []int) []int {

	var num = len(items)

	if num == 1 {
		//timeLapsed = time.Since(start).Seconds()
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}
	var sorted []int = merge(mergeSort(left), mergeSort(right))
	//timeLapsed = time.Since(start).Seconds()
	return sorted
}

// https://www.golangprograms.com/golang-program-for-implementation-of-mergesort.html
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

// https://www.golangprograms.com/golang-program-for-implementation-of-insertionsort.html
func insertionsort(items []int)  {
	start := time.Now()
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	var timeLapsed float64 = time.Since(start).Seconds()
}



package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	sortme := []int{60,56,38,42,87,31,5,54,10,88,97,95,6,13,52,12,98,10,90,12,42,46,87,29,22,98,8,66,64,12,66,17,57,40,46,76,16,16,43,43,22,80,5,54,10,88,97,95,6,13,52,12,98,10,90,12,42,46,87,29,22,98,8,66,64,12,66,17,57,40,46,76,16,16,43,43,22,80,50,45,55,48,24,91,78,75,47,95,28,55,12,57,83,38,34,77,28,46,28,68,11,44,43,95,50,45,55,48,24,91,78,75,47,95,28,55,12,57,83,38,34,77,28,46,28,68,11,44,43,95,43,22,57,40,46,76,16,16,43,43,22,80,50,45,55,48,24,91,78,75,47,95,28,55,12,57,86,72,53,72,53,71,29,81,82,52,85,88,12,47,8,66,64,12,66,17,57,40,46,76,16,16,43,43,22,80,50,45,55,48,24,91,78,75,47,95,28,55,12,57,83,38,34,77,28,46,28,68,11,44,43,95,17,4,20,53,65,10,98,58,42,11,1,12,6,82,5}
	all(sortme)
}

func all(sortme []int) {
	c1 := make(chan []int)
	c2 := make(chan []int)
	c3 := make(chan []int)
	c4 := make(chan []int)
	//sorted := []int{}
	//sorted2 := []int{}
	//sorted3 := []int{}
	//var timeLapsed int64
	start := time.Now()
	
	go doQuick(sortme, c1)
	go bubblesort(sortme, c2)
	go doMerge(sortme, c3)
	go doInstertion(sortme, c4)
	
	sorted := <- c1
	sorted2 := <- c2
	sorted3 := <- c3
	sorted4 := <- c4
	
	fmt.Println(sorted)
	//fmt.Println(sorted1)
	fmt.Println(sorted2)
	fmt.Println(sorted3)
	fmt.Println(sorted4)
	fmt.Println(time.Since(start))
	
	
}
 
func doQuick(a []int, c1 chan []int) {
	quicksort(a)
	c1 <- a
}

func quicksort(a []int) ([]int) {
	//fmt.Println("qs")
	
	if len(a) < 2 {
		return a	
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])
	
	return a	
}

func bubblesort(numbers []int, c2 chan []int) {
	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				intermediate := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = intermediate
			}
		}
	}
	
	c2 <- numbers
	//return numbers
}

func doMerge(a []int, c3 chan []int) {
	mergeSort(a)
	c3 <- a
}

//([]int)
func mergeSort(items []int) ([]int)  {

	var num = len(items)

	if num == 1 {
		
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
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

func doInstertion(a []int,  c4 chan []int) {
	insertionsort(a)
	c4 <- a
}

func insertionsort(items []int)  {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	
}



