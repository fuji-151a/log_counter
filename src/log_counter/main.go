package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	dir := os.Args[1]
	time := strings.Trim(dir, "/")
	max := logCount(dir, dir)
	fmt.Println(time + "\t" + strconv.Itoa(max))
}

func logCount(rootPath, searchPath string) int {
	fis, err := ioutil.ReadDir(searchPath)

	if err != nil {
		panic(err)
	}

	var max = 0
	for _, fi := range fis {
		fullPath := filepath.Join(searchPath, fi.Name())

		if fi.IsDir() {
			logCount(rootPath, fullPath)
		} else {
			rel, err := filepath.Rel(rootPath, fullPath)

			if err != nil {
				panic(err)
			}
			max += readFiles(rootPath + rel)
		}
	}
	return max
}

func readFiles(fileName string) int {
	var fp *os.File
	var err error

	fp, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	var cnt = 0
	for scanner.Scan() {
		cnt++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return cnt
}
