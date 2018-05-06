package main

import (
	"algos"
	"fmt"
	"gragh"
)

func main() {
	graghObj := gragh.NewGragh(9, false)
	graghObj.InsertEdge(0, 1, 10)
	graghObj.InsertEdge(0, 7, 22)
	graghObj.InsertEdge(1, 2, 20)
	graghObj.InsertEdge(1, 4, 89)
	graghObj.InsertEdge(2, 3, 33)
	graghObj.InsertEdge(3, 4, 12)
	graghObj.InsertEdge(3, 5, 32)
	graghObj.InsertEdge(5, 6, 78)
	graghObj.InsertEdge(5, 7, 66)
	graghObj.InsertEdge(7, 8, 13)
	dfsObj := algos.NewDfs(5, 9)
	dfsObj.Dfs(graghObj)
	fmt.Println(dfsObj.GetDfs())

	bfsObj := algos.NewBfs(0, 9)
	bfsObj.Bfs(graghObj)
	fmt.Println(bfsObj.GetBfs())
}
