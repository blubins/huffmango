package huffmantree

import "sort"

func (h *Huffman) getFrequencyNodesSorted() []Data {
	var out []Data
	for k, v := range h.ByteFrequencyTable {
		out = append(out, Data{
			B:    k,
			Freq: v,
		})
		h.NumEncodings++
	}
	// sort the slice into a min prio queue
	sort.Slice(out, func(i, j int) bool {
		return out[i].Freq < out[j].Freq
	})

	return out
}
