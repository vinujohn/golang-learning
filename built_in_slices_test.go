package learning

import (
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://go.dev/blog/slices-intro
func TestSlices(t *testing.T) {
	// an array is a fixed length data type
	arr := [3]int{1, 2, 3}

	// in fact, arrays of different lenths are different data types
	assert.NotEqual(t, reflect.TypeOf([4]int{}), reflect.TypeOf(arr)) // [4]int not equal to [3]int

	// you can create a slice from an array as so
	assert.Equal(t, reflect.TypeOf([]int{}), reflect.TypeOf(arr[:])) // using the slice expression '[:]' turns arr into a slice

	// sending an array to a function and then modifying the array does not change it. A copy of the arr and its contents is made
	modifyArr(arr)
	assert.Equal(t, 1, arr[0]) // no change to arr in the modifyArr function

	// however, this is not the same for a slice
	modifySlice(arr[:])
	assert.Equal(t, 99, arr[0]) // since modifySlice changed the value of the underlying array, we see changes here

	// you can also send a pointer of an array to get modified
	modifyArrPtr(&arr)
	assert.Equal(t, 89, arr[0]) // since modifyArrPtr has a reference to the array, no copy of the arr actually takes place

	// you can also send an array of int pointers and modify them
	one, two, three := 1, 2, 3
	arrOfIntPtrs := [3]*int{&one, &two, &three}
	modifyArrOfIntPtr(arrOfIntPtrs)
	assert.Equal(t, 89, *arrOfIntPtrs[0]) // arrOfIntPtrs[0] is a pointer to an int which changes to 89

	// there are a couple of ways to make slices
	a := []int{0, 0, 0}    // slice literal
	b := make([]int, 3)    // make([]Type, len)
	c := make([]int, 3, 4) // make([]Type, len, cap)
	assert.Equal(t, a, b)
	assert.Equal(t, a, c)

	// zero value of slice
	var arrNil []int
	assert.Nil(t, arrNil)
	assert.Equal(t, 0, len(arrNil))              // len function still works on nil slices
	clear(arrNil)                                // clear works on nil slice
	assert.Equal(t, []int{5}, append(arrNil, 5)) // even appending works
	assert.Nil(t, arrNil)
	assert.Panics(t, func() { _ = arrNil[0] }) // only indexing causes a panic

	// append will append values from another slice to a slice
	appendTo, appendFrom := []int{99}, []int{1, 2, 3}
	appended := append(appendTo, appendFrom...)
	assert.Equal(t, []int{99, 1, 2, 3}, appended)

	// you can also prepend
	prepend := append(appendFrom, appendTo...)
	assert.Equal(t, []int{1, 2, 3, 99}, prepend)

	// you can also use append to delete elements
	del := []int{1, 2, 3, 4, 5}
	assert.Equal(t, []int{1, 2, 5}, append(del[:2], del[4:]...))
	// new slices.Delete in action which does the same thing
	assert.Equal(t, []int{1, 2, 5}, slices.Delete(del, 2, 4))

	// you can copy a slice to another slice using the copy function
	copyTo, copyFrom := []int{}, []int{1, 2, 3}
	copy(copyTo, copyFrom)
	assert.Equal(t, []int{}, copyTo) // since the copyTo slice has len of zero, nothing can be copied
	copyTo = append(copyTo, 99)
	assert.Equal(t, []int{99}, copyTo)
	copy(copyTo, copyFrom)
	assert.Equal(t, []int{1}, copyTo) // since the copyTo slice of len of 1, 1 element was copied

	// you can use slice expressions to create subsets of other slices
	bigSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	subsetSlice := bigSlice[1:4] // slice expression: [startIdx:desiredEndIdx+1]
	assert.Equal(t, []int{1, 2, 3}, subsetSlice)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, bigSlice)

	// be careful when modifying slices that point to the same underlying array
	subsetSlice[0] = 99
	assert.Equal(t, []int{99, 2, 3}, subsetSlice)
	assert.Equal(t, []int{0, 99, 2, 3, 4, 5, 6, 7, 8, 9}, bigSlice)
	assert.Equal(t, subsetSlice[0], bigSlice[1])

	// also be careful when appending to a subset array. the larger array will have values overwritten
	subsetSlice = append(subsetSlice, []int{98, 97}...)
	assert.Equal(t, []int{99, 2, 3, 98, 97}, subsetSlice)
	assert.Equal(t, []int{0, 99, 2, 3, 98, 97, 6, 7, 8, 9}, bigSlice)

	// you can set slice elements to their zero value
	toZeroOut := []bool{true, true, true}
	clear(toZeroOut)
	assert.Equal(t, []bool{false, false, false}, toZeroOut)
}

func modifySlice(s []int) {
	s[0] = 99
}

func modifyArr(arr [3]int) {
	arr[0] = 99
}

func modifyArrPtr(arr *[3]int) {
	arr[0] = 89
}

func modifyArrOfIntPtr(arr [3]*int) {
	*arr[0] = 89
}
