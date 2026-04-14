package huffmantree

import (
	"huffmango/linkedlist"
	"huffmango/node"
	"strings"
)

// creates a map of leaf node pointers to binary paths
func getLeafTableHelper(root *node.Node, stack *linkedlist.LinkedList, key *map[node.Node]string) {
	if root == nil {
		return
	}
	// leaf found
	if root.Left == nil && root.Right == nil {
		pathToLeaf := stack.String()
		path := strings.ReplaceAll(pathToLeaf, ",", "")
		(*key)[*root] = path
	}
	// left branch taken
	stack.Append(0)
	getLeafTableHelper(root.Left, stack, key)
	// popping the stack
	stack.DeleteNode(int(stack.Size()))
	// right branch taken
	stack.Append(1)
	getLeafTableHelper(root.Right, stack, key)
	stack.DeleteNode(int(stack.Size()))
}

func GetLeafTable(root *node.Node) map[node.Node]string {
	if root == nil {
		return nil
	}
	stack := linkedlist.New()
	key := make(map[node.Node]string)
	getLeafTableHelper(root, stack, &key)
	return key
}
