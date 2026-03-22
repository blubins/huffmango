package huffmantree

func (h *Huffman) Stat() Stat {
	var numEncodedBits uint64
	for byt, fq := range h.ByteFrequencyTable {
		numEncodedBits += uint64(len(h.EncodingTable[byt])) * fq
	}

	numBits := uint64((*h.F).Size() * 8)
	percentImprovedOverall := (float64(numEncodedBits) / float64(numBits))

	return Stat{
		NumBits:            numBits,
		NumEncodedBits:     numEncodedBits,
		PercentImprovement: 1 - float64(numEncodedBits)/float64(numBits),
		AvgSymbolSize:      percentImprovedOverall * 8,
	}
}
