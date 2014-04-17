package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	bs, err := ioutil.ReadFile("src/foo.txt")
	if err != nil {
		return }
	str := string(bs)
	fmt.Println(str)

	filepath.Walk(".", func(path string, _ os.FileInfo, _ error) error {
			fmt.Println(path)
			return nil})
}
