# Note: This is not safe for concurrency
# BenchMarks
```
    BenchmarkTreeToArr1000-4                2000000000               0.00 ns/op
    BenchmarkTreeToArr10000-4               2000000000               0.01 ns/op
    BenchmarkTreeToArr100000-4              2000000000               0.17 ns/op
    BenchmarkTreeToArr1000000-4                    1        3490666994 ns/op
    BenchmarkEdgeCount1000-4                2000000000               0.00 ns/op
    BenchmarkEdgeCount10000-4               2000000000               0.01 ns/op
    BenchmarkEdgeCount100000-4              2000000000               0.12 ns/op
    BenchmarkEdgeCount1000000-4                    1        3609856257 ns/op
    BenchmarkNewTreeRecusively1000-4        2000000000               0.00 ns/op
    BenchmarkNewTreeRecusively10000-4       2000000000               0.01 ns/op
    BenchmarkNewTreeRecusively100000-4      2000000000               0.15 ns/op
    BenchmarkNewTreeRecusively1000000-4            1        4484977592 ns/op
    BenchmarkNewTreeIteratively1000-4       2000000000               0.00 ns/op
    BenchmarkNewTreeIteratively10000-4      2000000000               0.00 ns/op
    BenchmarkNewTreeIteratively100000-4     2000000000               0.08 ns/op
    BenchmarkNewTreeIteratively1000000-4           1        2226629506 ns/op
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
  * [func (t *Tree) AddIteratively(data int, balanceTree bool) (err error)](#Tree.AddIteratively)
  * [func (t *Tree) AddRecusively(data int, balanceTree bool) (err error)](#Tree.AddRecusively)
  * [func (t *Tree) CountEdges() (edges int)](#Tree.CountEdges)
  * [func (t *Tree) FindNode(data int) (node *Node, err error)](#Tree.FindNode)
  * [func (t *Tree) GenerateRandomTreeIteratively(numberOfNodesToCreate int) (err error)](#Tree.GenerateRandomTreeIteratively)
  * [func (t *Tree) GenerateRandomTreeRecusively(numberOfNodesToCreate int) (err error)](#Tree.GenerateRandomTreeRecusively)
  * [func (t *Tree) GetDepth() float64](#Tree.GetDepth)
  * [func (t *Tree) GetRootData() int](#Tree.GetRootData)
  * [func (t *Tree) GetTreeTotal() int](#Tree.GetTreeTotal)
  * [func (t *Tree) InOrderTraversal()](#Tree.InOrderTraversal)
  * [func (t *Tree) IsBalanced() bool](#Tree.IsBalanced)
  * [func (t *Tree) PrintTree()](#Tree.PrintTree)
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







### <a name="ArrToTree">func</a> [ArrToTree](/src/target/tree.go?s=9168:9199#L427)
``` go
func ArrToTree(arr []int) *Tree
```
ArrToTree converts an interger slice into a tree


### <a name="NewTree">func</a> [NewTree](/src/target/tree.go?s=789:809#L28)
``` go
func NewTree() *Tree
```
NewTree creates a pointer to the Tree struct





### <a name="Tree.AddIteratively">func</a> (\*Tree) [AddIteratively](/src/target/tree.go?s=3453:3522#L153)
``` go
func (t *Tree) AddIteratively(data int, balanceTree bool) (err error)
```
AddIteratively appends a new Node to a branch in a balanced manner interatively, which in most cases is faster then
recursion in Go




### <a name="Tree.AddRecusively">func</a> (\*Tree) [AddRecusively](/src/target/tree.go?s=2482:2550#L105)
``` go
func (t *Tree) AddRecusively(data int, balanceTree bool) (err error)
```
AddRecusively appends a new Node to a branch in a balanced manner recusively




### <a name="Tree.CountEdges">func</a> (\*Tree) [CountEdges](/src/target/tree.go?s=5281:5320#L252)
``` go
func (t *Tree) CountEdges() (edges int)
```
CountEdges returns the number of edges the tree contains




### <a name="Tree.FindNode">func</a> (\*Tree) [FindNode](/src/target/tree.go?s=1133:1190#L39)
``` go
func (t *Tree) FindNode(data int) (node *Node, err error)
```
FindNode recursively looks for the Node with the specified value




### <a name="Tree.GenerateRandomTreeIteratively">func</a> (\*Tree) [GenerateRandomTreeIteratively](/src/target/tree.go?s=6605:6688#L312)
``` go
func (t *Tree) GenerateRandomTreeIteratively(numberOfNodesToCreate int) (err error)
```
GenerateRandomTreeIteratively uses time (time.Now().Unix()) to create enthorpy for a source of random numbers to append to the Tree




### <a name="Tree.GenerateRandomTreeRecusively">func</a> (\*Tree) [GenerateRandomTreeRecusively](/src/target/tree.go?s=6038:6120#L289)
``` go
func (t *Tree) GenerateRandomTreeRecusively(numberOfNodesToCreate int) (err error)
```
GenerateRandomTreeRecusively uses time (time.Now().Unix()) to create enthorpy for a source of random numbers to append to the Tree




### <a name="Tree.GetDepth">func</a> (\*Tree) [GetDepth](/src/target/tree.go?s=8522:8555#L399)
``` go
func (t *Tree) GetDepth() float64
```
GetDepth gets the maximum depth of the tree




### <a name="Tree.GetRootData">func</a> (\*Tree) [GetRootData](/src/target/tree.go?s=7133:7165#L334)
``` go
func (t *Tree) GetRootData() int
```
GetRootData returns the data stored at the root, however this does not return the root Node




### <a name="Tree.GetTreeTotal">func</a> (\*Tree) [GetTreeTotal](/src/target/tree.go?s=7259:7292#L339)
``` go
func (t *Tree) GetTreeTotal() int
```
GetTreeTotal returns the sum of the collective nodes on the Tree




### <a name="Tree.InOrderTraversal">func</a> (\*Tree) [InOrderTraversal](/src/target/tree.go?s=4351:4384#L204)
``` go
func (t *Tree) InOrderTraversal()
```
InOrderTraversal prints out the values in order




### <a name="Tree.IsBalanced">func</a> (\*Tree) [IsBalanced](/src/target/tree.go?s=8695:8727#L404)
``` go
func (t *Tree) IsBalanced() bool
```
IsBalanced checks to see if the tree is balanced, where the threshold is defaulted to a difference of two




### <a name="Tree.PrintTree">func</a> (\*Tree) [PrintTree](/src/target/tree.go?s=8209:8235#L383)
``` go
func (t *Tree) PrintTree()
```
PrintTree uses json.MarshalIndent() to print the Tree in an organized fashion, which can then be analysized as a JSON
object




### <a name="Tree.Sum">func</a> (\*Tree) [Sum](/src/target/tree.go?s=1810:1842#L69)
``` go
func (t *Tree) Sum() (total int)
```
Sum added up all the values stored in the Nodes.. It is a redundant function because total value is kept as a Tree
value




### <a name="Tree.Traversal">func</a> (\*Tree) [Traversal](/src/target/tree.go?s=4839:4865#L228)
``` go
func (t *Tree) Traversal()
```
Traversal prints out the values by branch side, left, right, ect...




### <a name="Tree.TreeToArray">func</a> (\*Tree) [TreeToArray](/src/target/tree.go?s=7364:7398#L344)
``` go
func (t *Tree) TreeToArray() []int
```
TreeToArray converts to Tree into an int slice








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
