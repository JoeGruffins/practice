# Trapping Rain Water

Given `n` non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

## Examples

```
Input: height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
Output: 6

Input: height = [4, 2, 0, 3, 2, 5]
Output: 9
```

## Function Signature

```go
func trap(height []int) int {

}
```

## Hint

For each position, the water level is determined by min(max height to its left, max height to its right) minus its own height.
