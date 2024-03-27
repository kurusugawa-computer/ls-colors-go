package main

import (
	"fmt"
	"os"

	lscolors "github.com/kurusugawa-computer/ls-colors-go"
)

func printStringPointer(s *string) string {
	if s == nil {
		return "undefined"
	} else {
		return *s
	}
}

func main() {
	lsColors := os.Getenv("LS_COLORS")
	result, err := lscolors.ParseLS_COLORS(lsColors, false)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	fmt.Printf("no: %s\n", printStringPointer(result.Normal))
	fmt.Printf("fi: %s\n", printStringPointer(result.FileDefault))
	fmt.Printf("di: %s\n", printStringPointer(result.Directory))
	fmt.Printf("ln: %s\n", printStringPointer(result.Symlink))
	fmt.Printf("pi: %s\n", printStringPointer(result.Pipe))
	fmt.Printf("so: %s\n", printStringPointer(result.Socket))
	fmt.Printf("bd: %s\n", printStringPointer(result.BlockDevice))
	fmt.Printf("cd: %s\n", printStringPointer(result.CharDevice))
	fmt.Printf("mi: %s\n", printStringPointer(result.MissingFile))
	fmt.Printf("or: %s\n", printStringPointer(result.OrphanedSymlink))
	fmt.Printf("ex: %s\n", printStringPointer(result.Executable))
	fmt.Printf("do: %s\n", printStringPointer(result.Door))
	fmt.Printf("su: %s\n", printStringPointer(result.SetUID))
	fmt.Printf("sg: %s\n", printStringPointer(result.SetGID))
	fmt.Printf("st: %s\n", printStringPointer(result.Sticky))
	fmt.Printf("ow: %s\n", printStringPointer(result.OtherWritable))
	fmt.Printf("tw: %s\n", printStringPointer(result.OtherWritableSticky))
	fmt.Printf("ca: %s\n", printStringPointer(result.Cap))
	fmt.Printf("mh: %s\n", printStringPointer(result.MultiHardLink))

	for _, ext := range result.Extensions {
		if ext.ExactMatch {
			fmt.Printf("*%s: %s (case sensitive)\n", ext.Extension, ext.Sequence)
		} else {
			fmt.Printf("*%s: %s\n", ext.Extension, ext.Sequence)
		}
	}
}
