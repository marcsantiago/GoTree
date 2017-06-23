# Note: This is not safe for concurrency
# BenchMarks
```
    BenchmarkTreeToArr1000-4                2000000000               0.00 ns/op
    BenchmarkTreeToArr10000-4               2000000000               0.01 ns/op
    BenchmarkTreeToArr100000-4              2000000000               0.19 ns/op
    BenchmarkTreeToArr1000000-4                    1        5391724267 ns/op
    BenchmarkEdgeCount1000-4                2000000000               0.00 ns/op
    BenchmarkEdgeCount10000-4               1000000000               0.02 ns/op
    BenchmarkEdgeCount100000-4              2000000000               0.18 ns/op
    BenchmarkEdgeCount1000000-4                    1        3909434842 ns/op
    BenchmarkRootShift1000-4                2000000000               0.00 ns/op
    BenchmarkRootShift10000-4               2000000000               0.01 ns/op
    BenchmarkRootShift100000-4              2000000000               0.12 ns/op
    BenchmarkRootShift1000000-4                    1        4984063181 ns/op
    BenchmarkNewTreeRecusively1000-4        2000000000               0.00 ns/op
    BenchmarkNewTreeRecusively10000-4       2000000000               0.01 ns/op
    BenchmarkNewTreeRecusively100000-4      2000000000               0.15 ns/op
    BenchmarkNewTreeRecusively1000000-4            1        4218410453 ns/op
    BenchmarkNewTreeIteratively1000-4       2000000000               0.00 ns/op
    BenchmarkNewTreeIteratively10000-4      2000000000               0.00 ns/op
    BenchmarkNewTreeIteratively100000-4     2000000000               0.07 ns/op
    BenchmarkNewTreeIteratively1000000-4           1        2182573527 ns/op
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
  * [func (t *Tree) Rebalance()](#Tree.Rebalance)
  * [func (t *Tree) ShiftRoot(newRoot int)](#Tree.ShiftRoot)
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







### <a name="ArrToTree">func</a> [ArrToTree](/src/target/tree.go?s=8630:8661#L387)
``` go
func ArrToTree(arr []int) *Tree
```
ArrToTree converts an interger slice into a tree


### <a name="NewTree">func</a> [NewTree](/src/target/tree.go?s=789:809#L28)
``` go
func NewTree() *Tree
```
NewTree creates a pointer to the Tree struct





### <a name="Tree.AddIteratively">func</a> (\*Tree) [AddIteratively](/src/target/tree.go?s=2647:2716#L110)
``` go
func (t *Tree) AddIteratively(data int, balanceTree bool) (err error)
```
AddIteratively appends a new Node to a branch in a balanced manner interatively, which in most cases is faster then
recursion in Go




### <a name="Tree.AddRecusively">func</a> (\*Tree) [AddRecusively](/src/target/tree.go?s=1763:1831#L68)
``` go
func (t *Tree) AddRecusively(data int, balanceTree bool) (err error)
```
AddRecusively appends a new Node to a branch in a balanced manner recusively




### <a name="Tree.CountEdges">func</a> (\*Tree) [CountEdges](/src/target/tree.go?s=4388:4427#L203)
``` go
func (t *Tree) CountEdges() (edges int)
```
CountEdges returns the number of edges the tree contains




### <a name="Tree.FindNode">func</a> (\*Tree) [FindNode](/src/target/tree.go?s=1133:1190#L39)
``` go
func (t *Tree) FindNode(data int) (node *Node, err error)
```
FindNode recursively looks for the Node with the specified value




### <a name="Tree.GenerateRandomTreeIteratively">func</a> (\*Tree) [GenerateRandomTreeIteratively](/src/target/tree.go?s=5629:5712#L256)
``` go
func (t *Tree) GenerateRandomTreeIteratively(numberOfNodesToCreate int) (err error)
```
GenerateRandomTreeIteratively uses time (time.Now().Unix()) to create enthorpy for a source of random numbers to append to the Tree




### <a name="Tree.GenerateRandomTreeRecusively">func</a> (\*Tree) [GenerateRandomTreeRecusively](/src/target/tree.go?s=5145:5227#L240)
``` go
func (t *Tree) GenerateRandomTreeRecusively(numberOfNodesToCreate int) (err error)
```
GenerateRandomTreeRecusively uses time (time.Now().Unix()) to create enthorpy for a source of random numbers to append to the Tree




### <a name="Tree.GetDepth">func</a> (\*Tree) [GetDepth](/src/target/tree.go?s=7817:7850#L356)
``` go
func (t *Tree) GetDepth() float64
```
GetDepth gets the maximum depth of the tree




### <a name="Tree.GetRootData">func</a> (\*Tree) [GetRootData](/src/target/tree.go?s=6075:6107#L272)
``` go
func (t *Tree) GetRootData() int
```
GetRootData returns the data stored at the root, however this does not return the root Node




### <a name="Tree.GetTreeTotal">func</a> (\*Tree) [GetTreeTotal](/src/target/tree.go?s=6201:6234#L277)
``` go
func (t *Tree) GetTreeTotal() int
```
GetTreeTotal returns the sum of the collective nodes on the Tree




### <a name="Tree.InOrderTraversal">func</a> (\*Tree) [InOrderTraversal](/src/target/tree.go?s=3458:3491#L155)
``` go
func (t *Tree) InOrderTraversal()
```
InOrderTraversal prints out the values in order




### <a name="Tree.IsBalanced">func</a> (\*Tree) [IsBalanced](/src/target/tree.go?s=7990:8022#L361)
``` go
func (t *Tree) IsBalanced() bool
```
IsBalanced checks to see if the tree is balanced, where the threshold is defaulted to a difference of two




### <a name="Tree.PrintTree">func</a> (\*Tree) [PrintTree](/src/target/tree.go?s=7504:7530#L340)
``` go
func (t *Tree) PrintTree()
```
PrintTree uses json.MarshalIndent() to print the Tree in an organized fashion, which can then be analysized as a JSON
object




### <a name="Tree.Rebalance">func</a> (\*Tree) [Rebalance](/src/target/tree.go?s=8325:8351#L373)
``` go
func (t *Tree) Rebalance()
```
Rebalance converts the tree into an array, sorts the array, creates a new tree from that array, and assigns it's pointer




### <a name="Tree.ShiftRoot">func</a> (\*Tree) [ShiftRoot](/src/target/tree.go?s=7067:7104#L320)
``` go
func (t *Tree) ShiftRoot(newRoot int)
```
ShiftRoot rebuilds the tree with a new root




### <a name="Tree.Traversal">func</a> (\*Tree) [Traversal](/src/target/tree.go?s=3946:3972#L179)
``` go
func (t *Tree) Traversal()
```
Traversal prints out the values by branch side, left, right, ect...




### <a name="Tree.TreeToArray">func</a> (\*Tree) [TreeToArray](/src/target/tree.go?s=6306:6340#L282)
``` go
func (t *Tree) TreeToArray() []int
```
TreeToArray converts to Tree into an int slice








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
