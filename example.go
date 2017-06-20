package main

import (
	"fmt"

	"./tree"
)

func main() {
	t := tree.NewTree()
	t.GenerateRandomTree()
	t.PrintTree()
	fmt.Println("total:", t.Sum())

	// t := Tree{}
	t.Add(10)
	t.Add(100)
	t.Add(2)
	t.Add(3)

	fmt.Println(t.Sum()) // should be 142
	fmt.Println(t.GetTreeTotal())

	fmt.Println("edges", t.CountEdges())
	fmt.Printf("%+v\n", t.TreeToArray())

}
