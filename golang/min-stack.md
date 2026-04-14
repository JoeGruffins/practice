# Min Stack

Design a stack that supports push, pop, top, and retrieving the minimum element, all in O(1) time.

Implement the following:
- `Push(val)` - pushes the element onto the stack
- `Pop()` - removes the element on top of the stack
- `Top()` - gets the top element
- `GetMin()` - retrieves the minimum element in the stack

## Examples

```
Push(-2)
Push(0)
Push(-3)
GetMin() → -3
Pop()
Top() → 0
GetMin() → -2
```

## Type Signature

```go
type MinStack struct {

}

func (s *MinStack) Push(val int)
func (s *MinStack) Pop()
func (s *MinStack) Top() int
func (s *MinStack) GetMin() int
```

## Hint

How do you track the minimum as elements are pushed and popped? One stack isn't enough.
