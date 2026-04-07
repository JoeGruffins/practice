package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(doTwoSums())
}

//# Two Sum
//
//Given an array of integers `nums` and an integer `target`, return the indices of the two numbers that add up to `target`.
//
//- Each input has exactly one solution
//- You can't use the same element twice
//- Return the answer in any order
//
//## Examples
//
//```
//Input: nums = [2, 7, 11, 15], target = 9
//Output: [0, 1]   (because 2 + 7 = 9)
//
//Input: nums = [3, 2, 4], target = 6
//Output: [1, 2]   (because 2 + 4 = 6)
//
//Input: nums = [3, 3], target = 6
//Output: [0, 1]
//```
//
//## Function Signature
//
//```go
//func twoSum(nums []int, target int) []int {
//
//}
//```
//
//## Try to solve it two ways:
//1. The brute force way (easy, but slow)
//2. The optimal way (think about what data structure lets you look things up fast)

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
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if seen[nums[i]] == nums[j] {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
			seen[nums[j]] = nums[i]
			seen[nums[i]] = nums[j]
		}
	}
	return nil
}
