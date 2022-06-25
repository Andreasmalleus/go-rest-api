package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func GetFileContents(filename string) string {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
