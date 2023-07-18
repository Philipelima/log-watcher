package main

import (
	"fmt"
	"log"
	"os"
	"github.com/fsnotify/fsnotify"
)

func Watch(path string) {

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal("Failed on create new watcher")
	}

	defer watcher.Close()

	done := make(chan bool)

	go readFile(done, watcher, path)

	err = watcher.Add(path)

	if err != nil {
		log.Fatal("Add failed:", err)
	}

	<-done
}



func readFile(done chan bool, watcher *fsnotify.Watcher, path string) {

	defer close(done)

	for {
		select {
		case _, ok := <-watcher.Events:

			if !ok {
				return
			}

			file, err := os.ReadFile(path)

			if err != nil {
				panic(err)
			}

			fmt.Printf(" %s ", string(file))

		case err, ok := <-watcher.Errors:

			if !ok {
				return
			}

			log.Println("error:", err)
		}
	}
}
