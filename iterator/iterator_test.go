package iterator

import (
	"fmt"
	"testing"
)

func TestArrayIterator(t *testing.T) {
	array := []interface{}{1, 2, 3, 4, 5, 6}
	index := 0
	iterator := ArrayIterator{array: array, index: &index}
	for it := iterator; iterator.HasNext(); iterator.Next() {
		index, value := it.Index(), it.Value().(int)
		for value != array[*index] {
			fmt.Println("error")
		}

		fmt.Println(*index, value)
	}
}
