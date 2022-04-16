package leetcode

import (
	"sort"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func greaterSum(node *TreeNode, base int) int {
	if node == nil {
		return base
	}
	sum := greaterSum(node.Right, base)
	sum += node.Val
	node.Val = sum
	return greaterSum(node.Left, sum) // return all child nodes sum to upper level
}

func convertBST(root *TreeNode) *TreeNode {
	greaterSum(root, 0)
	return root
}

func iInt(src int) *int {
	return &src
}

// toNodeTree converts sorted []int to a tree
func toNodeTree(src []int) *TreeNode {
	if len(src) == 0 {
		return nil
	}

	sort.Ints(src)

	medIndex := (len(src) - 1) / 2
	root := &TreeNode{
		Val:   src[medIndex],
		Left:  toNodeTree(src[:medIndex]),
		Right: toNodeTree(src[medIndex+1:]),
	}
	return root
}

func equalTrees(expected, actual *TreeNode) bool {
	if expected == actual {
		return true
	}
	if expected == nil || actual == nil {
		return false
	}

	if expected.Val != actual.Val {
		return false
	}
	if !equalTrees(expected.Left, actual.Left) {
		return false
	}
	if !equalTrees(expected.Right, actual.Right) {
		return false
	}
	return true
}

func assertEqualTrees(t *testing.T, expected, actual *TreeNode) {
	if !equalTrees(expected, actual) {
		t.Fatal("trees are not equal")
	}
}

func TestBSTConvertToGreater(t *testing.T) {
	cases := map[string]struct {
		input    *TreeNode
		expected *TreeNode
	}{
		"1": {
			toNodeTree([]int{4, 1, 6, 0, 2, 5, 7, 3, 8}),
			&TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val: 36,
					Left: &TreeNode{
						Val: 36,
					},
					Right: &TreeNode{
						Val: 35,
						Right: &TreeNode{
							Val: 33,
						},
					},
				},
				Right: &TreeNode{
					Val: 21,
					Left: &TreeNode{
						Val: 26,
					},
					Right: &TreeNode{
						Val: 15,
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
	}
	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			assertEqualTrees(t, test.expected, convertBST(test.input))
		})
	}

}
