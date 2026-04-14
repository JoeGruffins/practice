# Container With Most Water

Given an integer array `height` of length `n`, where each element represents the height of a vertical line at that position. Find two lines that together with the x-axis form a container that holds the most water.

Return the maximum amount of water the container can store.

## Examples

```
Input: height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
Output: 49    (between height[1]=8 and height[8]=7, width=7, area=7*7=49)

Input: height = [1, 1]
Output: 1
```

## Function Signature

```go
func maxArea(height []int) int {

}
```

## Hint

Two pointers, start from both ends. Which pointer should you move inward and why?
