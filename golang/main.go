package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(doMaxArea())
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
	seen := make(map[int]int)
	for _, v := range nums {
		seen[v]++
	}
	buckets := make([][]int, len(nums)+1)
	for k, v := range seen {
		buckets[v] = append(buckets[v], k)
	}
	var result []int
	for i := len(buckets) - 1; i >= 0 && len(result) < n; i-- {
		result = append(result, buckets[i]...)
	}
	return result
}

//# Product of Array Except Self
//
//Given an integer array `nums`, return an array `answer` where `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.
//
//You must solve it without using division.
//
//## Examples
//
//```
//Input: nums = [1, 2, 3, 4]
//Output: [24, 12, 8, 6]
//
//Input: nums = [-1, 1, 0, -3, 3]
//Output: [0, 0, 9, 0, 0]
//```
//
//## Function Signature
//
//```go
//func productExceptSelf(nums []int) []int {
//
//}
//```
//
//## Hint
//
//Think about what's to the left of each element and what's to the right.

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
	currentLargest := 0
	set := make(map[int]struct{})
	for _, n := range nums {
		set[n] = struct{}{}
	}
	for k := range set {
		if _, has := set[k-1]; !has {
			delete(set, k)
			i := 1
			for _, has := set[k+i]; has; _, has = set[k+i] {
				delete(set, k+i)
				i++
			}
			if i > currentLargest {
				currentLargest = i
			}
		}
	}
	return currentLargest
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
	isAlpha := func(r byte) bool {
		return r >= 'a' && r <= 'z' || r >= '0' && r <= '9'
	}
	for i, j := 0, len(lower)-1; j > i; {
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

func threeSum(nums []int) [][]int {
	var three [][]int
	slices.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
				continue
			}
			if sum > 0 {
				k--
				continue
			}
			three = append(three, []int{nums[i], nums[j], nums[k]})
			j++
			k--
		}
	}
	return three
}

//# Container With Most Water
//
//Given an integer array `height` of length `n`, where each element represents the height of a vertical line at that position. Find two lines that together with the x-axis form a container that holds the most water.
//
//Return the maximum amount of water the container can store.

//Input: height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
//Output: 49    (between height[1]=8 and height[8]=7, width=7, area=7*7=49)

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
	i, j, best := 0, len(height)-1, 0
	for j > i {
		n := (j - i) * min(height[i], height[j])
		if n > best {
			best = n
		}
		if height[i] > height[j] {
			j--
			continue
		}
		i++
	}
	return best
}
