package hsort

import (
	"errors"
)


// Sort the list of ints using an insertion algorithm.
func InsertionInt(list []int) error {
	// We're going to follow this sequence for each item in the list:
	// 1. Get the value at the current index.
	// 2. For all previous items--starting from the current position and going down to the beginning--
	//    if the item at the index has a greater value, then shift it one to the right.
	// 3. Insert the value at the now-open index.
	if len(list) < 1 {
		return errors.New("Invalid list size")
	}

	for i, v := range list {
		for i > 0 {
			previous := list[i-1]
			if previous > v {
				// Shift one to the right.
				list[i] = previous
				i--
			} else {
				break
			}
		}
		list[i] = v
	}

	return nil
}

// Sort the list of ints using a selection algorithm.
func SelectionInt(list []int) error {
	// We're going to follow this sequence for each item in the list:
	// 1. Scan the entire list from the current position forward for the lowest value.
	// 2. Swap the current value and the lowest value.
	length := len(list)
	if length < 1 {
		return errors.New("Invalid list size")
	}

	for i := range list {
		pos := i
		for j := i+1; j < length; j++ {
			// Check each value to see if it's lower than our current lowest.
			if list[j] < list[pos] {
				// We found a value lower than we currently have. Select it.
				pos = j
			}
		}
		// Swap the selected value with the current value.
		list[i], list[pos] = list[pos], list[i]
	}

	return nil
}

// Sort the list of ints using a merging algorithm.
func MergeInt(list []int) error {

	return nil
}

