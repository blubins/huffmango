package utils

import (
	"fmt"
	"slices"
)

func printProgInfo() {
	fmt.Println("Usage: huffman <inputfile> -o <outputpath> [flags]")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("  <inputfile>       path to the file to compress or encoded file to decode")
	fmt.Println("  			         must have _encoded file ext. to decode")
	fmt.Println("  -o <outputpath>   path for the output file")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -viewer                render the huffman tree and compression stats")
	fmt.Println("  -decode                decode a given huffmantree encoded file")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  huffman image.png -o image.png")
	fmt.Println("  huffman image.png -o image.png -viewer")
	fmt.Println("  huffman image.png_encoded -o image.png -decode")
}

func HandleArgv(osArgs []string) *Config {
	args := osArgs[1:] // strip binary call

	if len(args) < 3 || args[1] != "-o" {
		printProgInfo()
		return nil
	}

	return &Config{
		InputFile:          args[0],
		OutputPath:         args[2],
		IsDecode:           slices.Contains(args, "-decode"), // check for -d flag, default encode
		ShouldDisplayStats: slices.Contains(args, "-stats"),  // check for -s flag
		ShouldGenerateHTML: slices.Contains(args, "-viewer"),
	}
}
