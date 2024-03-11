package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGreedy(t *testing.T) {

	test := []int64{86847, 2772, 33707, 79495, 44679, 18234, 63456, 54255, 73826, 62828, 42782, 46857, 9414, 56871, 20764, 58379, 30682, 80548, 10559, 15127, 15766, 11108, 37652, 3092, 33195}
	sum := int64(108135)
	//{test[6], test[17], test[21]}
	start := time.Now()

	matchedNums := Greedy(test, 108135)
	assert.NotEmpty(t, matchedNums, "")
	x := int64(0)
	for _, num := range matchedNums {
		x += num
	}
	// assert.Equal(t, len(set), len(inds))
	assert.Equal(t, sum, x, "why %d is not the sum of numbers of these indices %v?", sum, matchedNums)

	// Limitation of greedy. 177389 is the sum of these numbers
	// 2772, 63456, 42782, 58379,
	matchedNums2 := Greedy(test, 177389)
	// assert.NotEmpty(t, matchedNums2, "")
	assert.Empty(t, matchedNums2, "")

	nomatchedNums := Greedy(test, 1000000)
	assert.Empty(t, nomatchedNums, "")

	fmt.Printf("\nTotal time %d (ms)\n", time.Since(start).Milliseconds())
	PrintMemUsage()
}
