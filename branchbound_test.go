package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBranchBound(t *testing.T) {

	test := []int64{86847, 2772, 33707, 79495, 44679, 18234, 63456, 54255, 73826, 62828, 42782, 46857, 9414, 56871, 20764, 58379, 30682, 80548, 10559, 15127, 15766, 11108, 37652, 3092, 33195} //24   =>183ms, TotalAlloc Mem 856 MiB
	sum := int64(252170)
	start := time.Now()

	matchedNums := BranchBound(test, sum)
	assert.NotEmpty(t, matchedNums, "")
	x := int64(0)
	for _, num := range matchedNums {
		x += num
	}
	// assert.Equal(t, len(set), len(inds))
	assert.Equal(t, sum, x, "why %d is not the sum of numbers of these indices %v?", sum, matchedNums)

	matchedNums = BranchBound(test, 252170)
	assert.NotEmpty(t, matchedNums, "")

	nomatchedNums := BranchBound(test, 10)
	assert.Empty(t, nomatchedNums, "")

	fmt.Printf("\nTotal time %d (ms)\n", time.Since(start).Milliseconds())
	PrintMemUsage()

	// Stress test 250
	largetest := []int64{81742, 80098, 5011, 93027, 17057, 45923, 12678, 34255, 96277, 51836, 8755, 57267, 31358,
		28499, 27959, 47219, 41251, 88993, 30304, 68384, 83653, 94213, 42235, 30741, 29665, 97179, 40590, 13352,
		71808, 44811, 16101, 68414, 60190, 30712, 51637, 76802, 97112, 39301, 48978, 59086, 34964, 3723, 69876,
		38190, 4835, 91040, 50401, 49837, 70871, 79325, 85059, 86586, 92335, 48675, 87661, 83658, 28248, 8847,
		85459, 5129, 16852, 17638, 4186, 36732, 4106, 53953, 45960, 50365, 93150, 25123, 199, 24304, 41836, 27884,
		24884, 22730, 80651, 13670, 88488, 9442, 94894, 81346, 44527, 90448, 49376, 24215, 10153, 4845, 29031,
		26399, 35378, 23911, 32252, 98910, 10618, 27554, 21688, 12493, 69341, 15770, 86079, 73787, 57486, 63386,
		95220, 30473, 24473, 95397, 74213, 10985, 79579, 36296, 87815, 84349, 68157, 65253, 55632, 24923, 67451,
		15652, 68958, 61844, 57311, 31924, 25081, 84299, 78848, 13380, 86101, 8156, 40360, 42567, 10661, 71281,
		98665, 82742, 52219, 65925, 92466, 37592, 20438, 29313, 17109, 26199, 54124, 5845, 53672, 94103, 52547,
		62330, 42694, 44162, 15009, 28857, 9326, 96031, 22773, 95375, 84811, 46731, 30965, 11882, 4431, 69247, 96927,
		43956, 64165, 51324, 96878, 82361, 91342, 82614, 35030, 59955, 59612, 52443, 45295, 24892, 10366, 21635,
		74616, 45176, 24712, 80330, 44393, 51049, 45930, 21161, 54876, 23015, 84250, 36020, 82589, 69806, 15343,
		58552, 37986, 36159, 48911, 26473, 46291, 32863, 70086, 77394, 41396, 97933, 28561, 49201, 17201, 53765,
		27634, 68314, 80805, 851, 86318, 27129, 74873, 14515, 94905, 74519, 90035, 31689, 85793, 96202, 17697,
		9299, 88919, 51251, 51716, 97134, 67865, 45420, 69265, 63969, 4033, 63179, 90450, 77433, 13667, 53915,
		78375, 85661, 47208, 38661, 67870, 38479, 53920, 55524, 92049, 77963,
	}

	sum = int64(1803598)
	
	// Run time and memorycusage are both negligible.
	matchedNums = BranchBound(largetest, sum)
	PrintMemUsage()
	assert.NotEmpty(t, matchedNums, "")

	sum = int64(180359800)
	matchedNums = BranchBound(largetest, sum)
	PrintMemUsage()
	assert.Empty(t, matchedNums, "")

}
