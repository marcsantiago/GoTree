// Package tree implements a basic balanced binary tree
package tree

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Node is a fundemental part of what makes a tree a tree. Many Nodes creates a tree
type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

// Tree basic tree structure
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
	// HeightMininum set the depth difference for rebalancing
	HeightMininum = 2.0
)

// NewTree creates a pointer to the Tree struct
func NewTree() *Tree {
	return new(Tree)
}

// SetMaxDepthDifference changes the default from two, which would decrease the balance of the tree, but
// allow faster random tree generations.
func SetMaxDepthDifference(maxDifference float64) {
	HeightMininum = maxDifference
}

// FindNode recursively looks for the Node with the specified value
func (t *Tree) FindNode(data int) (node *Node, err error) {
	newNode := Node{
		Data: data,
	}
	if t.Root != nil {
		node := t.findNode(t.Root, newNode)
		if node != nil {
			return node, nil
		}
	}
	return nil, ErrNodeNotFound
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

// GetTotalAndNodeCount added up all the values stored in the Nodes.. It is a redundant function because total value is kept as a Tree
// value
func (t *Tree) GetTotalAndNodeCount() (total int, nodeCount int) {
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			return 0, 1
		}
		total, nodeCount = t.sum(currentNode)
	}
	return total, nodeCount
}

// Instead of using recursive or channels, I decided to demonstrate another way to tranverse the tree..
// for an channel Example see the CountEdges() method,
func (t *Tree) sum(n *Node) (int, int) {
	total := 0
	nodeCount := 0
	// create a starting stack
	var stack []*Node
	for n != nil {
		// push
		stack = append(stack, n)
		n = n.Left
	}
	for len(stack) > 0 {
		// pop a value out
		// https://github.com/golang/go/wiki/SliceTricks
		n, stack = stack[len(stack)-1], stack[:len(stack)-1]
		total += n.Data
		nodeCount++
		if n.Right != nil {
			n = n.Right
			for n != nil {
				stack = append(stack, n)
				n = n.Left
			}
		}
	}
	return total, nodeCount
}

// AddRecusively appends a new Node to a branch in a balanced manner recusively
func (t *Tree) AddRecusively(data int) (err error) {
	t.Total += data
	t.NodeCount++
	if data < 0 {
		return ErrPositiveIntegers
	}
	NodeToAdd := Node{Data: data}
	if t.Root == nil {
		t.Root = &NodeToAdd
		return
	}
	t.addRecusively(t.Root, NodeToAdd)
	if t.Root != nil {
		newTree := rebalance(t)
		if newTree.Root != nil {
			*t = newTree
		}

	}
	return
}

func (t *Tree) addRecusively(oldNode *Node, newNode Node) {
	if newNode.Data < oldNode.Data {
		if oldNode.Left == nil {
			oldNode.Left = &newNode
		} else {
			t.addRecusively(oldNode.Left, newNode)
		}
	} else if newNode.Data > oldNode.Data {
		if oldNode.Right == nil {
			oldNode.Right = &newNode
		} else {
			t.addRecusively(oldNode.Right, newNode)
		}
	}
	return
}

// AddIteratively appends a new Node to a branch in a balanced manner interatively, which in most cases is faster then
// recursion in Go
func (t *Tree) AddIteratively(data int) (err error) {
	t.Total += data
	t.NodeCount++
	if data < 0 {
		return ErrPositiveIntegers
	}
	NodeToAdd := Node{Data: data}
	if t.Root == nil {
		t.Root = &NodeToAdd
		return
	}
	t.addIteratively(t.Root, NodeToAdd)
	if t.Root != nil {
		newTree := rebalance(t)
		if newTree.Root != nil {
			*t = newTree
		}

	}
	return
}

func (t *Tree) addIteratively(oldNode *Node, newNode Node) {
	for {
		if newNode.Data < oldNode.Data {
			if oldNode.Left == nil {
				oldNode.Left = &newNode
				break
			} else {
				oldNode = oldNode.Left
			}
		} else if newNode.Data > oldNode.Data {
			if oldNode.Right == nil {
				oldNode.Right = &newNode
				break
			} else {
				oldNode = oldNode.Right
			}
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
	arr := r.Perm(numberOfNodesToCreate + 1)
	newTree := ArrToTree(arr)
	*t = *newTree
	t.Total, t.NodeCount = newTree.GetTotalAndNodeCount()
	return
}

// GetRootData returns the data stored at the root, however this does not return the root Node
func (t *Tree) GetRootData() int {
	return t.Root.Data
}

// GetTreeTotal returns the sum of the collective nodes on the Tree
func (t *Tree) GetTreeTotal() int {
	return t.Total
}

// TreeToArray converts to Tree into an int slice
func (t *Tree) TreeToArray() []int {
	var wg sync.WaitGroup
	c := make(chan int, 100)
	arr := make([]int, 0, t.NodeCount)
	if t.Root != nil {
		currentNode := t.Root
		if currentNode.Left == nil && currentNode.Right == nil {
			return []int{currentNode.Data}
		}
		// send the root node
		c <- currentNode.Data
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

// PrintTree uses json.MarshalIndent() to print the Tree in an organized fashion, which can then be analysized as a JSON
// object
func (t *Tree) PrintTree() {
	b, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func getDepth(root *Node) float64 {
	if root == nil {
		return 0
	}
	return math.Max(getDepth(root.Left), getDepth(root.Right)) + 1
}

// GetDepth gets the maximum depth of the tree
func (t *Tree) GetDepth() float64 {
	return getDepth(t.Root)
}

// IsBalanced checks to see if the tree is balanced, where the threshold is defaulted to a difference of two
func (t *Tree) IsBalanced() bool {
	if t.Root == nil {
		return true
	}
	heightDiff := getDepth(t.Root.Left) - getDepth(t.Root.Right)
	if math.Abs(heightDiff) > HeightMininum {
		return false
	}
	return true
}

func rebalance(t *Tree) (newTree Tree) {
	if !t.IsBalanced() {
		list := t.TreeToArray()
		newTree := ArrToTree(list)
		newTree.Total = t.Total
		newTree.NodeCount = t.NodeCount
		return *newTree
	}
	return
}

// ArrToTree converts an interger slice into a tree, but it won't store the NodeCount or Total
func ArrToTree(arr []int) *Tree {
	sort.Ints(arr)
	t := new(Tree)
	t.Root = arrToTree(arr, 0, len(arr)-1)
	return t
}

func arrToTree(arr []int, start, end int) *Node {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	n := new(Node)
	n.Data = arr[mid]
	n.Left = arrToTree(arr, start, mid-1)
	n.Right = arrToTree(arr, mid+1, end)
	return n
}
