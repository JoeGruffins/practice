package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(doTwoSums())
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

//# Valid Anagram
//
//Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.
//
//An anagram uses all the original letters exactly once, rearranged.
//
//## Examples
//
//```
//Input: s = "anagram", t = "nagaram"
//Output: true
//
//Input: s = "rat", t = "car"
//Output: false
//
//Input: s = "listen", t = "silent"
//Output: true
//```
//
//## Function Signature
//
//```go
//func isAnagram(s string, t string) bool {
//
//}
//```
//
//## Try to solve it two ways:
//1. The simple way (sort both strings and compare)
//2. The optimal way (think about counting)

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
	getValue := func(w string) uint32 {
		var total uint32
		for _, r := range w {
			x := uint32('z') - uint32(r)
			total += 1 >> x
		}
		return total
	}
	return getValue(s) == getValue(t)
}
