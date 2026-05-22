package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {

	testSlice := generateRandomElements(-5)
	require.Nil(t, testSlice)

	testSlice = generateRandomElements(SIZE)
	assert.Len(t, testSlice, SIZE)
	assert.NotNil(t, testSlice)
}

func TestMaximum(t *testing.T) {

	// handling edge cases
	testStruct := []struct {
		slice  []int
		result int
	}{
		{slice: []int{10, 20, 30}, result: 30},
		{slice: []int{-50, 45}, result: 45},
		{slice: []int{7, 0, 777, 7777}, result: 7777},
		{slice: []int{}, result: 0},
		{slice: []int{1}, result: 1},
		{slice: []int{0}, result: 0},
		{slice: []int{-1}, result: 0},
		{slice: []int{-1, -10, -100}, result: 0},
	}

	for _, s := range testStruct {
		max := maximum(s.slice)
		assert.Equal(t, s.result, max)

	}

	// testing when the number of elements is equal to CHUNKS
	testSlice := make([]int, CHUNKS)
	n := 10
	for i := range testSlice {
		testSlice[i] = n
		n += 10
	}

	max := maximum(testSlice)
	assert.Equal(t, 80, max)

	// testing when the number of elements is equal to CHUNKS+1
	testSliceWhenMoreChunks := make([]int, (CHUNKS + 1))
	v := 10
	for i := range testSliceWhenMoreChunks {
		testSliceWhenMoreChunks[i] = v
		v += 10
	}

	max = maximum(testSliceWhenMoreChunks)
	assert.Equal(t, 90, max)
}

func TestMaxChunks(t *testing.T) {

	testStruct := []struct {
		slice  []int
		result int
	}{
		{slice: []int{10, 20, 30, 40, 50, 60, 70, 80}, result: 80},
		{slice: []int{10, 20, 30, 40, 50, 60, 70, 80, 90}, result: 90},
		{slice: []int{}, result: 0},
	}

	for _, s := range testStruct {

		maximum := maxChunks(s.slice)
		assert.Equal(t, s.result, maximum)

	}
}
