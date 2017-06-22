# Note: This is not safe for concurrency
# BenchMarks
```
    BenchmarkTreeToArr1000-4                2000000000               0.00 ns/op
    BenchmarkTreeToArr10000-4               2000000000               0.01 ns/op
    BenchmarkTreeToArr100000-4              1000000000               0.25 ns/op
    BenchmarkTreeToArr1000000-4                    1        4040629920 ns/op
    BenchmarkEdgeCount1000-4                2000000000               0.00 ns/op
    BenchmarkEdgeCount10000-4               2000000000               0.01 ns/op
    BenchmarkEdgeCount100000-4              2000000000               0.12 ns/op
    BenchmarkEdgeCount1000000-4                    1        4207954738 ns/op
    BenchmarkRootShift1000-4                2000000000               0.00 ns/op
    BenchmarkRootShift10000-4               2000000000               0.00 ns/op
    BenchmarkRootShift100000-4              2000000000               0.09 ns/op
    BenchmarkRootShift1000000-4                    1        3081776385 ns/op
    BenchmarkNewTreeRecusively1000-4        2000000000               0.00 ns/op
    BenchmarkNewTreeRecusively10000-4       2000000000               0.00 ns/op
    BenchmarkNewTreeRecusively100000-4      2000000000               0.11 ns/op
    BenchmarkNewTreeRecusively1000000-4            1        3445230256 ns/op
    BenchmarkNewTreeIteratively1000-4       2000000000               0.00 ns/op
    BenchmarkNewTreeIteratively10000-4      2000000000               0.00 ns/op
    BenchmarkNewTreeIteratively100000-4     2000000000               0.02 ns/op
    BenchmarkNewTreeIteratively1000000-4           1        1028531304 ns/op
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
  * [func (t *Tree) AddIteratively(data int) (err error)](#Tree.AddIteratively)
  * [func (t *Tree) AddRecusively(data int) (err error)](#Tree.AddRecusively)
  * [func (t *Tree) CountEdges() (edges int)](#Tree.CountEdges)
  * [func (t *Tree) FindNode(data int) (node *Node, err error)](#Tree.FindNode)
  * [func (t *Tree) GenerateRandomTreeIteratively(numberOfNodesToCreate int) (err error)](#Tree.GenerateRandomTreeIteratively)
  * [func (t *Tree) GenerateRandomTreeRecusively(numberOfNodesToCreate int) (err error)](#Tree.GenerateRandomTreeRecusively)
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





### <a name="Tree.AddIteratively">func</a> (\*Tree) [AddIteratively](/src/target/tree.go?s=2264:2315#L97)
``` go
func (t *Tree) AddIteratively(data int) (err error)
```
AddIteratively appends a new Node to a branch in a balanced manner interatively, which in most cases is faster then
recursion in Go




### <a name="Tree.AddRecusively">func</a> (\*Tree) [AddRecusively](/src/target/tree.go?s=1435:1485#L58)
``` go
func (t *Tree) AddRecusively(data int) (err error)
```
AddRecusively appends a new Node to a branch in a balanced manner recusively




### <a name="Tree.CountEdges">func</a> (\*Tree) [CountEdges](/src/target/tree.go?s=4669:4708#L224)
``` go
func (t *Tree) CountEdges() (edges int)
```
CountEdges returns the number of edges the tree contains




### <a name="Tree.FindNode">func</a> (\*Tree) [FindNode](/src/target/tree.go?s=805:862#L29)
``` go
func (t *Tree) FindNode(data int) (node *Node, err error)
```
FindNode recursively looks for the Node with the specified value




### <a name="Tree.GenerateRandomTreeIteratively">func</a> (\*Tree) [GenerateRandomTreeIteratively](/src/target/tree.go?s=5884:5967#L276)
``` go
func (t *Tree) GenerateRandomTreeIteratively(numberOfNodesToCreate int) (err error)
```
GenerateRandomTreeIteratively uses time (time.Now().Unix()) to create enthorpy for a source of random numbers to append to the Tree




### <a name="Tree.GenerateRandomTreeRecusively">func</a> (\*Tree) [GenerateRandomTreeRecusively](/src/target/tree.go?s=5426:5508#L261)
``` go
func (t *Tree) GenerateRandomTreeRecusively(numberOfNodesToCreate int) (err error)
```
GenerateRandomTreeRecusively uses time (time.Now().Unix()) to create enthorpy for a source of random numbers to append to the Tree




### <a name="Tree.GetRootData">func</a> (\*Tree) [GetRootData](/src/target/tree.go?s=6304:6336#L291)
``` go
func (t *Tree) GetRootData() int
```
GetRootData returns the data stored at the root, however this does not return the root Node




### <a name="Tree.GetTreeTotal">func</a> (\*Tree) [GetTreeTotal](/src/target/tree.go?s=6430:6463#L296)
``` go
func (t *Tree) GetTreeTotal() int
```
GetTreeTotal returns the sum of the collective nodes on the Tree




### <a name="Tree.InOrderTraversal">func</a> (\*Tree) [InOrderTraversal](/src/target/tree.go?s=3020:3053#L139)
``` go
func (t *Tree) InOrderTraversal()
```
InOrderTraversal prints out the values in order




### <a name="Tree.PrintTree">func</a> (\*Tree) [PrintTree](/src/target/tree.go?s=7590:7616#L351)
``` go
func (t *Tree) PrintTree()
```
PrintTree uses json.MarshalIndent() to print the Tree in an organized fashion, which can then be analysized as a JSON
object




### <a name="Tree.ShiftRoot">func</a> (\*Tree) [ShiftRoot](/src/target/tree.go?s=7296:7333#L339)
``` go
func (t *Tree) ShiftRoot(newRoot int)
```
ShiftRoot rebuilds the tree with a new root




### <a name="Tree.Sum">func</a> (\*Tree) [Sum](/src/target/tree.go?s=4017:4049#L188)
``` go
func (t *Tree) Sum() (total int)
```
Sum added up all the values stored in the Nodes.. It is a redundant function because total value is kept as a Tree
value




### <a name="Tree.Traversal">func</a> (\*Tree) [Traversal](/src/target/tree.go?s=3508:3534#L163)
``` go
func (t *Tree) Traversal()
```
Traversal prints out the values by branch side, left, right, ect...




### <a name="Tree.TreeToArray">func</a> (\*Tree) [TreeToArray](/src/target/tree.go?s=6535:6569#L301)
``` go
func (t *Tree) TreeToArray() []int
```
TreeToArray converts to Tree into an int slice








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
