# Group Anagrams

Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

## Examples

```
Input: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
Output: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]

Input: strs = [""]
Output: [[""]]

Input: strs = ["a"]
Output: [["a"]]
```

## Function Signature

```go
func groupAnagrams(strs []string) [][]string {

}
```

## Hint

How can you create a key that is the same for all anagrams of a word?
