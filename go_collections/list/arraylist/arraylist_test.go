package arraylist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewArrayList(t *testing.T) {
	al := NewArrayList()
	assert.Equal(t, []interface{}{}, al.elements)
}

func TestArrayList_Add(t *testing.T) {
	al := NewArrayList()
	al.Add(1)
	assert.Len(t, al.elements, 1)
	res, ok := al.elements[0].(int)
	if !ok {
		t.Errorf("type is not int")
	}

	assert.Equal(t, 1, res)
	assert.Equal(t, 1, al.size)
}

func TestArrayList_AddAll(t *testing.T) {
	al := NewArrayList()
	cases := []int{1, 2, 3}
	al.AddAll(1, 2, 3)
	assert.Len(t, al.elements, 3)

	for idx, val := range al.elements {
		assert.Equal(t, cases[idx], val)
	}

	assert.Equal(t, 3, al.size)
}

func TestArrayList_Get(t *testing.T) {
	al := NewArrayList()
	al.AddAll(1, 2, 3)
	res, err := al.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, 2, res.(int))
}

func TestArrayList_Clear(t *testing.T) {
	al := NewArrayList()
	al.AddAll(1, 2, 3)
	al.Clear()
	assert.Len(t, al.elements, 0)
}

func TestArrayList_Contains(t *testing.T) {
	al := NewArrayList()
	al.AddAll(1, 2, 3)
	assert.True(t, al.Contains(1))
}

func TestArrayList_IsEmpty(t *testing.T) {
	al := NewArrayList()
	assert.True(t, al.IsEmpty())

	al.AddAll(1, 2, 3)
	assert.False(t, al.IsEmpty())
}

func TestArrayList_Remove(t *testing.T) {
	al := NewArrayList()
	al.AddAll(1, 2, 3)
	err := al.Remove(0)
	assert.Nil(t, err)
	assert.Len(t, al.elements, 2)
	for _, i := range al.elements {
		t.Log(i)
	}
	assert.Equal(t, al.size, 2)
}

func TestArrayList_Size(t *testing.T) {
	al := NewArrayList()
	al.AddAll(1, 2, 3)
	assert.Equal(t, 3, al.Size())
}

func TestNewArrayList_resize(t *testing.T) {
	al := NewArrayList()
	al.resize()
	assert.Equal(t, 5, al.capacity)

	al.AddAll(1, 2, 3, 4, 5, 6)
	assert.Equal(t, 10, al.capacity)
}
