package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(doContainsDuplicate())
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
		output := bruteTwoSum(test.inputs, test.target)
		if !slices.Equal(test.output, output) {
			return fmt.Errorf("brute force bad output at index %d got %v want %v", i, output, test.output)
		}
		output = optimalTwoSum(test.inputs, test.target)
		if !slices.Equal(test.output, output) {
			return fmt.Errorf("optimal bad output at index %d got %v want %v", i, output, test.output)
		}
	}
	return nil
}

func bruteTwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func optimalTwoSum(nums []int, target int) []int {
	want := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if widx, ok := want[nums[i]]; ok {
			return []int{widx, i}
		}
		want[target-nums[i]] = i
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

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	seen := make(map[rune]int)
	for _, r := range s {
		seen[r]++
	}
	for _, r := range t {
		if seen[r] == 0 {
			return false
		}
		seen[r]--
	}
	return true
}

//# Contains Duplicate
//
//Given an integer array `nums`, return `true` if any value appears at least twice in the array, and `false` if every element is distinct.
//
//## Examples
//
//```
//Input: nums = [1, 2, 3, 1]
//Output: true
//
//Input: nums = [1, 2, 3, 4]
//Output: false
//
//Input: nums = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
//Output: true
//```
//
//## Function Signature
//
//```go
//func containsDuplicate(nums []int) bool {
//
//}

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
	seen := make(map[int]struct{})
	for _, v := range nums {
		if _, has := seen[v]; has {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
