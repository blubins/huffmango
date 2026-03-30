package huffmantree

// returns the stats of a Huffman encoding function
// Huffman.Encode() must be executed before trying to call this or it will fail
func (h *Huffman) Stat() Stat {
	var numEncodedBits uint64
	var totalSymbols uint64
	for byt, fq := range h.ByteFrequencyTable {
		numEncodedBits += uint64(len(h.EncodingTable[byt])) * fq
		totalSymbols += fq
	}

	numBits := uint64((*h.F).Size() * 8)

	// each entry is = 1 (byte value) + 1 (comma) + len(path) + 1 (pipe)
	var tableSize uint64
	for _, enc := range h.EncodingTable {
		tableSize += uint64(3 + len(enc))
	}

	s := Stat{
		NumBits:                  numBits,
		NumEncodedBits:           numEncodedBits,
		NumTotBytesWritten:       h.NumTotBytesWritten,
		TableSize:                tableSize + 2, // 2 reserve bytes
		PercentImprovement:       1 - float64(numEncodedBits)/float64(numBits),
		PercentActualImprovement: 1 - float64(h.NumTotBytesWritten*8)/float64(numBits),
		AvgSymbolSize:            float64(numEncodedBits) / float64(totalSymbols),
	}

	h.Stats = &s
	return s
}
