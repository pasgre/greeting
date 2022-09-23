// Counting has features to count anything
package counting

import (
	"errors"
)

// Count counts in a range from start to end
func Count(from, to int) ([]int, error) {
	if from > to {
		return []int{}, errors.New("Start is greater than end. Can't count this way. Set from lower than to.")
	}

	var arr []int
	for i := from; from <= to; i++ {
		arr = append(arr, i)
	}

	return arr, nil
}

// Contstant that hold its corresponding number
const (
	One      = 1
	Ten      = 10
	Hunderd  = 100
	Thousand = 1000
)
