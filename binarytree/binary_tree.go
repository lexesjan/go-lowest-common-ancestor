package binarytree

import "go-lowest-common-ancestor/binarytree/treenode"

// BinaryTree type
type BinaryTree struct {
	Root *treenode.TreeNode
}

func lowestCommonAncestor(
	root *treenode.TreeNode,
	node1 *treenode.TreeNode,
	node2 *treenode.TreeNode,
	result **treenode.TreeNode,
) bool {
	if root == nil {
		return false
	}

	rootIsEitherNode := root.Value == node1.Value || root.Value == node2.Value
	left := lowestCommonAncestor(root.Left, node1, node2, result)
	right := lowestCommonAncestor(root.Right, node1, node2, result)

	if left && right || rootIsEitherNode && (left || right) {
		*result = root
	}

	return left || right || rootIsEitherNode
}

// LowestCommonAncestor returns the lowest common ancestor if node1 and node2 in the binary tree given
func LowestCommonAncestor(tree *BinaryTree, node1, node2 *treenode.TreeNode) *treenode.TreeNode {
	if node1 == nil || node2 == nil {
		return nil
	}

	if node1.Value == node2.Value {
		return node1
	}

	var result **treenode.TreeNode = new(*treenode.TreeNode)
	lowestCommonAncestor(tree.Root, node1, node2, result)

	return *result
}
