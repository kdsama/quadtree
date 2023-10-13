package main

import (
	"fmt"
)

type QuadTree struct {
	lvl      int
	val      int
	x        [2]int // max value of x
	y        [2]int // max value of y
	children []*QuadTree
}

func New(arr [][]int, dp int) *QuadTree {

	var newquadtree func(arr [][]int, depth int, lc, rc, tr, br int) *QuadTree
	newquadtree = func(arr [][]int, depth int, lc, rc, tr, br int) *QuadTree {
		if lc+1 == rc || tr+1 == br {
			return nil
		}
		sum := 0
		for i := tr; i < br; i++ {
			for j := lc; j < rc; j++ {
				sum += arr[i][j]
			}
		}
		qt := &QuadTree{
			val: sum,
			x:   [2]int{tr, br},
			y:   [2]int{lc, rc},
			lvl: dp - depth,
		}
		if depth == 0 {
			return qt
		}
		midR := tr + (br-tr)/2
		midC := lc + (rc-lc)/2
		qt.children = append(qt.children, newquadtree(arr, depth-1, lc, midC, tr, midR))
		qt.children = append(qt.children, newquadtree(arr, depth-1, lc, midC, midR, br))
		qt.children = append(qt.children, newquadtree(arr, depth-1, midC, rc, tr, midR))
		qt.children = append(qt.children, newquadtree(arr, depth-1, midC, rc, midR, br))
		return qt
	}

	return newquadtree(arr, dp, 0, len(arr[0]), 0, len(arr))
}

func (qt *QuadTree) Traverse() {
	var dfs func(node *QuadTree, depth int)
	dfs = func(node *QuadTree, depth int) {
		if node == nil {
			return
		}
		fmt.Println("Node value-->", node.val, "at depth", depth, node.x, node.y)
		for _, n := range node.children {
			dfs(n, depth+1)
		}
	}
	dfs(qt, 0)
}

func (qt QuadTree) FindRegions(value int) []*QuadTree {
	arr := []*QuadTree{}
	var dfs func(node *QuadTree, depth int)
	dfs = func(node *QuadTree, depth int) {
		if node == nil {
			return
		}
		if node.val < value {
			return
		}
		flag := false
		for _, child := range node.children {
			fmt.Println("For ", node, "child value::", child.val)
			if child.val > value {

				dfs(child, depth+1)
				flag = true
			}
		}

		if !flag {
			fmt.Println("Flag", flag, "for ", node)
			arr = append(arr, node)
		}
	}
	dfs(&qt, 0)
	for _, v := range arr {
		fmt.Println(v.val, v.x, v.y)
	}
	return arr
}

func (qt *QuadTree) Add(value, x, y int) {
	var dfs func(node *QuadTree)
	dfs = func(node *QuadTree) {
		if node == nil {
			return
		}
		node.val += value
		for _, n := range node.children {
			if n == nil {
				continue
			}
			if n.x[0] <= x && n.x[1] >= x {
				if n.y[0] <= y && n.y[1] >= y {
					dfs(n)
				}
			}

		}
	}
	dfs(qt)
}

func main() {

	arr := newFunction(256)
	qt := New(arr, 4)
	qt.FindRegions(256)
	// qt.Traverse()
	// qt.Add(5, 2, 5)
	// qt.Traverse()
}

func newFunction(size int) [][]int {
	x := size
	arr := make([][]int, x)
	for i := range arr {
		arr[i] = make([]int, x)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			arr[i][j] = 1

		}
	}
	return arr
}
