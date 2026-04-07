package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(doGroupAnagrams())
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

//# Group Anagrams
//
//Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.
//
//## Examples
//
//```
//Input: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//Output: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
//
//Input: strs = [""]
//Output: [[""]]
//
//Input: strs = ["a"]
//Output: [["a"]]
//```
//
//## Function Signature
//
//```go
//func groupAnagrams(strs []string) [][]string {
//
//}
//```
//
//## Hint
//
//How can you create a key that is the same for all anagrams of a word?

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
	zVal := byte('z')
	for _, w := range strs {
		key := [26]byte{}
		for _, r := range w {
			key[zVal-byte(r)]++
		}
		anagrams[key] = append(anagrams[key], w)
	}
	for _, group := range anagrams {
		grouped = append(grouped, group)
	}
	return
}
