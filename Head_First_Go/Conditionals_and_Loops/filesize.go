package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileInfo, err := os.Stat("condional.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}

/*
A Go program that prints the size of a file is on the fridge. It calls the os.Stat function, which returns an os.FileInfo value, and possibly an error value. Then it calls the Size method on the FileInfo value to get the file size.
*/
