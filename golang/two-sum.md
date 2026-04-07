# Two Sum

Given an array of integers `nums` and an integer `target`, return the indices of the two numbers that add up to `target`.

- Each input has exactly one solution
- You can't use the same element twice
- Return the answer in any order

## Examples

```
Input: nums = [2, 7, 11, 15], target = 9
Output: [0, 1]   (because 2 + 7 = 9)

Input: nums = [3, 2, 4], target = 6
Output: [1, 2]   (because 2 + 4 = 6)

Input: nums = [3, 3], target = 6
Output: [0, 1]
```

## Function Signature

```go
func twoSum(nums []int, target int) []int {

}
```

## Try to solve it two ways:
1. The brute force way (easy, but slow)
2. The optimal way (think about what data structure lets you look things up fast)
