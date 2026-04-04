package huffmantree

import "math"

// calcualtes the theoretical minimum average bits needed to encode each byte
// as described in C. E. Shannon's paper
// "A Mathematical Theory of Communication"
// https://en.wikipedia.org/wiki/Entropy_(information_theory)
func (h *Huffman) ShannonEntropy() float64 {
	entropy := 0.0
	for _, freq := range h.ByteFrequencyTable {
		p := float64(freq) / float64((*h.F).Size())
		entropy -= p * math.Log2(p)
	}
	return entropy
}
