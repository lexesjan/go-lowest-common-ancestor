package treenode

// TreeNode type
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// New TreeNode
func New(value int) *TreeNode {
	result := TreeNode{Value: value}
	return &result
}
