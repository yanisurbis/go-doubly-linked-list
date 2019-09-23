package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_Value(t *testing.T) {
	var l List
	l.PushFront(1)

	elm := l.First()
	assert.Equal(t, elm.Value(), 1)
}

func TestNode_Next(t *testing.T) {
	var l List
	l.PushFront(1)
	l.PushFront(2)

	elm := l.First().Next()
	assert.Equal(t, elm.Value(), 1)
}

func TestNode_Last(t *testing.T) {
	var l List
	l.PushFront(1)
	l.PushFront(2)

	elm := l.Last().Prev()
	assert.Equal(t, elm.Value(), 2)
}

func TestList_Len(t *testing.T) {
	var l1 List
	assert.Equal(t, l1.Len(), 0, "works with empty list")

	var l2 List
	l2.PushFront(1)
	l2.PushFront(1)
	assert.Equal(t, l2.Len(), 2, "works with non empty list")
}

func TestList_First(t *testing.T) {
	var l1 List
	assert.Nil(t, l1.First(), "works with empty list")

	var l2 List
	l2.PushFront(1)
	l2.PushFront(2)
	assert.Equal(t, l2.First().Value(), 2, "works with non empty list")
}

func TestList_Last(t *testing.T) {
	var l1 List
	assert.Nil(t, l1.Last(), "works with empty list")

	var l2 List
	l2.PushFront(1)
	l2.PushFront(2)
	assert.Equal(t, l2.Last().Value(), 1, "works with non empty list")
}

func TestList_PushFront(t *testing.T) {
	t.Run("works correctly with only one element", func(t *testing.T) {
		var l1 List
		l1.PushFront(100)

		assert.Equal(t, l1.Len(), 1)
		assert.Equal(t, l1.First().Value(), 100)
		assert.Equal(t, l1.Last().Value(), 100)
	})

	t.Run("works correctly with multiple elements", func(t *testing.T) {
		var l2 List
		l2.PushFront(100)
		l2.PushFront(200)

		assert.Equal(t, l2.Len(), 2, )
		assert.Equal(t, l2.First().Value(), 200)
		assert.Equal(t, l2.Last().Value(), 100)
	})
}

func TestList_PushBack(t *testing.T) {
	t.Run("works correctly with only one element", func(t *testing.T) {
		var l1 List
		l1.PushBack(100)

		assert.Equal(t, l1.Len(), 1)
		assert.Equal(t, l1.First().Value(), 100)
		assert.Equal(t, l1.Last().Value(), 100)
	})

	t.Run("works correctly with multiple elements", func(t *testing.T) {
		var l2 List
		l2.PushBack(100)
		l2.PushBack(200)

		assert.Equal(t, l2.Len(), 2, )
		assert.Equal(t, l2.First().Value(), 100)
		assert.Equal(t, l2.Last().Value(), 200)
	})
}

func TestList_Remove(t *testing.T) {
	t.Run("works with empty list", func(t *testing.T) {
		var l List

		r := Node{}
		l.Remove(r)

		assert.Equal(t, l.Len(), 0)
	})

	t.Run("works with non existing values", func(t *testing.T) {
		var l List
		l.PushFront(100)
		l.PushFront(200)
		l.PushFront(300)
		l.PushFront(400)
		r1 := Node{
			next: nil,
			prev: nil,
			data: 5,
		}
		r2 := Node{}

		l.Remove(r1)
		l.Remove(r2)

		assert.Equal(t, l.Len(), 4)
	})

	t.Run("works with one element list", func(t *testing.T) {
		var l List
		l.PushFront(1)
		r := l.First()

		l.Remove(*r)

		assert.Equal(t, l.Len(), 0)
		assert.Nil(t, l.First())
		assert.Nil(t, l.Last())
	})

	t.Run("works when deleting first element", func(t *testing.T) {
		var l List
		l.PushFront(100)
		l.PushFront(200)
		l.PushFront(300)
		l.PushFront(400)
		r := l.First()

		l.Remove(*r)

		assert.Equal(t, l.Len(), 3)
		assert.Equal(t, l.First().Value(), 300)
	})

	t.Run("works when deleting last element", func(t *testing.T) {
		var l List
		l.PushFront(100)
		l.PushFront(200)
		l.PushFront(300)
		l.PushFront(400)
		r := l.Last()

		l.Remove(*r)

		assert.Equal(t, l.Len(), 3)
		assert.Equal(t, l.Last().Value(), 200)
	})

	t.Run("works in general case", func(t *testing.T) {
		var l List
		l.PushFront(100)
		l.PushFront(200)
		l.PushFront(300)
		l.PushFront(400)
		r := l.First().Next().Next()

		l.Remove(*r)

		assert.Equal(t, l.Len(), 3)
	})
}
