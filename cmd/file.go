package main

import (
	"fmt"
	"log"
	"os"
)


func IsFile(file string) bool {
	return fileExists(file)
}

func FileInfo(file string) {

	info, err := os.Stat(file)

	if err != nil {
		log.Println(err)
		return
	}

	mode := info.Mode()

	fmt.Printf("\n File: %s \t \t Mode: %s \n \n ", file, mode)

}

func fileExists(file string) ( bool ){
	info, err := os.Stat(file)

	if os.IsNotExist(err) {	
		return false
	}

	return !info.IsDir()
}