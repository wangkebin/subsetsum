package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBrutalForceGetMatchedCollections(t *testing.T) {

	// 25 numbers
	test := []int64{86847, 2772, 33707, 79495, 44679, 18234, 63456, 54255, 73826, 62828, 42782, 46857, 9414, 56871, 20764, 58379, 30682, 80548, 10559, 15127, 15766, 11108, 37652, 3092, 33195}
	//There are 1 matched sets.
    // 108135 is the sum of: 44679, 63456, 
	sum := int64(108135)
	start := time.Now()
	subs := BrutalForceAllSums(test)
	indSets := BFGetMatchedCollections(subs, len(test), sum)
	fmt.Printf("There are %d matched sets.\n", len(indSets))
	for _, inds := range indSets {
		fmt.Printf("\n%d is the sum of: ", sum)

		x := int64(0)
		for _, ind := range inds {
			x += test[ind]
			// assert.Contains(t,set, test[ind])
			fmt.Printf("%d, ", test[ind])
		}
		// assert.Equal(t, len(set), len(inds))
		assert.Equal(t, sum, x, "why %d is not the sum of numbers of these indices %v?", sum, inds)
	}
	assert.Greater(t, len(indSets), 0, "no matching subset found.")
	fmt.Printf("\nTotal time %d (ms)\n", time.Since(start).Milliseconds())
	PrintMemUsage()
}
