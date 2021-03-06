package collection

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStrCollection_Base(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	// Len, Size, Count
	assert.Equal(t, 3, c.Len())
	assert.Equal(t, 3, c.Size())
	assert.Equal(t, 3, c.Count())

	// Add, Item
	c.Add("d")
	assert.Equal(t, []string{
		"a", "b", "c", "d",
	}, c.Items())

	// Remove
	c.Remove("b")
	assert.Equal(t, []string{
		"a", "c", "d",
	}, c.Items())
	c.Add("b").Add("b")
	assert.Equal(t, []string{
		"a", "c", "d", "b", "b",
	}, c.Items())
	c.Remove("b")
	assert.Equal(t, []string{
		"a", "c", "d", "b",
	}, c.Items())
	c.Add("b")
	c.RemoveAll("b")
	assert.Equal(t, []string{
		"a", "c", "d",
	}, c.Items())
}

func TestStrCollection_Contains(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	// Contains
	assert.True(t, c.Contains("a"))
	assert.False(t, c.Contains("d"))
}

func TestStrCollection_ClearAndLenAndIsEmpty(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	assert.False(t, c.IsEmpty())
	assert.Equal(t, 3, c.Len())

	// Clear
	c.Clear()
	assert.True(t, c.IsEmpty())
	assert.Equal(t, 0, c.Len())
}

func TestStrCollection_Each(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	// Each
	c.Each(func(i int, v string) {
		if i == 0 {
			assert.Equal(t, "a", v)
		} else if i == 1 {
			assert.Equal(t, "b", v)
		} else if i == 2 {
			assert.Equal(t, "c", v)
		}
	})
}

func TestStrCollection_Map(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	r := c.Map(func(i int, v string) interface{} {
		return v + v
	})

	assert.Equal(t, []interface{}{
		"aa", "bb", "cc",
	}, r)
}

func TestStrCollection_Filter(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	r1 := c.Filter(func(i int, v string) bool {
		return v == "a"
	})
	r2 := c.Filter(func(i int, v string) bool {
		return i == 1
	})

	assert.Equal(t, []string{
		"a",
	}, r1.Items())
	assert.Equal(t, []string{
		"b",
	}, r2.Items())
}

func TestStrCollection_Reject(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	r1 := c.Reject(func(i int, v string) bool {
		return v == "a"
	})
	r2 := c.Reject(func(i int, v string) bool {
		return i == 1
	})

	assert.Equal(t, []string{
		"b", "c",
	}, r1.Items())
	assert.Equal(t, []string{
		"a", "c",
	}, r2.Items())
}

func TestStrCollection_FindAndFindIndex(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	// Find Use Value
	index, value := c.Find(func(i int, v string) bool {
		return v == "a"
	})
	assert.Equal(t, 0, index)
	assert.Equal(t, "a", value)

	// Find Use Index
	index, value = c.Find(func(i int, v string) bool {
		return i == 1
	})

	assert.Equal(t, 1, index)
	assert.Equal(t, "b", value)

	// FindIndex
	index = c.FindIndex(func(i int, v string) bool {
		return v == "c"
	})
	assert.Equal(t, 2, index)
}

func TestStrCollection_FindLastAndFindLastIndex(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	// FindLast Use Value
	index, value := c.FindLast(func(i int, v string) bool {
		return v == "a"
	})
	assert.Equal(t, 3, index)
	assert.Equal(t, "a", value)

	// FindLast Use Index
	index, value = c.FindLast(func(i int, v string) bool {
		return i == 1
	})
	assert.Equal(t, 1, index)
	assert.Equal(t, "b", value)

	// FindLastIndex
	index = c.FindLastIndex(func(i int, v string) bool {
		return v == "a"
	})
	assert.Equal(t, 3, index)
}

func TestStrCollection_Reduce(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	// Reduce
	r := c.Reduce(func(i int, v string, r interface{}) interface{} {
		return r.(string) + v
	}, "")
	assert.Equal(t, "abc", r)
	assert.Equal(t, "abc", r.(string))
}

func TestStrCollection_ReduceRight(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	// ReduceRight
	r := c.ReduceRight(func(i int, v string, r interface{}) interface{} {
		return r.(string) + v
	}, "")
	assert.Equal(t, "cba", r)
	assert.Equal(t, "cba", r.(string))
}

func TestStrCollection_Every(t *testing.T) {
	c := NewStrCollection([]string{
		"aa", "bb", "cc",
	})

	// Every Use Value
	assert.True(t, c.Every(func(i int, v string) bool {
		return v[0] == v[1]
	}))
	assert.False(t, c.Every(func(i int, v string) bool {
		return v == "aa" || v == "bb"
	}))

	// Every Use Index
	assert.True(t, c.Every(func(i int, v string) bool {
		return i <= 2
	}))
	assert.False(t, c.Every(func(i int, v string) bool {
		return i <= 1
	}))
}

func TestStrCollection_Some(t *testing.T) {
	c := NewStrCollection([]string{
		"aa", "b", "c",
	})

	// Some Use Value
	assert.True(t, c.Some(func(i int, v string) bool {
		return len(v) == 2 && v[0] == v[1]
	}))
	assert.False(t, c.Some(func(i int, v string) bool {
		return len(v) == 2 && v[0] != v[1]
	}))

	// Some Use Index
	assert.True(t, c.Some(func(i int, v string) bool {
		return i <= 1
	}))
	assert.False(t, c.Some(func(i int, v string) bool {
		return i > 2
	}))
}

