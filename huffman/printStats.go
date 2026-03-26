package huffmantree

import (
	"fmt"
	"time"
)

func (h *Huffman) PrintStats(duration *time.Duration) {
	for b, path := range h.EncodingTable {
		fmt.Printf("B: %-5v Path: %s\n", b, path)
	}
	stats := h.Stat()

	fmt.Printf("%-32s %.4f\n", "Average Symbol Size (Bits)", stats.AvgSymbolSize)
	fmt.Printf("%-32s %d\n", "Input Read (Bits)", stats.NumBits)
	fmt.Printf("%-32s %d\n", "Output Write (Bits)", stats.NumEncodedBits)
	fmt.Printf("%-32s %.4f\n", "Theoretical Percent improvement", stats.PercentImprovement)
	fmt.Printf("%-32s %d\n", "Total bits written", stats.NumTotBytesWritten*8)
	fmt.Printf("%-32s %.4f\n", "Actual Percent improvement", stats.PercentActualImprovement)
	fmt.Printf("%-32s %d\n", "Table Size (Bytes)", stats.TableSize)
	fmt.Printf("%-32s %d\n", "Tree Encodings", uint16(h.NumEncodings)+1)
	fmt.Printf("%-32s %.3f\n", "Duration (ms)", float64(duration.Microseconds())/1000.0)
}
