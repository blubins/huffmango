package huffmantree

import (
	"fmt"
	"huffmango/frequency"
	"huffmango/linkedlist"
	"huffmango/node"
	"os"
)

// initializes the huffman encoding table required for encoding data
func (h *Huffman) InitEncodingTable(filePath string) error {
	h.FilePath = filePath

	stat, err := os.Stat(h.FilePath)
	if err != nil {
		return err
	}
	h.F = &stat
	// catch edge case if the file is empty
	if stat.Size() == 0 {
		return fmt.Errorf("empty file")
	}

	data, err := os.ReadFile(h.FilePath)
	if err != nil {
		return err
	}

	// build the frequency byte table
	tbl := frequency.GetByteFrequency(&data)
	h.ByteFrequencyTable = tbl

	frequencyNodes := h.getFrequencyNodesSorted()
	// huffman tree construction
	// create a linked list we will treat as a queue to dequeue from
	nodeTreeList := linkedlist.New()
	// fill our queue up with our frequency nodes
	for _, fN := range frequencyNodes {
		nodeTreeList.Append(
			&node.Node{
				Data: fN,
			},
		)
	}

	for nodeTreeList.Size() > 1 {
		// huffman algorithm:
		// dequeue 2 nodes with lowest frequency
		// create new internal node with the two removed children as nodes
		// sum of their frequencies as new frequency
		// enqueue new node sorted into the queue

		// grab the first two nodes, list is sorted so they have lowest freq
		ele1 := nodeTreeList.DeleteNode(1)
		ele2 := nodeTreeList.DeleteNode(1)
		// grab their data element and type cast to node.Node ptrs
		node1 := ele1.Data.(*node.Node)
		node2 := ele2.Data.(*node.Node)
		// extract the data field to get their freq values
		data1 := node1.Data.(Data)
		data2 := node2.Data.(Data)
		// combine the frequencies
		combineFrequency := data1.Freq + data2.Freq
		// create a new internal node with the combine frequency
		internalNode := &node.Node{
			Data: Data{
				Freq: combineFrequency,
				B:    0, // placeholder 0 value instead of byte value since internal node
			},
			Left:  node1,
			Right: node2,
		}
		// insert back into the queue but sorted by freq
		nodeTreeList.InsertSorted(
			internalNode,
			func(a, b any) bool {
				return a.(*node.Node).Data.(Data).Freq < b.(*node.Node).Data.(Data).Freq
			},
		)

	}

	// last node left is the root of the huffman tree
	if nodeTreeList.Size() > 0 {
		h.Root = nodeTreeList.DeleteNode(1).Data.(*node.Node)
		h.LeafTable = GetLeafTable(h.Root)
		// only 1 byte value edge case
		// check if there is only 1 leaf in the entire table assign it 0 path
		if len(h.LeafTable) == 1 {
			for k := range h.LeafTable {
				h.LeafTable[k] = "0"
			}
		}

		h.EncodingTable = GetKeyTable(h.LeafTable)
	}

	return nil
}
