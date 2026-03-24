package main

import (
	"fmt"
	huffmantree "huffmango/huffman"
	"huffmango/renderer"
	"huffmango/utils"
	"os"
	"time"
)

func main() {
	conf := utils.HandleArgv(os.Args)
	if conf == nil {
		return
	}

	// if encode
	if !conf.IsDecode {
		begin := time.Now()

		Tree := huffmantree.New()

		err := Tree.InitEncodingTable(conf.InputFile)
		if err != nil {
			fmt.Printf("Error initializing encoding table: %s\n", err.Error())
			return
		}

		_, err = Tree.Encode(conf.OutputPath + "_encoded")
		if err != nil {
			fmt.Printf("Error encoding file: %s\n", err.Error())
		}
		duration := time.Since(begin)

		if conf.ShouldDisplayStats {
			Tree.PrintStats(duration)
		}

		renderer.Render(Tree.Root)
		return
	}

	// if decode
	err := huffmantree.Decode(conf.InputFile, conf.OutputPath+"_decoded")
	if err != nil {
		fmt.Printf("Error decoding file: %s\n", err.Error())
	}
}
