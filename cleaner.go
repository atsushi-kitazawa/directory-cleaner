package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const fromDir string = "./desktop"
const toDir string = "./desktop/move"

var (
	excludeDir []string = []string{"move", "exclude1", "exclude2"}
	excludeExt []string = []string{"ext1", "ext2"}
)

func main() {
	doMain()
}

func doMain() {
	if err := mkdir(toDir); err != nil {
		log.Fatalln(err)
	}

	move()
}

func mkdir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func move() {
	files, err := ioutil.ReadDir(fromDir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		if isExclude(f.Name()) {
			continue
		}

		from := fromDir + "/" + f.Name()
		to := toDir + "/" + f.Name()
		if err := os.Rename(from, to); err != nil {
			log.Println(err)
		}

	}
}

func isExclude(path string) bool {
	for _, e := range excludeDir {
		if path == e {
			return true
		}
	}

	for _, e := range excludeExt {
		if strings.HasSuffix(path, e) {
			return true
		}
	}

	return false
}
