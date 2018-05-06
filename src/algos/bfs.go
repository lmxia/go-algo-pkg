package algos

import (
	"bytes"
	"containers"
	"gragh"
	"strconv"
)

type Bfs struct {
	isMarked []bool
	begin    int
	count    int
	edgeTo   []int
	queue    *containers.Queue
}

func NewBfs(begin int, count int) *Bfs {
	isMarked := make([]bool, count)
	edgeTo := make([]int, count)
	queue := containers.NewQueue()
	return &Bfs{isMarked, begin, 0, edgeTo, queue}
}

func (bfs *Bfs) Bfs(graghObj *gragh.Gragh) {
	adjacencyList := graghObj.GetAdjaList()
	bfs.isMarked[bfs.begin] = true
	bfs.edgeTo[bfs.count] = bfs.begin
	bfs.count++
	bfs.queue.Push(bfs.begin)

	for !bfs.queue.Empty() {
		vetex := adjacencyList[bfs.queue.Pop().(int)]
		for edgeNode := vetex.GetNext(); edgeNode != nil; edgeNode = edgeNode.GetNext() {
			//travelled before
			if !bfs.isMarked[edgeNode.GetConent()] {
				bfs.isMarked[edgeNode.GetConent()] = true
				bfs.edgeTo[bfs.count] = edgeNode.GetConent()
				bfs.count++
				bfs.queue.Push(edgeNode.GetConent())
			}
		}
	}
}

func (dfs *Bfs) GetBfs() string {
	var buffer bytes.Buffer
	for index, value := range dfs.edgeTo {
		if dfs.count == index+1 {
			buffer.WriteString(strconv.Itoa(value))
		} else {
			buffer.WriteString(strconv.Itoa(value) + "->")
		}
	}
	return buffer.String()
}
