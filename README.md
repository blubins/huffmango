## huffman-go

Golang [Huffman coding](https://en.wikipedia.org/wiki/Huffman_coding) compression CLI and library with interactive HTML tree visualization.

### Requirements

- Go 1.21+

### Installation

```bash
git clone https://github.com/blubins/huffmango.git
cd huffmango
go build -o huffmango .
```

### Usage

#### Encode a file
```bash
huffmango myfile.png -o myfile.png
```

### Encode with stats and tree visualization

```bash
huffmango myfile.png -o myfile.png -stats -viewer
```
This prints compression statistics and generates an interactive `huffman_tree.html` to the same name and directory as the output parameter.

### Decode a file
#### To decode the file must have the suffix _encoded or the file will be rejected.

```bash
huffmango myfile.png_encoded -o myfile.png -decode
```

### Library Usage

```go
package main

import (
	huffmantree "huffmango/huffman"
	"huffmango/renderer"
	"fmt"
)

func main() {
	begin := time.Now()

	// create a new tree
	h := huffmantree.New()
	
	// Build the Huffman tree from a file any format
	err := h.InitEncodingTable("input.txt")
	if err != nil {
		panic(err)
	}

	// Encode the file
	_, err = h.Encode("input.txt_encoded")
	if err != nil {
		panic(err)
	}
	duration := time.Since(begin)

	// Calculate and display the compression stats with duration taken
	stats := h.Stat()
	h.PrintStats(&duration)



	// Generate interactive HTML tree visualization html page
	renderer.CreateHTMLView(h.Root, &stats, "huffman_tree.html")

	// Decode
	huffmantree.Decode("input.txt_encoded", "input_decoded.txt")
}
```

### How Huffman coding works

1. **Byte frequency analysis** - Count how often each byte value appears in the input file
2. **Tree construction algorithm**

    **i.**  Create a leaf node for each symbol and insert into a priority queue
    
    **ii.** Create a new internal node with these two nodes as children with the value as the combine frequency of the children

    **iii.** Repeat until there is one node left (the root)

3. **Path assignment** - Traverse the tree to assign variable length binary paths (left = 0, right = 1). Frequent bytes get shorter codes
4. **Encoding** - Replace each byte with its binary code and pack the bits into bytes
5. **Decoding** - Read the encoding table from the file header, then read each bit and translate them into the original byte representation

### Encoded file format
#### Every encoded file has a header containing the required information to decode it
| Section | Size | Description |
|---------|------|-------------|
| Padding count | 1 byte | Number of padding bits in the last byte |
| Encoding count | 1 byte | Number of unique byte values - 1 |
| Encoding table | Variable | `byte,code\|byte,code\|...` (3 + code length bytes per entry) |
| Encoded data | Variable | Packed bits |
#### The size of the encoding table has a maximum of 256 entries as there are only 256 unique byte values possible 0->255

### Stats

When using the `-stats` flag with the encoder, the following statistics are displayed to the console:

| Stat | Description |
|------|-------------|
| NumBits | Original file size in bits |
| NumEncodedBits | Total bits in the encoded data (excluding table overhead) |
| TableSize | Size of the encoding table in bytes |
| NumTotBytesWritten | Total output size (encoded data + table) |
| PercentImprovement | Compression ratio of encoded bits vs original bits |
| PercentActualImprovement | Compression ratio including table overhead |
| AvgSymbolSize | Average bits per byte in the encoded output, always <= 8 for effective compression |

### HTML Tree Viewer

The `-viewer` flag generates an interactive HTML file with the encoders statistics and tree visualization:
- Pan and zoom (mouse drag + scroll wheel)
- Huffman tree visualization showing byte values and frequencies
- Compression statistics overlay