package containers

import (
	"container/list"
)

type Queue struct {
	content *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (queue *Queue) Pop() interface{} {
	start := queue.content.Front()
	if start != nil {
		queue.content.Remove(start)
		return start.Value
	}
	return nil
}

func (queue *Queue) Push(obj interface{}) {
	queue.content.PushBack(obj)
}

func (queue *Queue) Empty() bool {
	return queue.content.Len() == 0
}

func (queue *Queue) Len() int {
	return queue.content.Len()
}
