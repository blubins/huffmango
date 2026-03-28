package huffmantree

import (
	"huffmango/node"
)

// simple inorder traversal of a tree
func InOrderTraversal(root *node.Node, v *[]*node.Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left, v)
	*v = append(*v, root)
	InOrderTraversal(root.Right, v)
}
