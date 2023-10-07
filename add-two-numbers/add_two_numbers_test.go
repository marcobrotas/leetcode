package add_two_numbers

import (
	"reflect"
	"testing"
)

type TestCase struct {
	l1       *ListNode
	l2       *ListNode
	Expected *ListNode
}

var tests []TestCase = []TestCase{
	{l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}, l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}, Expected: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}},
	{l1: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}}, l2: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}, Expected: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}}}},
}

var benchmark = TestCase{l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}, l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}, Expected: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}}}

func TestAddTwoNumbers(t *testing.T) {
	for _, test := range tests {
		actual := addTwoNumbers(test.l1, test.l2)

		if !reflect.DeepEqual(test.Expected, actual) {
			t.Errorf("addTwoNumbers(%v, %v) = %v; expected %v", test.l1, test.l2, actual, test.Expected)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		addTwoNumbers(benchmark.l1, benchmark.l2)
	}
}
