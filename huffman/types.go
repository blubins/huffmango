package huffmantree

import (
	"image-compression/node"
	"os"
)

type Huffman struct {
	F                  *os.FileInfo
	FilePath           string
	NumEncodings       uint16
	ByteFrequencyTable map[byte]uint64
	LeafTable          map[node.Node]string
	EncodingTable      map[byte]string
	Root               *node.Node
}

type Data struct {
	B    byte   // byte value
	Freq uint64 // frequency
}

type Stat struct {
	NumBits            uint64
	NumEncodedBits     uint64
	PercentImprovement float64
	AvgSymbolSize      float64
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
