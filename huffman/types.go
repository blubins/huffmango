package huffmantree

import (
	"huffmango/node"
	"os"
)

type Huffman struct {
	F                  *os.FileInfo
	FilePath           string
	NumEncodings       uint8                // populates after Huffman.InitalizeEncodingTree()
	NumTotBytesWritten uint64               // populates after Huffman.Encode()
	ByteFrequencyTable map[byte]uint64      // populates after Huffman.InitalizeEncodingTree()
	LeafTable          map[node.Node]string // populates after Huffman.InitalizeEncodingTree()
	EncodingTable      map[byte]string      // populates after Huffman.InitalizeEncodingTree()
	Stats              *Stat                // populates after Huffman.Stats()
	Root               *node.Node           // populates after Huffman.InitalizeEncodingTree()
}

type Data struct {
	B    byte
	Freq uint64
}

type Stat struct {
	NumBits                  uint64  // number of actual ingested bits
	NumEncodedBits           uint64  // number of encoded bits
	NumTotBytesWritten       uint64  // encoded bits + translation table size
	TableSize                uint64  // size of the encoder table written into the encoded file
	PercentImprovement       float64 // percent improvement num bits vs. encoded bits
	PercentActualImprovement float64 // percent improvement num bits vs. encoded bits + table size
	AvgSymbolSize            float64 // average byte value encoded size always < 8
	ShannonEntropy           float64 // theoretical minimum average bits needed to encode each byte
}

func New() *Huffman {
	return &Huffman{
		FilePath:           "",
		ByteFrequencyTable: make(map[byte]uint64),
		LeafTable:          make(map[node.Node]string),
		EncodingTable:      make(map[byte]string),
		Root:               nil,
	}
}
