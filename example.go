package main

import (
	"fmt"
	"log"

	"./tree"
	// "github.com/marcsantiago/GoTree/tree"
)

func main() {
	// Using time go run example.go, takes
	// go run example.go  81.86s user 11.17s system 195% cpu 47.536 total
	// While suming with this function is a slower then just calling GetTreeTotal()
	// the real bottle neck is adding that many Nodes to the tree. It uses a lot of memory, so be careful when
	// executing this funcion
	t := tree.NewTree()
	log.Println("Creating a tree with 16777216 nodes")
	t.GenerateRandomTreeIteratively(16777216)
	log.Println("Adding Nodes using the slower Sum() function")
	log.Println("Total:", t.Sum())

	t = tree.NewTree()
	t.AddIteratively(10, false)
	t.AddIteratively(100, false)
	t.AddIteratively(2, false)
	t.AddIteratively(3, false)
	fmt.Println(t.Sum()) // should be 142
	fmt.Println(t.GetTreeTotal())
	fmt.Println("edges", t.CountEdges())
	fmt.Printf("%+v\n", t.TreeToArray())
	t.PrintTree()

	t = tree.NewTree()
	t.AddIteratively(1, false)
	t.AddIteratively(2, false)
	t.AddIteratively(3, false)
	t.AddIteratively(4, false)
	t.AddIteratively(10, false)
	t.AddIteratively(9, false)
	t.AddIteratively(8, false)
	t.PrintTree()
	fmt.Println(t.GetDepth())
	fmt.Println(t.IsBalanced())
	t.AddIteratively(20, true)
	fmt.Println(t.GetDepth())
	fmt.Println(t.IsBalanced())
	t.PrintTree()

}
