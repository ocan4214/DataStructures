package main

import (
	CommonHeapUtilitiesPackage "Heaps/HeapUtilities"
	"fmt"
)

var util CommonHeapUtilitiesPackage.HeapUtilities = CommonHeapUtilitiesPackage.HeapUtilities{}
var Heap []int = []int{5, 4, 9, 7, 19, 8, 17, 2, 6, 5, 21}

var MaxHeap []int = []int{}

var MinHeap []int = []int{}

func TopDownMaxHeap(arr []int) []int {
	for i := 0; i < len(arr); {
		MaxHeap = append(MaxHeap, arr[i])
		fmt.Printf("****Max Heap is %v **** \n", MaxHeap)

		Sift_Up(MaxHeap, i)
		i++

	}

	fmt.Printf("%v", MaxHeap)
	return MaxHeap
}

func Sift_Up(heap []int, index int) {
	parent := util.Parent(index)
	if parent < 0 {
		return
	}
	fmt.Printf("Slice Informations length = %v, capacity = %v \n", len(heap), cap(heap))
	fmt.Printf(" Parent Index is %v, Child is %v \n", parent, index)
	if heap[parent] < heap[index] {
		util.Swap(&heap[parent], &heap[index])
		index = parent
		Sift_Up(heap, index)
		fmt.Printf("Heap is %v \n", heap)
	}

}

func main() {
	TopDownMaxHeap(Heap)
}
