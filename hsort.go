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

	return nil
}

// Sort the list of ints using a merging algorithm.
func MergeInt(list []int) error {

	return nil
}

