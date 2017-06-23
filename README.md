# Note: This is not safe for concurrency
# BenchMarks
```
    BenchmarkTreeToArr1000-4        2000000000               0.00 ns/op
    BenchmarkTreeToArr10000-4       2000000000               0.00 ns/op
    BenchmarkTreeToArr100000-4      2000000000               0.05 ns/op
    BenchmarkTreeToArr1000000-4            1        1244557003 ns/op
    BenchmarkEdgeCount1000-4        2000000000               0.00 ns/op
    BenchmarkEdgeCount10000-4       2000000000               0.01 ns/op
    BenchmarkEdgeCount100000-4      1000000000               0.12 ns/op
    BenchmarkEdgeCount1000000-4            1        1412587095 ns/op
    BenchmarkNewTree1000-4          2000000000               0.00 ns/op
    BenchmarkNewTree10000-4         2000000000               0.00 ns/op
    BenchmarkNewTree1000000-4       2000000000               0.20 ns/op
```


# tree
`import "github.com/marcsantiago/GoTree/tree"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package tree implements a basic balanced binary tree




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func SetMaxDepthDifference(maxDifference float64)](#SetMaxDepthDifference)
* [type Node](#Node)
* [type Tree](#Tree)
  * [func ArrToTree(arr []int) *Tree](#ArrToTree)
  * [func NewTree() *Tree](#NewTree)
  * [func (t *Tree) AddIteratively(data int) (err error)](#Tree.AddIteratively)
  * [func (t *Tree) AddRecusively(data int) (err error)](#Tree.AddRecusively)
  * [func (t *Tree) CountEdges() (edges int)](#Tree.CountEdges)
  * [func (t *Tree) FindNode(data int) (node *Node, err error)](#Tree.FindNode)
  * [func (t *Tree) GenerateRandomTree(numberOfNodesToCreate int) (err error)](#Tree.GenerateRandomTree)
  * [func (t *Tree) GetDepth() float64](#Tree.GetDepth)
  * [func (t *Tree) GetRootData() int](#Tree.GetRootData)
  * [func (t *Tree) GetTotalAndNodeCount() (total int, nodeCount int)](#Tree.GetTotalAndNodeCount)
  * [func (t *Tree) GetTreeTotal() int](#Tree.GetTreeTotal)
  * [func (t *Tree) InOrderTraversal()](#Tree.InOrderTraversal)
  * [func (t *Tree) IsBalanced() bool](#Tree.IsBalanced)
  * [func (t *Tree) PrintTree()](#Tree.PrintTree)
  * [func (t *Tree) Traversal()](#Tree.Traversal)
  * [func (t *Tree) TreeToArray() []int](#Tree.TreeToArray)


#### <a name="pkg-files">Package files</a>
[tree.go](/src/github.com/marcsantiago/GoTree/tree/tree.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    // ErrPositiveIntegers reports that only positive intergers may be added to the tree
    ErrPositiveIntegers = fmt.Errorf("only postive integers may be added")
    // ErrNodeNotFound reports that a Node wasn't found
    ErrNodeNotFound = fmt.Errorf("Node not found")
    // HeightMininum set the depth difference for rebalancing
    HeightMininum = 2.0
)
```


## <a name="SetMaxDepthDifference">func</a> [SetMaxDepthDifference](/src/target/tree.go?s=979:1028#L34)
``` go
func SetMaxDepthDifference(maxDifference float64)
```
SetMaxDepthDifference changes the default from two, which would decrease the balance of the tree, but
allow faster random tree generations.




## <a name="Node">type</a> [Node](/src/target/tree.go?s=236:293#L5)
``` go
type Node struct {
    Left  *Node
    Right *Node
    Data  int
}
```
Node is a fundemental part of what makes a tree a tree. Many Nodes creates a tree










## <a name="Tree">type</a> [Tree](/src/target/tree.go?s=324:391#L12)
``` go
type Tree struct {
    Root      *Node
    Total     int
    NodeCount int
}
```
Tree basic tree structure







### <a name="ArrToTree">func</a> [ArrToTree](/src/target/tree.go?s=8767:8798#L390)
``` go
func ArrToTree(arr []int) *Tree
```
ArrToTree converts an interger slice into a tree, but it won't store the NodeCount or Total


