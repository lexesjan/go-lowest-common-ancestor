package binarytree

import (
	"go-lowest-common-ancestor/binarytree/treenode"
	"testing"
)

func getTestBinaryTreeRoot() *treenode.TreeNode {
	/*
	       0
	      / \
	     1   2
	    / \
	   3   4
	        \
	         5
	*/
	testBinaryTreeRoot := treenode.New(0)
	testBinaryTreeRoot.Left = treenode.New(1)
	testBinaryTreeRoot.Right = treenode.New(2)
	testBinaryTreeRoot.Left.Left = treenode.New(3)
	testBinaryTreeRoot.Left.Right = treenode.New(4)
	testBinaryTreeRoot.Left.Right.Right = treenode.New(5)

	return testBinaryTreeRoot
}

var testBinaryTreeRoot = getTestBinaryTreeRoot()

type LowestCommonAncestorTest struct {
	message    string
	binaryTree *BinaryTree
	node1      *treenode.TreeNode
	node2      *treenode.TreeNode
	output     *treenode.TreeNode
}

var LowestCommonAncestorTests = []LowestCommonAncestorTest{
	// null cases
	{"LCA should be null", &BinaryTree{Root: nil}, nil, nil, nil},
	// base cases
	{
		"LCA should be 0",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left,
		testBinaryTreeRoot.Right,
		testBinaryTreeRoot,
	},
	{
		"LCA should be 0",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left.Left,
		testBinaryTreeRoot.Right,
		testBinaryTreeRoot,
	},
	{
		"LCA should be 0",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left.Right.Right,
		testBinaryTreeRoot.Right,
		testBinaryTreeRoot,
	},
	{
		"LCA should be 1",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left.Left,
		testBinaryTreeRoot.Left.Right.Right,
		testBinaryTreeRoot.Left,
	},
	{
		"LCA should be 1",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left.Left,
		testBinaryTreeRoot.Left.Right,
		testBinaryTreeRoot.Left,
	},
	// lca is one of the input nodes cases
	{
		"LCA should be 0",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left,
		testBinaryTreeRoot,
		testBinaryTreeRoot,
	},
	{
		"LCA should be 1",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left.Left,
		testBinaryTreeRoot.Left,
		testBinaryTreeRoot.Left,
	},
	{
		"LCA should be 1",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left,
		testBinaryTreeRoot.Left.Right,
		testBinaryTreeRoot.Left,
	},
	// lca is both of the input nodes cases
	{
		"LCA should be 0",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot,
		testBinaryTreeRoot,
		testBinaryTreeRoot,
	},
	{
		"LCA should be 3",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left.Left,
		testBinaryTreeRoot.Left.Left,
		testBinaryTreeRoot.Left.Left,
	},
	// node is not in the graph case
	{
		"LCA should be nil",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left.Left,
		treenode.New(6),
		nil,
	},
	{
		"LCA should be nil",
		&BinaryTree{Root: testBinaryTreeRoot},
		testBinaryTreeRoot.Left.Left,
		nil,
		nil,
	},
	{
		"LCA should be nil",
		&BinaryTree{Root: testBinaryTreeRoot},
		treenode.New(6),
		treenode.New(6),
		nil,
	},
}

func TestLowestCommonAncestor(t *testing.T) {
	for _, test := range LowestCommonAncestorTests {
		output := LowestCommonAncestor(test.binaryTree, test.node1, test.node2)
		expected := test.output
		if output != expected {
			t.Errorf("%s but output was %d", test.message, expected.Value)
		}
	}
}
