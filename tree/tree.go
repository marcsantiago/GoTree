package tree

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
    All bench marks where down on a 13inch macbook pro with a 2.5 GHz Intel Core i5 processor and 8gb of memory

	BenchmarkTreeToArr1000-4        2000000000               0.00 ns/op
	BenchmarkTreeToArr10000-4       2000000000               0.01 ns/op
	BenchmarkTreeToArr100000-4      1000000000               0.26 ns/op
	BenchmarkTreeToArr1000000-4            1        3997007417 ns/op
	BenchmarkEdgeCount1000-4        2000000000               0.00 ns/op
	BenchmarkEdgeCount10000-4       2000000000               0.01 ns/op
	BenchmarkEdgeCount100000-4      2000000000               0.15 ns/op
	BenchmarkEdgeCount1000000-4            1        3882999625 ns/op
	BenchmarkRootShift1000-4        2000000000               0.00 ns/op
	BenchmarkRootShift10000-4       1000000000               0.02 ns/op
	BenchmarkRootShift100000-4      2000000000               0.21 ns/op
	BenchmarkRootShift1000000-4            1        7592046369 ns/op
	BenchmarkNewTree1000-4          2000000000               0.00 ns/op
	BenchmarkNewTree10000-4         2000000000               0.01 ns/op
	BenchmarkNewTree100000-4        2000000000               0.09 ns/op
	BenchmarkNewTree1000000-4              1        3040027533 ns/op
*/

// Node is a fudemental part of what makes a tree a tree. Many Nodes creates a tree
type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

// Tree basic tree structure.. Root is of type Node,
type Tree struct {
	Root      *Node
	Total     int
	NodeCount int
}

var (
	// ErrPositiveIntegers reports that only positive intergers may be added to the tree
	ErrPositiveIntegers = fmt.Errorf("only postive integers may be added")
	// ErrNodeNotFound reports that a Node wasn't found
	ErrNodeNotFound = fmt.Errorf("Node not found")
)

// NewTree ...
func NewTree() *Tree {
	return new(Tree)
}

// FindNode ...
func (t *Tree) FindNode(data int) (err error) {
	newNode := Node{
		Data: data,
	}
	if t.Root != nil {
		if t.findNode(t.Root, newNode) != nil {
			return
		}
	}
	return ErrNodeNotFound
}

func (t *Tree) findNode(search *Node, target Node) *Node {
	var returnNode *Node
	if search == nil {
		return returnNode
	}
	if search.Data == target.Data {
		return search
	}
	returnNode = t.findNode(search.Left, target)
	if returnNode == nil {
		returnNode = t.findNode(search.Right, target)
	}
	return returnNode
}

// Add appends a new node to a branch in a balanced manner
func (t *Tree) Add(data int) (err error) {
	t.Total += data
	t.NodeCount++
	if data < 0 {
		return ErrPositiveIntegers
	}
	NodeToAdd := Node{
		Data: data,
	}
	if t.Root == nil {
		t.Root = new(Node)
	}
	if t.Root.Data == 0 {
		t.Root = &NodeToAdd
		return
	}
	t.add(t.Root, NodeToAdd)
	return
}

func (t *Tree) add(oldNode *Node, newNode Node) {
	if newNode.Data < oldNode.Data {
		if oldNode.Left == nil {
			oldNode.Left = &newNode
		} else {
			t.add(oldNode.Left, newNode)
		}
	} else if newNode.Data > oldNode.Data {
		if oldNode.Right == nil {
			oldNode.Right = &newNode
		} else {
			t.add(oldNode.Right, newNode)
		}
	}
	return
}

// InOrderTraversal prints out the values in order
func (t *Tree) InOrderTraversal() {
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			fmt.Println(currentNode.Data)
		} else {
			t.inOrderTraversal(currentNode)
		}
	}
	return
}

