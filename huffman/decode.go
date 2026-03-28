package huffmantree

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// main decoding function for a Huffman encoded file
// returns an error if any
func Decode(filepath, outputPath string) error {
	if !strings.Contains(filepath, "_encoded") {
		return fmt.Errorf("invalid file extension, expected example.png_huffman, got %s", filepath)
	}
	fin, err := os.OpenFile(filepath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer fin.Close()

	reader := bufio.NewReader(fin)
	// read in the first 2 bytes
	// first byte is reserved for the padding count
	// second byte reserved for how many paths exist
	var reserve [2]byte
	for i := range 2 {
		byt, err := reader.ReadByte()
		if err != nil {
			return err
		}
		reserve[i] = byt
	}
	// parse the encoder key table
	expectedEncodingBars := uint16(reserve[1]) + 1 // |
	// char,01010101|char,1010100011|...|
	numParsedBars := uint16(0)
	decodingTable := make(map[string]byte)
	for expectedEncodingBars > numParsedBars {
		// read byte value
		currByteEnc, err := reader.ReadByte()
		if err != nil {
			return err
		}
		// skip the comma byte
		_, err = reader.ReadByte()
		if err != nil {
			return err
		}
		// read path until '|'
		var currPath []byte
		for {
			b, err := reader.ReadByte()
			if err != nil {
				return err
			}
			if b == '|' {
				break
			}
			currPath = append(currPath, b)
		}
		numParsedBars++
		decodingTable[string(currPath)] = currByteEnc
	}

	fout, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fout.Close()

	writer := bufio.NewWriter(fout)

	paddingCount := reserve[0]
	remaining, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	totalBits := len(remaining)*8 - int(paddingCount)
	var currPath string
	bitIndex := 0
	// decode alg
	for _, b := range remaining {
		for bit := 7; bit >= 0; bit-- {
			if bitIndex >= totalBits {
				break
			}
			if b&(1<<bit) != 0 {
				currPath += "1"
			} else {
				currPath += "0"
			}
			bitIndex++

			if val, ok := decodingTable[currPath]; ok {
				writer.WriteByte(val)
				currPath = ""
			}
		}
	}

	return writer.Flush()
}
