package main

import (
	"golang.org/x/mod/modfile"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		source := os.Args[1]
		target := os.Args[2]
		sourceFile, err := os.ReadFile(source)
		if err != nil {
			panic(err)
		}
		sourceMod, err := modfile.Parse(source, sourceFile, nil)
		if err != nil {
			panic(err)
		}
		targetFile, err := os.ReadFile(target)
		if err != nil {
			panic(err)
		}
		targetMod, err := modfile.Parse(target, targetFile, nil)
		if err != nil {
			panic(err)
		}
		for _, replace := range sourceMod.Replace {
			_ = targetMod.AddReplace(replace.Old.Path, replace.Old.Version, replace.New.Path, replace.New.Version) // don't have any error
		}
		targetFile, err = targetMod.Format()
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(target, targetFile, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Usage: update-go-mod-replace <source file> <target file>")
	}
}
