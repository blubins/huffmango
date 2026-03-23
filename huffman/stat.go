package huffmantree

func (h *Huffman) Stat() Stat {
	var numEncodedBits uint64
	for byt, fq := range h.ByteFrequencyTable {
		numEncodedBits += uint64(len(h.EncodingTable[byt])) * fq
	}

	numBits := uint64((*h.F).Size() * 8)
	percentImprovedOverall := (float64(numEncodedBits) / float64(numBits))

	// each entry is = 1 (byte value) + 1 (comma) + len(path) + 1 (pipe)
	var tableSize uint64
	for _, enc := range h.EncodingTable {
		tableSize += uint64(3 + len(enc))
	}

	return Stat{
		NumBits:                  numBits,
		NumEncodedBits:           numEncodedBits,
		NumTotBytesWritten:       h.NumTotBytesWritten,
		TableSize:                tableSize + 2, // 2 reserve bytes
		PercentImprovement:       1 - float64(numEncodedBits)/float64(numBits),
		PercentActualImprovement: 1 - float64(h.NumTotBytesWritten*8)/float64(numBits),
		AvgSymbolSize:            percentImprovedOverall * 8,
	}
}
