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
	fmt.Println("  -r                render the huffman tree and compression stats")
	fmt.Println("  -d                decode a given huffmantree encoded file")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  huffman image.png -o image.png")
	fmt.Println("  huffman image.png -o image.png -r")
	fmt.Println("  huffman image.png_encoded -o image.png -d")
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
		IsDecode:           slices.Contains(args, "-d"), // check for -d flag, default encode
		ShouldDisplayStats: slices.Contains(args, "-s"), // check for -s flag
	}
}
