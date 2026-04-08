# Longest Consecutive Sequence

Given an unsorted array of integers `nums`, return the length of the longest consecutive elements sequence.

Must run in O(n) time.

## Examples

```
Input: nums = [100, 4, 200, 1, 3, 2]
Output: 4    (the sequence is [1, 2, 3, 4])

Input: nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
Output: 9    (the sequence is [0, 1, 2, 3, 4, 5, 6, 7, 8])

Input: nums = []
Output: 0
```

## Function Signature

```go
func longestConsecutive(nums []int) int {

}
```

## Hint

Put everything in a set. A number is the start of a sequence if num-1 is NOT in the set.
