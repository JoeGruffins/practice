# Product of Array Except Self

Given an integer array `nums`, return an array `answer` where `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

You must solve it without using division.

## Examples

```
Input: nums = [1, 2, 3, 4]
Output: [24, 12, 8, 6]

Input: nums = [-1, 1, 0, -3, 3]
Output: [0, 0, 9, 0, 0]
```

## Function Signature

```go
func productExceptSelf(nums []int) []int {

}
```

## Hint

Think about what's to the left of each element and what's to the right.
