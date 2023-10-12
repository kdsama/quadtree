package main

type QuadTree struct {
	val      int
	x        [2]int // max value of x
	y        [2]int // max value of y
	children []*QuadTree
}

func New(arr [][]int, depth int) *QuadTree {

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

	return newquadtree(arr, depth, 0, len(arr[0]), 0, len(arr))
}

func (qt *QuadTree) Traverse() {
	var dfs func(node *QuadTree, depth int)
	dfs = func(node *QuadTree, depth int) {
		if node == nil {
			return
		}

		for _, n := range node.children {
			dfs(n, depth+1)
		}
	}
	dfs(qt, 0)
}

func main() {

	x := 256
	arr := make([][]int, x)
	for i := range arr {
		arr[i] = make([]int, x+37)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < x+37; j++ {
			arr[i][j] = 1

		}
	}
	qt := New(arr, 256)
	qt.Traverse()
}