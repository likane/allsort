package allsort

import (
	"crypto/rand"
	"math/rand"
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

func all(a []int) []int {
	go quicksort(a)
	go bubblesort(a)
	go mergesort(a)
	go insertionsort(a)

}

//https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
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

//https://www.golangprograms.com/golang-program-for-implementation-of-bubble-sort.html
func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

// func main() {

//     slice := generateSlice(20)
//     fmt.Println("\n--- Unsorted --- \n\n", slice)
//     fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice), "\n")
// }

// Generates a slice of size, size filled with random numbers
// func generateSlice(size int) []int {

//     slice := make([]int, size, size)
//     rand.Seed(time.Now().UnixNano())
//     for i := 0; i < size; i++ {
//         slice[i] = rand.Intn(999) - rand.Intn(999)
//     }
//     return slice
// }

func mergeSort(items []int) []int {
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

	return merge(mergeSort(left), mergeSort(right))
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
func insertionsort(items []int) {
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
