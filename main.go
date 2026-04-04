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

	// if should encode
	if !conf.IsDecode {
		// start recording execution time
		begin := time.Now()

		Tree := huffmantree.New()

		// create our leaf table and encoding table for compression
		err := Tree.InitEncodingTable(conf.InputFile)
		if err != nil {
			fmt.Printf("Error initializing encoding table: %s\n", err.Error())
			return
		}

		// encode the file from argv
		_, err = Tree.Encode(conf.OutputPath + "_encoded")
		if err != nil {
			fmt.Printf("Error encoding file: %s\n", err.Error())
			return
		}
		duration := time.Since(begin)

		// display stats if applicable
		if conf.ShouldDisplayStats {
			Tree.PrintStats(&duration)
			if conf.ShouldGenerateHTML {
				err = renderer.CreateHTMLView(Tree.Root, Tree.Stats, conf.OutputPath+".html")
				if err != nil {
					fmt.Printf("Error creating html view: %s\n", err.Error())
					return
				}
				fmt.Printf("HTML File exported to %s.html", conf.OutputPath)
				return
			}
		}
		return
	}

	// if should decode
	err := huffmantree.Decode(conf.InputFile, conf.OutputPath)
	if err != nil {
		fmt.Printf("Error decoding file: %s\n", err.Error())
		return
	}
}
