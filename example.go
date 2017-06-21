package main

import (
	"fmt"
	"log"

	"./tree"
	// "github.com/marcsantiago/GoTree/tree"
)

func main() {
	// Using time go run example.go, takes
	// go run example.go  235.89s user 11.92s system 174% cpu 2:21.75 total
	// While suming with this function is a bit slower then just calling GetTreeTotal()
	// the real bottle neck is adding that many Nodes to the tree. It uses a lot of memory, so be careful when
	// executing this funcion
	t := tree.NewTree()
	log.Println("Creating a tree with 16777216 nodes")
	t.GenerateRandomTreeIteratively(16777216)
	log.Println("Adding Nodes using the slower Sum() function")
	log.Println("Total:", t.Sum())

	t = tree.NewTree()
	t.AddIteratively(10)
	t.AddIteratively(100)
	t.AddIteratively(2)
	t.AddIteratively(3)
	fmt.Println(t.Sum()) // should be 142
	fmt.Println(t.GetTreeTotal())
	fmt.Println("edges", t.CountEdges())
	fmt.Printf("%+v\n", t.TreeToArray())
	t.PrintTree()

	t = tree.NewTree()
	t.AddIteratively(1)
	t.AddIteratively(2)
	t.AddIteratively(3)
	t.AddIteratively(4)
	t.AddIteratively(10)
	t.AddIteratively(9)
	t.AddIteratively(8)
	t.PrintTree()
}