### <a name="NewTree">func</a> [NewTree](/src/target/tree.go?s=789:809#L28)
``` go
func NewTree() *Tree
```
NewTree creates a pointer to the Tree struct





### <a name="Tree.AddIteratively">func</a> (\*Tree) [AddIteratively](/src/target/tree.go?s=3704:3755#L151)
``` go
func (t *Tree) AddIteratively(data int) (err error)
```
AddIteratively appends a new Node to a branch in a balanced manner interatively, which in most cases is faster then
recursion in Go




### <a name="Tree.AddRecusively">func</a> (\*Tree) [AddRecusively](/src/target/tree.go?s=2831:2881#L110)
``` go
func (t *Tree) AddRecusively(data int) (err error)
```
AddRecusively appends a new Node to a branch in a balanced manner recusively




### <a name="Tree.CountEdges">func</a> (\*Tree) [CountEdges](/src/target/tree.go?s=5434:5473#L243)
``` go
func (t *Tree) CountEdges() (edges int)
```
CountEdges returns the number of edges the tree contains




### <a name="Tree.FindNode">func</a> (\*Tree) [FindNode](/src/target/tree.go?s=1133:1190#L39)
``` go
func (t *Tree) FindNode(data int) (node *Node, err error)
```
FindNode recursively looks for the Node with the specified value




### <a name="Tree.GenerateRandomTree">func</a> (\*Tree) [GenerateRandomTree](/src/target/tree.go?s=6181:6253#L280)
``` go
func (t *Tree) GenerateRandomTree(numberOfNodesToCreate int) (err error)
```
GenerateRandomTree uses time (time.Now().Unix()) to create enthorpy for a source of random numbers to append to the Tree




### <a name="Tree.GetDepth">func</a> (\*Tree) [GetDepth](/src/target/tree.go?s=8078:8111#L362)
``` go
func (t *Tree) GetDepth() float64
```
GetDepth gets the maximum depth of the tree




### <a name="Tree.GetRootData">func</a> (\*Tree) [GetRootData](/src/target/tree.go?s=6641:6673#L295)
``` go
func (t *Tree) GetRootData() int
```
GetRootData returns the data stored at the root, however this does not return the root Node




### <a name="Tree.GetTotalAndNodeCount">func</a> (\*Tree) [GetTotalAndNodeCount](/src/target/tree.go?s=1827:1891#L69)
``` go
func (t *Tree) GetTotalAndNodeCount() (total int, nodeCount int)
```
GetTotalAndNodeCount added up all the values stored in the Nodes.. It is a redundant function because total value is kept as a Tree
value




### <a name="Tree.GetTreeTotal">func</a> (\*Tree) [GetTreeTotal](/src/target/tree.go?s=6767:6800#L300)
``` go
func (t *Tree) GetTreeTotal() int
```
GetTreeTotal returns the sum of the collective nodes on the Tree




### <a name="Tree.InOrderTraversal">func</a> (\*Tree) [InOrderTraversal](/src/target/tree.go?s=4504:4537#L195)
``` go
func (t *Tree) InOrderTraversal()
```
InOrderTraversal prints out the values in order




### <a name="Tree.IsBalanced">func</a> (\*Tree) [IsBalanced](/src/target/tree.go?s=8251:8283#L367)
``` go
func (t *Tree) IsBalanced() bool
```
IsBalanced checks to see if the tree is balanced, where the threshold is defaulted to a difference of two




### <a name="Tree.PrintTree">func</a> (\*Tree) [PrintTree](/src/target/tree.go?s=7765:7791#L346)
``` go
func (t *Tree) PrintTree()
```
PrintTree uses json.MarshalIndent() to print the Tree in an organized fashion, which can then be analysized as a JSON
object




### <a name="Tree.Traversal">func</a> (\*Tree) [Traversal](/src/target/tree.go?s=4992:5018#L219)
``` go
func (t *Tree) Traversal()
```
Traversal prints out the values by branch side, left, right, ect...




### <a name="Tree.TreeToArray">func</a> (\*Tree) [TreeToArray](/src/target/tree.go?s=6872:6906#L305)
``` go
func (t *Tree) TreeToArray() []int
```
TreeToArray converts to Tree into an int slice








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
