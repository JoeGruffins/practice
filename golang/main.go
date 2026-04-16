package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(doTwoSums())
	fmt.Println(doIsAnagram())
	fmt.Println(doContainsDuplicate())
	// fmt.Println(doGroupAnagrams())
	fmt.Println(doTopKFrequent())
	fmt.Println(doProductExceptSelf())
	fmt.Println(doLongestConsecutive())
	fmt.Println(doIsPalindrome())
	fmt.Println(doThreeSum())
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

// Given an array of integers `nums` and an integer `target`, return the indices of the two numbers that add up to `target`.
//
// - Each input has exactly one solution
// - You can't use the same element twice
// - Return the answer in any order

func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if j, has := diffs[n]; has {
			return []int{j, i}
		}
		diff := target - nums[i]
		diffs[diff] = i
	}
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

// Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.
//
// An anagram uses all the original letters exactly once, rearranged.

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	found := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		found[s[i]]++
		found[t[i]]--
	}
	for _, v := range found {
		if v != 0 {
			return false
		}
	}
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

// Given an integer array `nums`, return `true` if any value appears at least twice in the array, and `false` if every element is distinct.

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, n := range nums {
		if _, has := seen[n]; has {
			return true
		}
		seen[n] = struct{}{}
	}
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

// Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

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

// Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.

func topKFrequent(nums []int, n int) []int {
	amts := make(map[int]int)
	for _, num := range nums {
		amts[num]++
	}
	bucket := make([][]int, len(nums)+1)
	for k, v := range amts {
		bucket[v] = append(bucket[v], k)
	}
	var highest []int
	for i := len(bucket) - 1; i > 0 && len(highest) < n; i-- {
		ns := bucket[i]
		for j := 0; j < len(ns) && len(highest) < n; j++ {
			highest = append(highest, ns[j])
		}
	}
	return highest
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

// Given an integer array `nums`, return an array `answer` where `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.
//
// You must solve it without using division.

func productExceptSelf(nums []int) []int {
	left, right := make([]int, len(nums)), make([]int, len(nums))
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		right[i] = left[i] * r
		r *= nums[i]
	}
	return right
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

// Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.
//
// Must run in O(n) time.

func longestConsecutive(nums []int) int {
	var highest int
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	for k, _ := range set {
		if _, has := set[k-1]; !has {
			for i := 1; ; i++ {
				if _, has := set[k+i]; !has {
					if i > highest {
						highest = i
						break
					}
				}
			}
		}
	}
	return highest
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
	lower := strings.ToLower(s)
	i, j := 0, len(lower)-1
	isAlpha := func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
	}
	for i < j {
		if !isAlpha(lower[i]) {
			i++
			continue
		}
		if !isAlpha(lower[j]) {
			j--
			continue
		}
		if lower[i] != lower[j] {
			return false
		}
		i++
		j--
	}
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

// Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.
//
// The solution must not contain duplicate triplets.

func threeSum(nums []int) [][]int {
	var three [][]int
	slices.Sort(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[j] + nums[k] + nums[i]
			if sum > 0 {
				j++
				continue
			}
			if sum < 0 {
				k--
				continue
			}
			three = append(three, []int{nums[j], nums[k], nums[i]})
			j++
			k--
		}
	}
	return three
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
