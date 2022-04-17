package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	fromDir string
	toDir   string
	exDir   excludeDir
	exExt   excludeExt
)

type excludeDir []string
type excludeExt []string

func (ed *excludeDir) String() string {
	return fmt.Sprintf("%v", *ed)
}

func (ed *excludeDir) Set(v string) error {
	*ed = append(*ed, v)
	return nil
}

func (ee *excludeExt) String() string {
	return fmt.Sprintf("%v", *ee)
}

func (ee *excludeExt) Set(v string) error {
	*ee = append(*ee, v)
	return nil
}

func load() {
	if len(os.Args) > 1 {
		cmdline()
		return
	}

	config()
}

func cmdline() {
	flag.StringVar(&fromDir, "fromDir", "./desktop", "move from directory")
	flag.StringVar(&toDir, "toDir", "./desktop/move", "move to directory")
	flag.Var(&exDir, "exDir", "exclusive directory lists")
	flag.Var(&exExt, "exExt", "exclusive extension lists")
	flag.Parse()

	// if 1 == 1 {
	// 	fmt.Println(fromDir)
	// 	fmt.Println(toDir)
	// 	fmt.Println(exDir)
	// 	fmt.Println(exExt)
	// }
}

func config() {

}
