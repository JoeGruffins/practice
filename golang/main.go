package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(doTopKFrequent())
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
	anagrams := make(map[[26]byte][]string)
	aVal := byte('a')
	for _, w := range strs {
		key := [26]byte{}
		for _, r := range w {
			key[byte(r)-aVal]++
		}
		anagrams[key] = append(anagrams[key], w)
	}
	for _, group := range anagrams {
		grouped = append(grouped, group)
	}
	return
}

//# Top K Frequent Elements
//
//Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.
//
//## Examples
//
//```
//Input: nums = [1, 1, 1, 2, 2, 3], k = 2
//Output: [1, 2]
//
//Input: nums = [1], k = 1
//Output: [1]
//
//Input: nums = [4, 4, 4, 1, 1, 2, 2, 2, 3], k = 2
//Output: [4, 2]
//```
//
//## Function Signature
//
//```go
//func topKFrequent(nums []int, k int) []int {
//
//}
//```
//
//## Hint
//
//Count frequencies first. Then think about how to efficiently find the top k without sorting everything.

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

func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, v := range nums {
		seen[v]++
	}
	topValues, topKeys := make([]int, k), make([]int, k)
	for k, v := range seen {
		for i := range topValues {
			if topValues[i] < v {
				copy(topValues[i+1:], topValues[i:])
				topValues[i] = v
				copy(topKeys[i+1:], topKeys[i:])
				topKeys[i] = k
				break
			}
		}
	}
	return topKeys
}
