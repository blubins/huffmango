package huffmantree

import (
	"image-compression/node"
)

func InOrderTraversal(root *node.Node, v *[]*node.Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left, v)
	*v = append(*v, root)
	InOrderTraversal(root.Right, v)
}