func TestStrCollection_IndexOf(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	// IndexOf
	assert.Equal(t, 0, c.IndexOf("a"))
	assert.Equal(t, 1, c.IndexOf("b"))
	assert.Equal(t, 2, c.IndexOf("c"))
	assert.Equal(t, -1, c.IndexOf("d"))
}

func TestStrCollection_LastIndexOf(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	// LastIndexOf
	assert.Equal(t, 3, c.LastIndexOf("a"))
	assert.Equal(t, 1, c.LastIndexOf("b"))
	assert.Equal(t, 2, c.LastIndexOf("c"))
	assert.Equal(t, -1, c.LastIndexOf("d"))
}

func TestStrCollection_Slice(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	// Slice
	assert.Equal(t, []string{"a", "b", "c", "a"}, c.Slice(0, 4).Items())
	assert.Equal(t, []string{"a", "b", "c"}, c.Slice(0, 3).Items())
	assert.Equal(t, []string{"a", "b"}, c.Slice(0, 2).Items())
	assert.Equal(t, []string{"a"}, c.Slice(0, 1).Items())
	assert.Equal(t, []string{}, c.Slice(0, 0).Items())
	assert.Equal(t, []string{"a"}, c.Slice(1, 0).Items())
	assert.Equal(t, []string{}, c.Slice(1, 1).Items())
	assert.Equal(t, []string{"b", "c", "a"}, c.Slice(1, 4).Items())
	assert.Panics(t, func() {
		c.Slice(1, 5).Items()
	})

	// SliceFrom
	assert.Equal(t, []string{"b", "c", "a"}, c.SliceFrom(1).Items())
	assert.Equal(t, []string{"c", "a"}, c.SliceFrom(2).Items())
	assert.Equal(t, []string{"a"}, c.SliceFrom(3).Items())
	assert.Equal(t, []string{}, c.SliceFrom(4).Items())
	assert.Panics(t, func() {
		c.SliceFrom(5).Items()
	})

	// SliceTo
	assert.Equal(t, []string{"a", "b", "c"}, c.SliceTo(3).Items())
	assert.Equal(t, []string{"a", "b"}, c.SliceTo(2).Items())
	assert.Equal(t, []string{"a"}, c.SliceTo(1).Items())
	assert.Equal(t, []string{}, c.SliceTo(0).Items())
	assert.Equal(t, []string{"a", "b", "c", "a"}, c.SliceTo(-1).Items())
	assert.Equal(t, []string{"a", "b", "c", "a"}, c.SliceTo(4).Items())
	assert.Panics(t, func() {
		c.SliceTo(5).Items()
	})

}

func TestStrCollection_Copy(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	// Copy
	cp := c.Copy()
	assert.Equal(t, []string{"a", "b", "c", "a"}, cp.Items())
	assert.False(t, cp == c)
	assert.NotSame(t, cp, c)

	// Modify
	c.Add("d")
	assert.Equal(t, []string{"a", "b", "c", "a", "d"}, c.Items())
	assert.Equal(t, []string{"a", "b", "c", "a"}, cp.Items())

	// Rmove
	cp.Remove("a")
	assert.Equal(t, []string{"a", "b", "c", "a", "d"}, c.Items())
	assert.Equal(t, []string{"b", "c", "a"}, cp.Items())
}

func TestStrCollection_Reverse(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	cp := c.Reverse()
	assert.Equal(t, []string{"a", "b", "c", "a"}, c.Items())
	assert.Equal(t, []string{"a", "c", "b", "a"}, cp.Items())
}

func TestStrCollection_Shuffle(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c",
	})

	flag := false

	for {
		cs := c.Shuffle()

		if !reflect.DeepEqual(c, cs) {
			flag = true
			break
		}
	}

	assert.True(t, flag)
}

func TestStrCollection_Sort(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})
	c2 := c.Copy()

	// Sort
	c.Sort()
	assert.Equal(t, []string{"a", "a", "b", "c"}, c.Items())

	// Sort
	c2.Sort(func(a, b string) bool {
		return a > b
	})
	assert.Equal(t, []string{"c", "b", "a", "a"}, c2.Items())
}

func TestStrCollection_SortByAndSortByDesc(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	// SortBy
	c.SortBy(func(s string) string {
		return "z" + s
	})
	assert.Equal(t, []string{"a", "a", "b", "c"}, c.Items())

	// SortByDesc
	c.SortByDesc(func(s string) string {
		return "z" + s
	})
	assert.Equal(t, []string{"c", "b", "a", "a"}, c.Items())
}

func TestStrCollection_Unique(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	cu := c.Unique()
	assert.Equal(t, []string{"a", "b", "c"}, cu.Items())
}

func TestStrCollection_MustJson(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	assert.Equal(t, `["a","b","c","a"]`, c.MustJson())
}

func TestStrCollection_String(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	assert.Equal(t, "[a b c a]", c.String())
}

func TestStrCOllection_FirstAndLast(t *testing.T) {
	c := NewStrCollection([]string{
		"a", "b", "c", "a",
	})

	assert.Equal(t, "a", c.First())
	assert.Equal(t, "a", c.Last())
}
