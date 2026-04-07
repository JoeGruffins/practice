# Top K Frequent Elements

Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.

## Examples

```
Input: nums = [1, 1, 1, 2, 2, 3], k = 2
Output: [1, 2]

Input: nums = [1], k = 1
Output: [1]

Input: nums = [4, 4, 4, 1, 1, 2, 2, 2, 3], k = 2
Output: [4, 2]
```

## Function Signature

```go
func topKFrequent(nums []int, k int) []int {

}
```

## Hint

Count frequencies first. Then think about how to efficiently find the top k without sorting everything.
