package two_sum

import (
	"reflect"
	"slices"
	"testing"
)

type TestCase struct {
	Nums     []int
	Target   int
	Expected []int
}

var tests []TestCase = []TestCase{
	{Nums: []int{2, 7, 11, 15}, Target: 9, Expected: []int{1, 0}},
	{Nums: []int{3, 2, 4}, Target: 6, Expected: []int{1, 2}},
	{Nums: []int{3, 3}, Target: 6, Expected: []int{0, 1}},
	{Nums: []int{-1, -2, -3, -4, -5}, Target: -8, Expected: []int{2, 4}},
}

var benchmark = TestCase{Nums: []int{-1, -2, -3, -4, -5}, Target: -8, Expected: []int{2, 4}}

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		actual := twoSum(test.Nums, test.Target)

		slices.Sort(actual)
		slices.Sort(test.Expected)

		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("twoSum(%v, %v) = %v; expected %v", test.Nums, test.Target, actual, test.Expected)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		twoSum(benchmark.Nums, benchmark.Target)
	}
}
