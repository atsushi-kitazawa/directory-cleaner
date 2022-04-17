package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	fromDir string = "./desktop"
	toDir   string = "./desktop/move"

	//only available under the top directory
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

	traverse(fromDir)
}

func mkdir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func traverse(baseDir string) {
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatalln(err)
	}

	var tmp string
	for _, f := range files {
		//only the top directory is valid for exclusion listings
		if baseDir == fromDir && isExclude(f.Name()) {
			continue
		}

		if f.IsDir() {
			dir := baseDir + "/" + f.Name()
			tmp = toDir
			toDir = toDir + "/" + f.Name()
			traverse(dir)
		} else {
			file := baseDir + "/" + f.Name()
			move(f, file, toDir)
		}
	}
	toDir = tmp
}

func move(f os.FileInfo, file string, toDir string) {
	if err := mkdir(toDir); err != nil {
		log.Println(err)
		return
	}

	if err := copy(f, file, toDir); err != nil {
		panic(err)
	}

	delete(file)
}

func copy(f os.FileInfo, file string, toDir string) error {
	from, err := os.Open(file)
	if err != nil {
		log.Println("os.Open " + err.Error())
		return err
	}
	defer from.Close()

	to, err := os.Create(toDir + "/" + f.Name())
	if err != nil {
		log.Println("os.Create " + err.Error())
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Println("io.Copy " + err.Error())
		return err
	}

	return nil
}

func delete(file string) {

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
