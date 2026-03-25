package huffmantree

import (
	"huffmango/node"
	"os"
)

type Huffman struct {
	F                  *os.FileInfo
	FilePath           string
	NumEncodings       uint8
	NumTotBytesWritten uint64
	ByteFrequencyTable map[byte]uint64
	LeafTable          map[node.Node]string
	EncodingTable      map[byte]string
	Stats              *Stat // populates after Huffman.Stats() is called
	Root               *node.Node
}

type Data struct {
	B    byte
	Freq uint64
}

type Stat struct {
	NumBits                  uint64 // number of actual ingested bits
	NumEncodedBits           uint64 // number of encoded bits
	NumTotBytesWritten       uint64 // encoded bits + translation table size
	TableSize                uint64
	PercentImprovement       float64 // percent improvement num bits vs. encoded bits
	PercentActualImprovement float64 // percent improvement num bits vs. encoded bits + table size
	AvgSymbolSize            float64 // average byte value encoded size always < 8
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
