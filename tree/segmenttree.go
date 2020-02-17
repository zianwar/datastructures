package tree

type SegmentTreeNode struct {
	start, end int
	sum        int // can be sum, min or max
	left       *SegmentTreeNode
	right      *SegmentTreeNode
}

type SegmentTree struct {
	root *SegmentTreeNode
}

func NewSegmentTree(start, end int, values []int) *SegmentTree {
	root := buildTree(start, end, values)
	return &SegmentTree{
		root: root,
	}
}

// BuildTree builds the segment tree given start and end and the values.
// O(N)
func buildTree(start, end int, values []int) *SegmentTreeNode {
	if start == end {
		return &SegmentTreeNode{
			start: start,
			end:   end,
			sum:   values[start],
		}
	}

	mid := start + (end-start)/2
	left := buildTree(start, mid, values)
	right := buildTree(mid+1, end, values)

	return &SegmentTreeNode{
		start: start,
		end:   end,
		sum:   left.sum + right.sum,
		left:  left,
		right: right,
	}
}

// updateTree updates tree's sum with the given value and index
// O(log(N)) (similar to binary search)
func updateTree(root *SegmentTreeNode, index int, val int) {
	if root.start == index && root.end == index {
		root.sum = val
		return
	}

	mid := root.start + (root.end-root.start)/2

	if index <= mid {
		updateTree(root.left, index, val)
	} else {
		updateTree(root.left, index, val)
	}

	// important to update the root's sum at the end.
	root.sum = root.left.sum + root.right.sum
}

func querySum(root *SegmentTreeNode, i, j int) int {
	if root.start == i && root.end == j {
		return root.sum
	}

	mid := root.start + (root.end-root.start)/2

	if i <= mid {
		return querySum(root.left, i, j)
	} else if i > mid {
		return querySum(root.right, i, j)
	} else {
		return querySum(root.left, i, mid) + querySum(root.right, mid+1, j)
	}
}
