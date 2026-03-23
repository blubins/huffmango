package main

import (
	"fmt"
	huffmantree "image-compression/huffman"
	"os"
)

func main() {
	args := os.Args[1:] // strip binary call
	if len(args) < 3 || args[1] != "-o" {
		fmt.Println("Usage: compress <inputfile> -o <outputpath>")
		return
	}

	file := args[0]
	outputPath := args[2]

	Tree := huffmantree.New()
	err := Tree.InitEncodingTable(file)
	if err != nil {
		fmt.Printf("Error initializing encoding table: %s\n", err.Error())
		return
	}

	for b, path := range Tree.EncodingTable {
		fmt.Printf("b: %v, path: %s\n", b, path)
	}

	stats := Tree.Stat()
	fmt.Printf("stats.AvgSymbolSize: %v\n", stats.AvgSymbolSize)
	fmt.Printf("stats.NumBits: %v\n", stats.NumBits)
	fmt.Printf("stats.NumEncodedBits: %v\n", stats.NumEncodedBits)
	fmt.Printf("stats.PercentImprovement: %v\n", stats.PercentImprovement)
	fmt.Printf("Tree.NumEncodings: %v\n", uint16(Tree.NumEncodings)+1)
	err = Tree.Encode(outputPath + "_encoded")
	if err != nil {
		fmt.Printf("Error encoding file: %s\n", err.Error())
	}

	err = huffmantree.Decode(outputPath, outputPath+"_decoded")
	if err != nil {
		fmt.Printf("Error decoding file: %s\n", err.Error())
	}
}
