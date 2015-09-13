package sorting_algorithm

import "testing"

func TestSort(t *testing.T) {
	insertionSort([]int{4, 42, 5, 1, 11, 3, 15, 53, 31, 32})
	insertionSort2([]int{4, 42, 5, 1, 11, 3, 15, 53, 31, 32})
	shellSort([]int{4, 42, 5, 1, 11, 3, 15, 53, 31, 32})
	bubbleSort([]int{4, 42, 5, 1, 11, 3, 15, 53, 31, 32})
	t.Log("quickSort =========")
	quickSort([]int{4, 42, 5, 1, 11, 3, 15, 53, 31, 32})
	selectionSort([]int{4, 42, 5, 1, 11, 3, 15, 53, 31, 32})
	a := mergeSort([]int{4, 42, 5, 1, 11, 3, 15, 53, 31, 32})
	t.Logf("mergeSort %v", a)
	h, _ := NewHeap([]int{4, 42, 5, 1, 11, 3, 15, 53, 31, 32})
	t.Logf("heap %v", h.Data)
	h.Sort()
	t.Logf("heapSort %v", h.Data)
}
