package stack

import (
	"container/list"
)

type Stack struct {
	content *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (stack *Stack) Pop() interface{} {
	end := stack.content.Back()
	if end != nil {
		stack.content.Remove(end)
		return end
	}
	return nil
}

func (stack *Stack) Push(obj interface{}) {
	stack.content.PushBack(obj)
}

func (stack *Stack) Peak() interface{} {
	end := stack.content.Back()
	if end != nil {
		return end
	}
	return nil
}

func (stack *Stack) Empty() bool {
	return stack.content.Len() == 0
}

func (stack *Stack) Len() int {
	return stack.content.Len()
}
