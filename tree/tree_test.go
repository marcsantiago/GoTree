package tree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/marcsantiago/GoTree/tree"
)

func TestSum(t *testing.T) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(1000)
	sum := myTree.Sum()
	if myTree.Total != sum {
		t.Fatalf("The total values should be the same. They are not.  Tree.Total == %d and Tree.Sum() == %d", myTree.Total, sum)
	}
}

func TestNewRoot(t *testing.T) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(1000)
	currentNode := myTree.Root
	myTree.ShiftRoot(5)
	newNode := myTree.Root
	if currentNode == newNode {
		t.Fatalf("Tree roots should be different")
	}
}

func BenchmarkTreeToArr1000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(1000)
	b.StartTimer()
	myTree.TreeToArray()
	b.StopTimer()
}

func BenchmarkTreeToArr10000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(10000)
	b.StartTimer()
	myTree.TreeToArray()
	b.StopTimer()
}

func BenchmarkTreeToArr100000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(100000)
	b.StartTimer()
	myTree.TreeToArray()
	b.StopTimer()
}

func BenchmarkTreeToArr1000000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(1000000)
	b.StartTimer()
	myTree.TreeToArray()
	b.StopTimer()
}

func BenchmarkEdgeCount1000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(1000)
	b.StartTimer()
	myTree.CountEdges()
	b.StopTimer()
}

func BenchmarkEdgeCount10000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(10000)
	b.StartTimer()
	myTree.CountEdges()
	b.StopTimer()
}

func BenchmarkEdgeCount100000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(100000)
	b.StartTimer()
	myTree.CountEdges()
	b.StopTimer()
}

func BenchmarkEdgeCount1000000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(1000000)
	b.StartTimer()
	myTree.CountEdges()
	b.StopTimer()
}

func BenchmarkRootShift1000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(1000)
	u := time.Now()
	source := rand.NewSource(u.Unix())
	r := rand.New(source)
	arr := r.Perm(1)
	b.StartTimer()
	myTree.ShiftRoot(arr[0])
	b.StopTimer()
}

func BenchmarkRootShift10000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(10000)
	u := time.Now()
	source := rand.NewSource(u.Unix())
	r := rand.New(source)
	arr := r.Perm(1)
	b.StartTimer()
	myTree.ShiftRoot(arr[0])
	b.StopTimer()
}

func BenchmarkRootShift100000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(100000)
	u := time.Now()
	source := rand.NewSource(u.Unix())
	r := rand.New(source)
	arr := r.Perm(1)
	b.StartTimer()
	myTree.ShiftRoot(arr[0])
	b.StopTimer()
}

func BenchmarkRootShift1000000(b *testing.B) {
	myTree := tree.NewTree()
	myTree.GenerateRandomTree(1000000)
	u := time.Now()
	source := rand.NewSource(u.Unix())
	r := rand.New(source)
	arr := r.Perm(1)
	b.StartTimer()
	myTree.ShiftRoot(arr[0])
	b.StopTimer()
}

func BenchmarkNewTree1000(b *testing.B) {
	myTree := tree.NewTree()
	b.StartTimer()
	myTree.GenerateRandomTree(1000)
	b.StopTimer()
}

func BenchmarkNewTree10000(b *testing.B) {
	myTree := tree.NewTree()
	b.StartTimer()
	myTree.GenerateRandomTree(10000)
	b.StopTimer()
}

func BenchmarkNewTree100000(b *testing.B) {
	myTree := tree.NewTree()
	b.StartTimer()
	myTree.GenerateRandomTree(100000)
	b.StopTimer()
}

func BenchmarkNewTree1000000(b *testing.B) {
	myTree := tree.NewTree()
	b.StartTimer()
	myTree.GenerateRandomTree(1000000)
	b.StopTimer()
}