func (t *Tree) inOrderTraversal(n *Node) {
	if n.Left != nil {
		t.inOrderTraversal(n.Left)
	}
	fmt.Println(n.Data)
	if n.Right != nil {
		t.inOrderTraversal(n.Right)
	}
	return
}

// Traversal prints out the values by branch side, left, right, ect...
func (t *Tree) Traversal() {
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			fmt.Println(currentNode.Data)
		} else {
			t.traversal(currentNode)
		}
	}
	return
}

func (t *Tree) traversal(n *Node) {
	fmt.Println(n.Data)
	if n.Left != nil {
		t.traversal(n.Left)
	}
	if n.Right != nil {
		t.traversal(n.Right)
	}
	return
}

// Sum added up all the values stored in the Nodes.. It is a redundant function because total value is kept as a Tree
// value
func (t *Tree) Sum() (total int) {
	var wg sync.WaitGroup
	c := make(chan int, 100)
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			return 1
		}
		wg.Add(1)
		t.sum(currentNode, c, &wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for n := range c {
		total += n
	}
	return total
}

func (t *Tree) sum(n *Node, counter chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n.Left != nil {
		wg.Add(1)
		go t.sum(n.Left, counter, wg)
	}
	counter <- n.Data
	if n.Right != nil {
		wg.Add(1)
		go t.sum(n.Right, counter, wg)
	}
	return
}

// CountEdges returns the number of edges the tree contains
func (t *Tree) CountEdges() (edges int) {
	var wg sync.WaitGroup
	c := make(chan int, 100)
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			return 1
		}
		wg.Add(1)
		t.countEdges(currentNode, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()
	for n := range c {
		edges += n
	}
	return edges
}

func (t *Tree) countEdges(n *Node, counter chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n.Left != nil {
		wg.Add(1)
		go t.countEdges(n.Left, counter, wg)
	}
	counter <- 1
	if n.Right != nil {
		wg.Add(1)
		go t.countEdges(n.Right, counter, wg)
	}
	return
}

// GenerateRandomTree uses time (time.Now().Unix()) to create enthorpy for a source of random numbers to append to the Tree
func (t *Tree) GenerateRandomTree(numberOfNodesToCreate int) (err error) {
	if numberOfNodesToCreate < 0 {
		return ErrPositiveIntegers
	}
	u := time.Now()
	source := rand.NewSource(u.Unix())
	r := rand.New(source)
	arr := r.Perm(numberOfNodesToCreate)
	for _, a := range arr {
		t.Add(a)
	}
	return
}

// GetRootData returns the data stored at the root, however this does not return the root Node
func (t *Tree) GetRootData() int {
	return t.Root.Data
}

// GetTreeTotal returns the sum of the collecitve nodes on the Tree
func (t *Tree) GetTreeTotal() int {
	return t.Total
}

// TreeToArray converts to the into an int slice
func (t *Tree) TreeToArray() []int {
	var wg sync.WaitGroup
	c := make(chan int, 100)
	arr := make([]int, 0, t.NodeCount)
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			return []int{currentNode.Data}
		}
		wg.Add(1)
		t.traversalGetVals(currentNode, c, &wg)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	for n := range c {
		arr = append(arr, n)
	}
	return arr
}

func (t *Tree) traversalGetVals(n *Node, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	if n.Left != nil {
		c <- n.Left.Data
		wg.Add(1)
		go t.traversalGetVals(n.Left, c, wg)
	}
	if n.Right != nil {
		c <- n.Right.Data
		wg.Add(1)
		go t.traversalGetVals(n.Right, c, wg)
	}
	return
}

// ShiftRoot rebuilds the tree with a new root
func (t *Tree) ShiftRoot(newRoot int) {
	arr := t.TreeToArray()
	n := Tree{}
	n.Add(newRoot)
	for _, i := range arr {
		n.Add(i)
	}
	*t = n
}

// PrintTree uses json.MarshalIndent() to print the Tree in an organized fashion, which can then be analysized as a JSON
// object
func (t *Tree) PrintTree() {
	b, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
