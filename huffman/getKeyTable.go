package huffmantree

import (
	"huffmango/node"
)

// converts a node table to a byte value to binary path table
func GetKeyTable(leafTable map[node.Node]string) map[byte]string {
	out := make(map[byte]string)
	for node, path := range leafTable {
		out[node.Data.(Data).B] = path
	}

	return out
}
