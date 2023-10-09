package longest_substring_without_repeating_characters

import (
	"reflect"
	"testing"
)

type TestCase struct {
	s        string
	expected int
}

var tests = []TestCase{
	{s: "abcabcbb", expected: 3},
	{s: "bbbbb", expected: 1},
	{s: "pwwkew", expected: 3},
	{s: "dvdf", expected: 3},
}

var benchmark = TestCase{s: "abcabcbb", expected: 3}

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	for _, test := range tests {
		actual := longestSubstringWithoutRepeatingCharacters(test.s)

		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("longestSubstringWithoutRepeatingCharacters(%v) = %v; expected %v", test.s, actual, test.expected)
		}
	}
}

func BenchmarkLongestSubstringWithoutRepeatingCharacters(b *testing.B) {
	for n := 0; n < b.N; n++ {
		longestSubstringWithoutRepeatingCharacters(benchmark.s)
	}
}
