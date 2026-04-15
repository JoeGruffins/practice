package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(doTwoSums)
}

func doTwoSums() error {
	tests := []struct {
		inputs []int
		target int
		output []int
	}{{
		inputs: []int{2, 7, 11, 15},
		target: 9,
		output: []int{0, 1},
	}, {
		inputs: []int{3, 2, 4},
		target: 6,
		output: []int{1, 2},
	}, {
		inputs: []int{3, 3},
		target: 6,
		output: []int{0, 1},
	}}
	for i, test := range tests {
		output := twoSum(test.inputs, test.target)
		if !slices.Equal(test.output, output) {
			return fmt.Errorf("brute force bad output at index %d got %v want %v", i, output, test.output)
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	return nil
}

func doIsAnagram() error {
	tests := []struct {
		s, t string
		want bool
	}{{
		s:    "anagram",
		t:    "nagaram",
		want: true,
	}, {
		s: "rat",
		t: "car",
	}, {
		s:    "listen",
		t:    "silent",
		want: true,
	}, {
		s: "aab",
		t: "bba",
	}, {
		s: "ad",
		t: "bc",
	}, {
		s: "acc",
		t: "bbb",
	}}
	for i, test := range tests {
		is := isAnagram(test.s, test.t)
		if test.want != is {
			return fmt.Errorf("isanagram wrong at index %d", i)
		}
	}
	return nil
}

func isAnagram(s string, t string) bool {
	return true
}

func doContainsDuplicate() error {
	tests := []struct {
		nums []int
		want bool
	}{{
		nums: []int{1, 2, 3, 1},
		want: true,
	}, {
		nums: []int{1, 2, 3, 4},
	}, {
		nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		want: true,
	}}
	for i, test := range tests {
		got := containsDuplicate(test.nums)
		if test.want != got {
			return fmt.Errorf("containsDuplicate wrong at index %d got %v want %v", i, got, test.want)
		}
	}
	return nil
}

func containsDuplicate(nums []int) bool {
	return false
}

func doGroupAnagrams() error {
	tests := []struct {
		strs []string
		want [][]string
	}{{
		strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
	}, {
		strs: []string{""},
		want: [][]string{{""}},
	}, {
		strs: []string{"a"},
		want: [][]string{{"a"}},
	}}
	for i, test := range tests {
		got := groupAnagrams(test.strs)
		if len(got) != len(test.want) {
			return fmt.Errorf("groupAnagrams wrong at index %d got %d groups want %d", i, len(got), len(test.want))
		}
		// Sort inner slices and outer slice for comparison.
		for _, g := range got {
			slices.Sort(g)
		}
		for _, g := range test.want {
			slices.Sort(g)
		}
		slices.SortFunc(got, func(a, b []string) int {
			return strings.Compare(a[0], b[0])
		})
		slices.SortFunc(test.want, func(a, b []string) int {
			return strings.Compare(a[0], b[0])
		})
		for j := range got {
			if !slices.Equal(got[j], test.want[j]) {
				return fmt.Errorf("groupAnagrams wrong at index %d group %d got %v want %v", i, j, got[j], test.want[j])
			}
		}
	}
	return nil
}

func groupAnagrams(strs []string) (grouped [][]string) {
	return
}

func doTopKFrequent() error {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{{
		nums: []int{1, 1, 1, 2, 2, 3},
		k:    2,
		want: []int{1, 2},
	}, {
		nums: []int{1},
		k:    1,
		want: []int{1},
	}, {
		nums: []int{4, 4, 4, 1, 1, 2, 2, 2, 3},
		k:    2,
		want: []int{4, 2},
	}}
	for i, test := range tests {
		got := topKFrequent(test.nums, test.k)
		if len(got) != len(test.want) {
			return fmt.Errorf("topKFrequent wrong at index %d got %v want %v", i, got, test.want)
		}
		slices.Sort(got)
		slices.Sort(test.want)
		if !slices.Equal(got, test.want) {
			return fmt.Errorf("topKFrequent wrong at index %d got %v want %v", i, got, test.want)
		}
	}
	return nil
}

func topKFrequent(nums []int, n int) []int {
	return nil
}

func doProductExceptSelf() error {
	tests := []struct {
		nums []int
		want []int
	}{{
		nums: []int{1, 2, 3, 4},
		want: []int{24, 12, 8, 6},
	}, {
		nums: []int{-1, 1, 0, -3, 3},
		want: []int{0, 0, 9, 0, 0},
	}}
	for i, test := range tests {
		got := productExceptSelf(test.nums)
		if !slices.Equal(got, test.want) {
			return fmt.Errorf("productExceptSelf wrong at index %d got %v want %v", i, got, test.want)
		}
	}
	return nil
}

func productExceptSelf(nums []int) []int {
	return nil
}

func doLongestConsecutive() error {
	tests := []struct {
		nums []int
		want int
	}{{
		nums: []int{100, 4, 200, 1, 3, 2},
		want: 4,
	}, {
		nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		want: 9,
	}, {
		nums: nil,
		want: 0,
	}}
	for i, test := range tests {
		got := longestConsecutive(test.nums)
		if got != test.want {
			return fmt.Errorf("longestConsecutive wrong at index %d got %d want %d", i, got, test.want)
		}
	}
	return nil
}

func longestConsecutive(nums []int) int {
	return 0
}

func doIsPalindrome() error {
	tests := []struct {
		s    string
		want bool
	}{{
		s:    "A man, a plan, a canal: Panama",
		want: true,
	}, {
		s: "race a car",
	}, {
		s:    " ",
		want: true,
	}}
	for i, test := range tests {
		got := isPalindrome(test.s)
		if got != test.want {
			return fmt.Errorf("isPalindrome wrong at index %d got %v want %v", i, got, test.want)
		}
	}
	return nil
}

func isPalindrome(s string) bool {
	return true
}

func doThreeSum() error {
	tests := []struct {
		nums []int
		want [][]int
	}{{
		nums: []int{-1, 0, 1, 2, -1, -4},
		want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
	}, {
		nums: []int{0, 1, 1},
		want: [][]int{},
	}, {
		nums: []int{0, 0, 0},
		want: [][]int{{0, 0, 0}},
	}}
	for i, test := range tests {
		got := threeSum(test.nums)
		if len(got) != len(test.want) {
			return fmt.Errorf("threeSum wrong at index %d got %v want %v", i, got, test.want)
		}
		for _, g := range got {
			slices.Sort(g)
		}
		slices.SortFunc(got, func(a, b []int) int {
			for k := range a {
				if a[k] != b[k] {
					return a[k] - b[k]
				}
			}
			return 0
		})
		for j := range got {
			if !slices.Equal(got[j], test.want[j]) {
				return fmt.Errorf("threeSum wrong at index %d triplet %d got %v want %v", i, j, got[j], test.want[j])
			}
		}
	}
	return nil
}

func threeSum(nums []int) [][]int {
	return nil
}

func doMaxArea() error {
	tests := []struct {
		height []int
		want   int
	}{{
		height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		want:   49,
	}, {
		height: []int{1, 1},
		want:   1,
	}}
	for i, test := range tests {
		got := maxArea(test.height)
		if got != test.want {
			return fmt.Errorf("maxArea wrong at index %d got %d want %d", i, got, test.want)
		}
	}
	return nil
}

func maxArea(height []int) int {
	return 0
}

func doTrap() error {
	tests := []struct {
		height []int
		want   int
	}{{
		height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		want:   6,
	}, {
		height: []int{4, 2, 0, 3, 2, 5},
		want:   9,
	}}
	for i, test := range tests {
		got := trap(test.height)
		if got != test.want {
			return fmt.Errorf("trap wrong at index %d got %d want %d", i, got, test.want)
		}
	}
	return nil
}

func trap(height []int) int {
	return 0
}

func doIsValid() error {
	tests := []struct {
		s    string
		want bool
	}{{
		s:    "()",
		want: true,
	}, {
		s:    "()[]{}",
		want: true,
	}, {
		s: "(]",
	}, {
		s: "([)]",
	}, {
		s:    "{[]}",
		want: true,
	}}
	for i, test := range tests {
		got := isValid(test.s)
		if got != test.want {
			return fmt.Errorf("isValid wrong at index %d got %v want %v", i, got, test.want)
		}
	}
	return nil
}

func isValid(s string) bool {
	return false
}

//# Min Stack
//
//Design a stack that supports push, pop, top, and retrieving the minimum element, all in O(1) time.
//
//Implement the following:
//- `Push(val)` - pushes the element onto the stack
//- `Pop()` - removes the element on top of the stack
//- `Top()` - gets the top element
//- `GetMin()` - retrieves the minimum element in the stack

func doMinStack() error {
	s := &MinStack{}
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if got := s.GetMin(); got != -3 {
		return fmt.Errorf("GetMin wrong got %d want -3", got)
	}
	s.Pop()
	if got := s.Top(); got != 0 {
		return fmt.Errorf("Top wrong got %d want 0", got)
	}
	if got := s.GetMin(); got != -2 {
		return fmt.Errorf("GetMin wrong got %d want -2", got)
	}
	return nil
}

type MinStack struct {
}

func (s *MinStack) Push(val int) {
}
func (s *MinStack) Pop() {
}
func (s *MinStack) Top() int {
	return 0
}
func (s *MinStack) GetMin() int {
	return 0
}
