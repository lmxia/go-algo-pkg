package algos

import (
	"bytes"
	"gragh"
	"strconv"
)

type Dfs struct {
	isMarked []bool
	begin    int
	count    int
	edgeTo   []int
}

func NewDfs(begin int, count int) *Dfs {
	isMarked := make([]bool, count)
	edgeTo := make([]int, count)
	return &Dfs{isMarked, begin, 0, edgeTo}
}

func (dfs *Dfs) Dfs(graghObj *gragh.Gragh) {
	adjacencyList := graghObj.GetAdjaList()
	dfs.isMarked[dfs.begin] = true
	vetex := adjacencyList[dfs.begin]
	dfs.edgeTo[dfs.count] = dfs.begin
	dfs.count++
	var edgeNode *gragh.EdgeNode
	for edgeNode = vetex.GetNext(); edgeNode != nil; edgeNode = edgeNode.GetNext() {
		//travelled before
		if !dfs.isMarked[edgeNode.GetConent()] {
			dfs.begin = edgeNode.GetConent()
			dfs.Dfs(graghObj)
		}
	}
}

func (dfs *Dfs) GetDfs() string {
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
