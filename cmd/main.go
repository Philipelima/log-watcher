package main

import (
	"fmt"
	"log"
	"os"
)



func main() {

	fmt.Println("\n Wellcome to File-Watcher! ")

	args := os.Args
	file := args[1]

	if !IsFile(file) {
		log.Fatal("Sorry arg must be a file path")
	} 

	FileInfo(file)
	Watch(file)

	
}