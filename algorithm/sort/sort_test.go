package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

var arr []int

func TestBubbleSort(t *testing.T) {
	BubbleSort(arr, len(arr))
	fmt.Println(arr)
}

func TestInsertionSort(t *testing.T) {
	InsertionSort(arr, len(arr))
	fmt.Println(arr)
}

func TestShellSort(t *testing.T) {
	ShellSort(arr, len(arr))
	fmt.Println(arr)
}

func TestSelectionSort(t *testing.T) {
	SelectionSort(arr, len(arr))
	fmt.Println(arr)
}

func TestHeapSort(t *testing.T) {
	HeapSort(arr, len(arr))
	fmt.Println(arr)
}

func TestMergeSort(t *testing.T) {
	MergeSort(arr, len(arr))
	fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
	QuickSort(arr, len(arr))
	fmt.Println(arr)
}

func TestLSDRadixSort(t *testing.T) {
	LSDRadixSort(arr, len(arr))
	fmt.Println(arr)
}

func TestMSDRadixSort(t *testing.T) {
	MSDRadixSort(arr, len(arr))
	fmt.Println(arr)
}

func TestMain(m *testing.M) {
	n := 1000
	arr = make([]int, n)
	for i := range arr {
		data := rand.Intn(10000)
		arr[i] = data
	}
	m.Run()
}
