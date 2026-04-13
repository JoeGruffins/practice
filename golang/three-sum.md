# Three Sum

Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

The solution must not contain duplicate triplets.

## Examples

```
Input: nums = [-1, 0, 1, 2, -1, -4]
Output: [[-1, -1, 2], [-1, 0, 1]]

Input: nums = [0, 1, 1]
Output: []

Input: nums = [0, 0, 0]
Output: [[0, 0, 0]]
```

## Function Signature

```go
func threeSum(nums []int) [][]int {

}
```

## Hint

Sort first. Then for each number, use two pointers on the remaining elements - like Two Sum but with a sorted array.
