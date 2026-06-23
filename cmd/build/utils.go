package main

import "os"

func cleanOutputDir() {
	os.RemoveAll("public")
	os.MkdirAll("public", 0755)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
