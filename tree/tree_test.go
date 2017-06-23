package tree

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	myTree := NewTree()
	myTree.GenerateRandomTreeRecusively(1000)
	sum := myTree.Sum()
	if myTree.Total != sum {
		t.Fatalf("The total values should be the same. They are not.  Tree.Total == %d and Tree.Sum() == %d", myTree.Total, sum)
	}
}

func TestTreeIsTheSame(t *testing.T) {
	tree1 := NewTree()
	tree1.GenerateRandomTreeRecusively(1000)
	tree2 := NewTree()
	tree2.GenerateRandomTreeIteratively(1000)

	b1, err := json.Marshal(tree1)
	if err != nil {
		t.Error(err)
	}

	b2, err := json.Marshal(tree2)
	if err != nil {
		t.Error(err)
	}

	compare := func(b1, b2 []byte) bool {
		var o1 interface{}
		var o2 interface{}
		json.Unmarshal(b1, &o1)
		json.Unmarshal(b2, &o2)
		return reflect.DeepEqual(o1, o2)
	}(b1, b2)

	if !compare {
		t.Errorf("The two trees are not the same and they should be")
	}
}

func BenchmarkTreeToArr1000(b *testing.B) {
	myTree := NewTree()
	myTree.GenerateRandomTreeRecusively(1000)
	b.StartTimer()
	myTree.TreeToArray()
	b.StopTimer()
}

func BenchmarkTreeToArr10000(b *testing.B) {
	myTree := NewTree()
	myTree.GenerateRandomTreeRecusively(10000)
	b.StartTimer()
	myTree.TreeToArray()
	b.StopTimer()
}

func BenchmarkTreeToArr100000(b *testing.B) {
	myTree := NewTree()
	myTree.GenerateRandomTreeRecusively(100000)
	b.StartTimer()
	myTree.TreeToArray()
	b.StopTimer()
}

func BenchmarkTreeToArr1000000(b *testing.B) {
	myTree := NewTree()
	myTree.GenerateRandomTreeRecusively(1000000)
	b.StartTimer()
	myTree.TreeToArray()
	b.StopTimer()
}

func BenchmarkEdgeCount1000(b *testing.B) {
	myTree := NewTree()
	myTree.GenerateRandomTreeRecusively(1000)
	b.StartTimer()
	myTree.CountEdges()
	b.StopTimer()
}

func BenchmarkEdgeCount10000(b *testing.B) {
	myTree := NewTree()
	myTree.GenerateRandomTreeRecusively(10000)
	b.StartTimer()
	myTree.CountEdges()
	b.StopTimer()
}

func BenchmarkEdgeCount100000(b *testing.B) {
	myTree := NewTree()
	myTree.GenerateRandomTreeRecusively(100000)
	b.StartTimer()
	myTree.CountEdges()
	b.StopTimer()
}

func BenchmarkEdgeCount1000000(b *testing.B) {
	myTree := NewTree()
	myTree.GenerateRandomTreeRecusively(1000000)
	b.StartTimer()
	myTree.CountEdges()
	b.StopTimer()
}

func BenchmarkNewTreeRecusively1000(b *testing.B) {
	myTree := NewTree()
	b.StartTimer()
	myTree.GenerateRandomTreeRecusively(1000)
	b.StopTimer()
}

func BenchmarkNewTreeRecusively10000(b *testing.B) {
	myTree := NewTree()
	b.StartTimer()
	myTree.GenerateRandomTreeRecusively(10000)
	b.StopTimer()
}

func BenchmarkNewTreeRecusively100000(b *testing.B) {
	myTree := NewTree()
	b.StartTimer()
	myTree.GenerateRandomTreeRecusively(100000)
	b.StopTimer()
}

func BenchmarkNewTreeRecusively1000000(b *testing.B) {
	myTree := NewTree()
	b.StartTimer()
	myTree.GenerateRandomTreeRecusively(1000000)
	b.StopTimer()
}

func BenchmarkNewTreeIteratively1000(b *testing.B) {
	myTree := NewTree()
	b.StartTimer()
	myTree.GenerateRandomTreeIteratively(1000)
	b.StopTimer()
}

func BenchmarkNewTreeIteratively10000(b *testing.B) {
	myTree := NewTree()
	b.StartTimer()
	myTree.GenerateRandomTreeIteratively(10000)
	b.StopTimer()
}

func BenchmarkNewTreeIteratively100000(b *testing.B) {
	myTree := NewTree()
	b.StartTimer()
	myTree.GenerateRandomTreeIteratively(100000)
	b.StopTimer()
}

func BenchmarkNewTreeIteratively1000000(b *testing.B) {
	myTree := NewTree()
	b.StartTimer()
	myTree.GenerateRandomTreeIteratively(1000000)
	b.StopTimer()
}
