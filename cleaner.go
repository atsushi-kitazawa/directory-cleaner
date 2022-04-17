package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	load()

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

		modTime := f.ModTime().Format("20060102")
		if err := mkdir(fmt.Sprintf("%s/%s", toDir, modTime)); err != nil {
			log.Println(err)
			continue
		}

		from := fmt.Sprintf("%s/%s", fromDir, f.Name())
		to := fmt.Sprintf("%s/%s/%s", toDir, modTime, f.Name())
		if err := os.Rename(from, to); err != nil {
			log.Println(err)
		}

	}
}

func isExclude(path string) bool {
	for _, e := range exDir {
		if path == e {
			return true
		}
	}

	for _, e := range exExt {
		if strings.HasSuffix(path, e) {
			return true
		}
	}

	return false
}
