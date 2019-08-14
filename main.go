package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		err := errors.New("[args] pomfile configfile")
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	placementAfter := "</developers>"

	pomfile := os.Args[1]
	txtPomfile, err := ioutil.ReadFile(pomfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configfile := os.Args[2]
	txtConfigfile, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dst := bytes.Replace(txtPomfile, []byte(placementAfter), []byte(placementAfter + string(txtConfigfile)), -1)
	if err := ioutil.WriteFile(pomfile, dst, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("⚙ POM configuration updated ✨")
}
