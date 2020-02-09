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
			// Scan down the section of the list that is now sorted.
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
	// For this sorting function, we're going to focus on a stack of blocks. A block is a subsection of the total list.
	// First, we're going to create a block for the entire list. Then we're going to follow this sequence for each sub-block:
	// - Look at the top block on the stack.
	//     - If it hasn't been split yet, then make two blocks out of each half and add them to the stack.
	//     - If it has already been split, then merge its two halves together and throw away the block.
	type block struct {
		index  int
		length int
		merge  bool
	}

	length := len(list)
	if length < 1 {
		return errors.New("Invalid list size")
	}

	b := block{0, length, false}
	s := []block{b}
	for len(s) > 0 {
		// Pop the top block.
		b = s[len(s)-1]
		s = s[:(len(s)-1)]

		leftIndex := b.index
		leftLen := b.length / 2

		rightIndex := b.index + leftLen
		rightLen := b.length - leftLen
		if b.merge {
			// Merge the two halves.
			tmp := make([]int, b.length)
			for i := 0; i < b.length; i++ {
				if leftLen == 0 {
					// We only have values on the right side still.
					tmp[i] = list[rightIndex]
					rightIndex++
				} else if rightLen == 0 {
					// We only have values on the left side still.
					tmp[i] = list[leftIndex]
					leftIndex++
				} else if list[leftIndex] < list[rightIndex] {
					tmp[i] = list[leftIndex]
					leftIndex++
					leftLen--
				} else {
					tmp[i] = list[rightIndex]
					rightIndex++
					rightLen--
				}
			}
			copy(list[b.index:], tmp)
		} else {
			// We're still on the splitting phase.
			b.merge = true
			s = append(s, b)
			if leftLen > 1 {
				// Add left-side block to stack.
				s = append(s, block{leftIndex, leftLen, false})
			}
			if rightLen > 1 {
				// Add right-side block to stack.
				s = append(s, block{rightIndex, rightLen, false})
			}
		}
	}

	return nil
}
