package huffmantree

import (
	"bufio"
	"io"
	"os"
)

// main huffman encoding function used to encode data
// returns the amount of bytes written or error if any
func (h *Huffman) Encode(outputPath string) (int, error) {
	// open the input file
	fin, err := os.OpenFile(h.FilePath, os.O_RDONLY, 0)
	if err != nil {
		return 0, err
	}
	defer fin.Close()

	// open and prepare the output file for the encoded binary data
	fout, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer fout.Close()
	// prepare our reader and writer for file io
	reader := bufio.NewReader(fin)
	writer := bufio.NewWriter(fout)

	writer.Write([]byte{0})                    // reserve first byte for padding count
	writer.Write([]byte{byte(h.NumEncodings)}) // 2nd byte reserved for how many paths there are
	bytesWritten := 2
	// next is writing the encoder key table into the file before the encoded data
	// byteValue,encodedByteValue|byteValue,encodedByteValue|...
	for byt, enc := range h.EncodingTable {
		writer.Write([]byte{byt, ','})
		writer.Write([]byte(enc))
		writer.Write([]byte{'|'})
		bytesWritten += 3 + len(enc)
	}

	var bitPos int = 0
	var currByte byte
	for {
		// read in from input file an individual byte
		b, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break // EOF
		}
		// translate the byte into bits from the encoder table
		encB := h.EncodingTable[b]
		for _, r := range encB {
			// iterate through the specific bit position within the byte
			if r == '1' {
				currByte |= 1 << (7 - bitPos)
			}
			bitPos++
			// if the byte is full, then write the byte and iterate to the next byte
			if bitPos == 8 {
				writer.WriteByte(currByte)
				currByte = 0
				bitPos = 0
				bytesWritten++
			}
		}
	}
	// record how many bits are leftover in the last byte written for the padding value
	paddingCount := byte(0)
	if bitPos > 0 {
		paddingCount = byte(8 - bitPos)
		writer.WriteByte(currByte)
		bytesWritten++
	}

	err = writer.Flush()
	if err != nil {
		return 0, err
	}

	_, err = fout.WriteAt([]byte{paddingCount}, 0)
	if err != nil {
		return 0, err
	}

	h.NumTotBytesWritten = uint64(bytesWritten)
	return bytesWritten, nil
}
