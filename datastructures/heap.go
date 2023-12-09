package datastructures

import (
	"cmp"
)

type Comparator[T comparable] interface {
	compare(t1 T, t2 T) bool
}

type MaxComparator[T cmp.Ordered] struct{}

func (comparator *MaxComparator[T]) compare(t1 T, t2 T) bool {
	return t1 > t2
}

type MinComparator[T cmp.Ordered] struct{}

func (comparator *MinComparator[T]) compare(t1 T, t2 T) bool {
	return t1 < t2
}

type Heap[T cmp.Ordered] interface {
	Init() Heap[T]
	Insert(element T)
	Extract() T
	Length() int
	IsEmpty() bool
}

type BinaryHeap[T cmp.Ordered] struct {
	array      []T
	comparator Comparator[T]
}

func (heap *BinaryHeap[T]) Insert(element T) {
	heap.array = append(heap.array, element)
	lastIndex := len(heap.array) - 1
	heap.heapifyUp(lastIndex)
}

func (heap *BinaryHeap[T]) Extract() T {
	lastIndex := heap.Length() - 1
	element := heap.array[0]
	heap.swap(0, lastIndex)
	heap.array = heap.array[:lastIndex]
	heap.heapifyDown(0)

	return element
}

func (heap *BinaryHeap[T]) Length() int {
	return len(heap.array)
}

func (heap *BinaryHeap[T]) IsEmpty() bool {
	return heap.Length() == 0
}

func (heap *BinaryHeap[T]) heapifyUp(index int) {
	currentIndex := index

	for parentIndex := heap.getParent(currentIndex); parentIndex >= 0 && heap.comparator.compare(heap.array[currentIndex], heap.array[parentIndex]); parentIndex = heap.getParent(currentIndex) {
		heap.swap(currentIndex, parentIndex)
		currentIndex = parentIndex
	}
}

func (heap *BinaryHeap[T]) getParent(childIndex int) int {
	return (childIndex - 1) >> 1
}

func (heap *BinaryHeap[T]) heapifyDown(index int) {
	length := heap.Length()

	for currentIndex, leftChildIndex := index, heap.getLeftChild(index); leftChildIndex < length; leftChildIndex = heap.getLeftChild(currentIndex) {
		rightChildIndex := heap.getRightChild(currentIndex)
		maxValueIndex := currentIndex

		if heap.comparator.compare(heap.array[leftChildIndex], heap.array[maxValueIndex]) {
			maxValueIndex = leftChildIndex
		}

		if rightChildIndex < length && heap.comparator.compare(heap.array[rightChildIndex], heap.array[maxValueIndex]) {
			maxValueIndex = rightChildIndex
		}

		if maxValueIndex == currentIndex {
			break
		}

		heap.swap(currentIndex, maxValueIndex)
		currentIndex = maxValueIndex
	}
}

func (heap *BinaryHeap[T]) getLeftChild(parentIndex int) int {
	return (parentIndex << 1) + 1
}

func (heap *BinaryHeap[T]) getRightChild(parentIndex int) int {
	return (parentIndex << 1) + 2
}

func (heap *BinaryHeap[T]) swap(index1 int, index2 int) {
	heap.array[index1], heap.array[index2] = heap.array[index2], heap.array[index1]
}
