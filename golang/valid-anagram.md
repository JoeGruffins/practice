# Valid Anagram

Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

An anagram uses all the original letters exactly once, rearranged.

## Examples

```
Input: s = "anagram", t = "nagaram"
Output: true

Input: s = "rat", t = "car"
Output: false

Input: s = "listen", t = "silent"
Output: true
```

## Function Signature

```go
func isAnagram(s string, t string) bool {

}
```

## Try to solve it two ways:
1. The simple way (sort both strings and compare)
2. The optimal way (think about counting)
