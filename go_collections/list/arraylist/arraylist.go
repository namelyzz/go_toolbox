package arraylist

import (
	"errors"
	"fmt"
)

const (
	Capacity = 10
)

// ArrayList Constructing method
type ArrayList struct {
	size     int
	capacity int
	elements []interface{}
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		size:     0,
		capacity: Capacity,
		elements: make([]interface{}, 0),
	}
}

// Resize the array
func (a *ArrayList) resize() {
	if a.size >= a.capacity {
		a.capacity *= 2 // expansion
	} else {
		a.capacity /= 2 // shrinkage
	}

	newElements := make([]interface{}, a.size, a.capacity)
	copy(newElements, a.elements)
	a.elements = newElements
}

// Get elements
func (a *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= a.size {
		return nil, errors.New(fmt.Sprintf("arraylist out of bounds, index: %d", index))
	}

	return a.elements[index], nil
}

// Add elements
func (a *ArrayList) Add(value interface{}) {
	a.size += 1
	if a.size == a.capacity {
		a.resize()
	}

	a.elements = append(a.elements, value)
}

func (a *ArrayList) AddAll(values ...interface{}) {
	a.size += len(values)
	for a.size >= a.capacity {
		a.resize()
	}

	a.elements = append(a.elements, values...)
}

// Remove elements
func (a *ArrayList) Remove(index int) error {
	if index < 0 || index >= a.size {
		return errors.New(fmt.Sprintf("arraylist out of bounds, index: %d", index))
	}

	newArray := a.elements[:index]
	newArray = append(newArray, a.elements[index+1:]...)
	a.elements = newArray
	a.size--

	if a.size < a.capacity/2 {
		a.resize()
	}

	return nil
}

// Clear elements
func (a *ArrayList) Clear() {
	a.elements = make([]interface{}, 0)
	a.size = 0
	a.capacity = Capacity
}

// Size returns the size of arraylist
func (a *ArrayList) Size() int {
	return a.size
}

// Contains return true if the arraylist contains the given element
func (a *ArrayList) Contains(value interface{}) bool {
	for _, v := range a.elements {
		if v == value {
			return true
		}
	}

	return false
}

// IsEmpty returns true if the arraylist is empty
func (a *ArrayList) IsEmpty() bool {
	return a.size == 0
}
