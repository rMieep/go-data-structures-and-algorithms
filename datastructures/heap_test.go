package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyHeapExtractPanics(t *testing.T) {
	heap := new(BinaryHeap[int])
	heap.comparator = new(MaxComparator[int])

	assert.PanicsWithError(t, "runtime error: index out of range [0] with length 0", func() { heap.Extract() })
}

func TestExtract(t *testing.T) {
	heap := new(BinaryHeap[int])
	heap.array = []int{5, 4, 3, 2, 1}
	heap.comparator = new(MaxComparator[int])

	var extractedValue int = heap.Extract()

	assert.Equal(t, 5, extractedValue)
	assert.Equal(t, []int{4, 2, 3, 1}, heap.array)
}

func TestInsert(t *testing.T) {
	heap := new(BinaryHeap[int])
	heap.array = []int{5, 3, 2, 1}
	heap.comparator = new(MaxComparator[int])

	heap.Insert(6)

	assert.Equal(t, []int{6, 5, 2, 1, 3}, heap.array)
}

func TestLength(t *testing.T) {
	heap := new(BinaryHeap[int])
	heap.array = []int{5, 3, 2, 1}

	length := heap.Length()

	assert.Equal(t, 4, length)
}

func TestIsEmptyTrue(t *testing.T) {
	heap := new(BinaryHeap[int])

	isEmpty := heap.IsEmpty()

	assert.Equal(t, true, isEmpty)
}

func TestIsEmptyFalse(t *testing.T) {
	heap := new(BinaryHeap[int])
	heap.array = []int{1}

	isEmpty := heap.IsEmpty()

	assert.Equal(t, false, isEmpty)
}
