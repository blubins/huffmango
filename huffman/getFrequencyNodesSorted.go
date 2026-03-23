package huffmantree

import "sort"

func (h *Huffman) getFrequencyNodesSorted() []Data {
	var out []Data
	var numEncodings uint16
	for k, v := range h.ByteFrequencyTable {
		out = append(out, Data{
			B:    k,
			Freq: v,
		})
		numEncodings++
	}
	// subtract 1 so 255 = 256 to fit inside uint8
	h.NumEncodings = uint8(numEncodings - 1)
	// sort the slice into a min prio queue
	sort.Slice(out, func(i, j int) bool {
		return out[i].Freq < out[j].Freq
	})

	return out
}
