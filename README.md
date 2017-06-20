# Note: This is not safe for concurrency
# BenchMarks
```
    All benchmarks were done on a 13inch macbook pro with a 2.5 GHz Intel Core i5 processor and 8gb of memory
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
```

# tree
`import "github.com/marcsantiago/GoTree/tree"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package tree implements a basic balanced binary tree




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type Node](#Node)
* [type Tree](#Tree)
  * [func NewTree() *Tree](#NewTree)
  * [func (t *Tree) Add(data int) (err error)](#Tree.Add)
  * [func (t *Tree) CountEdges() (edges int)](#Tree.CountEdges)
  * [func (t *Tree) FindNode(data int) (err error)](#Tree.FindNode)
  * [func (t *Tree) GenerateRandomTree(numberOfNodesToCreate int) (err error)](#Tree.GenerateRandomTree)
  * [func (t *Tree) GetRootData() int](#Tree.GetRootData)
  * [func (t *Tree) GetTreeTotal() int](#Tree.GetTreeTotal)
  * [func (t *Tree) InOrderTraversal()](#Tree.InOrderTraversal)
  * [func (t *Tree) PrintTree()](#Tree.PrintTree)
  * [func (t *Tree) ShiftRoot(newRoot int)](#Tree.ShiftRoot)
  * [func (t *Tree) Sum() (total int)](#Tree.Sum)
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
)
```



## <a name="Node">type</a> [Node](/src/target/tree.go?s=220:277#L3)
``` go
type Node struct {
    Left  *Node
    Right *Node
    Data  int
}
```
Node is a fundemental part of what makes a tree a tree. Many Nodes creates a tree










## <a name="Tree">type</a> [Tree](/src/target/tree.go?s=308:375#L10)
``` go
type Tree struct {
    Root      *Node
    Total     int
    NodeCount int
}
```
Tree basic tree structure







### <a name="NewTree">func</a> [NewTree](/src/target/tree.go?s=693:713#L24)
``` go
func NewTree() *Tree
```
NewTree creates a pointer to the Tree struct





### <a name="Tree.Add">func</a> (\*Tree) [Add](/src/target/tree.go?s=1372:1412#L57)
``` go
func (t *Tree) Add(data int) (err error)
```
Add appends a new Node to a branch in a balanced manner




### <a name="Tree.CountEdges">func</a> (\*Tree) [CountEdges](/src/target/tree.go?s=3713:3752#L180)
``` go
func (t *Tree) CountEdges() (edges int)
```
CountEdges returns the number of edges the tree contains




### <a name="Tree.FindNode">func</a> (\*Tree) [FindNode](/src/target/tree.go?s=805:850#L29)
``` go
func (t *Tree) FindNode(data int) (err error)
```
FindNode recursively looks for the Node with the specified value




### <a name="Tree.GenerateRandomTree">func</a> (\*Tree) [GenerateRandomTree](/src/target/tree.go?s=4460:4532#L217)
``` go
func (t *Tree) GenerateRandomTree(numberOfNodesToCreate int) (err error)
```
GenerateRandomTree uses time (time.Now().Unix()) to create enthorpy for a source of random numbers to append to the Tree




### <a name="Tree.GetRootData">func</a> (\*Tree) [GetRootData](/src/target/tree.go?s=4858:4890#L232)
``` go
func (t *Tree) GetRootData() int
```
GetRootData returns the data stored at the root, however this does not return the root Node




### <a name="Tree.GetTreeTotal">func</a> (\*Tree) [GetTreeTotal](/src/target/tree.go?s=4984:5017#L237)
``` go
func (t *Tree) GetTreeTotal() int
```
GetTreeTotal returns the sum of the collective nodes on the Tree




### <a name="Tree.InOrderTraversal">func</a> (\*Tree) [InOrderTraversal](/src/target/tree.go?s=2064:2097#L95)
``` go
func (t *Tree) InOrderTraversal()
```
InOrderTraversal prints out the values in order




### <a name="Tree.PrintTree">func</a> (\*Tree) [PrintTree](/src/target/tree.go?s=6124:6150#L292)
``` go
func (t *Tree) PrintTree()
```
PrintTree uses json.MarshalIndent() to print the Tree in an organized fashion, which can then be analysized as a JSON
object




### <a name="Tree.ShiftRoot">func</a> (\*Tree) [ShiftRoot](/src/target/tree.go?s=5850:5887#L280)
``` go
func (t *Tree) ShiftRoot(newRoot int)
```
ShiftRoot rebuilds the tree with a new root




### <a name="Tree.Sum">func</a> (\*Tree) [Sum](/src/target/tree.go?s=3061:3093#L144)
``` go
func (t *Tree) Sum() (total int)
```
Sum added up all the values stored in the Nodes.. It is a redundant function because total value is kept as a Tree
value




### <a name="Tree.Traversal">func</a> (\*Tree) [Traversal](/src/target/tree.go?s=2552:2578#L119)
``` go
func (t *Tree) Traversal()
```
Traversal prints out the values by branch side, left, right, ect...




### <a name="Tree.TreeToArray">func</a> (\*Tree) [TreeToArray](/src/target/tree.go?s=5089:5123#L242)
``` go
func (t *Tree) TreeToArray() []int
```
TreeToArray converts to Tree into an int slice








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
