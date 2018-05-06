package gragh

import (
	"errors"
)

type VetexNode struct {
	content int
	next    *EdgeNode
}

func (vetex *VetexNode) GetConent() int {
	return vetex.content
}

func (vetex *VetexNode) GetNext() *EdgeNode {
	return vetex.next
}

type EdgeNode struct {
	content int
	weight  int
	next    *EdgeNode
}

func (edge *EdgeNode) GetConent() int {
	return edge.content
}

func (edge *EdgeNode) GetNext() *EdgeNode {
	return edge.next
}

func (edge *EdgeNode) GetWeight() int {
	return edge.weight
}

type Gragh struct {
	adjacencyList []*VetexNode
	directioned   bool
	vetexCount    int
}

func NewGragh(vetexCount int, directioned bool) *Gragh {
	adjacencyList := make([]*VetexNode, vetexCount)
	return &Gragh{adjacencyList, directioned, vetexCount}
}

func (gragh *Gragh) GetDirectioned() bool {
	return gragh.directioned
}

func (gragh *Gragh) GetVetexCount() int {
	return gragh.vetexCount
}

func (gragh *Gragh) insertEdge(start int, end int, weight int) error {
	if start > gragh.vetexCount || end > gragh.vetexCount {
		return errors.New("illegal start or end")
	}

	//first vetex
	if gragh.adjacencyList[start] == nil {
		gragh.adjacencyList[start] = &VetexNode{start, &EdgeNode{end, weight, nil}}
	} else {
		edgeNode := gragh.adjacencyList[start].next
		var previous *EdgeNode
		for edgeNode != nil && edgeNode.content < end {
			previous = edgeNode
			edgeNode = edgeNode.next
		}
		//override the same edge node
		if edgeNode != nil && edgeNode.content == end {
			edgeNode.weight = weight
			return nil
		}
		previous.next = &EdgeNode{end, weight, edgeNode}
	}

	return nil
}

func (gragh *Gragh) InsertEdge(start int, end int, weight int) error {
	err := gragh.insertEdge(start, end, weight)
	if err != nil {
		return err
	}
	if !gragh.directioned {
		gragh.insertEdge(end, start, weight)
	}
	return nil
}

func (gragh *Gragh) GetAdjaList() []*VetexNode {
	return gragh.adjacencyList
}
