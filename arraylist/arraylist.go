package arraylist

// 1. Constructing method
type ArrayList struct {
	size     int
	capacity int
	elements []interface{}
}

func NewArrayList(value ...interface{}) *ArrayList {
	return &ArrayList{
		size:     0,
		capacity: 10,
		elements: make([]interface{}, 0),
	}
}

// 2. Resize the array
func (a *ArrayList) resize() {
	if a.size == a.capacity {
		a.capacity *= 2 // expansion
	} else {
		a.capacity /= 2 // shrinkage
	}

	newElements := make([]interface{}, a.capacity)
	copy(newElements, a.elements)
	a.elements = newElements
}

// 3. Add elements
func (a *ArrayList) Add(element interface{}) {
	if a.size == a.capacity {
	}
}
