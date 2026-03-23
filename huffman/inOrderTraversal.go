package huffmantree

import (
	"huffmango/node"
)

func InOrderTraversal(root *node.Node, v *[]*node.Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left, v)
	*v = append(*v, root)
	InOrderTraversal(root.Right, v)
}
