package main

import (
	"fmt"
	"goexcrs/htmllinkparser"
	"log"
	"os"
)

func main() {
	filenames := []string{"main/testdata/ex1.html", "main/testdata/ex2.html", "main/testdata/ex3.html", "main/testdata/ex4.html"}
	for _, fn := range filenames {
		fd, err := os.Open(fn)
		if err != nil {
			log.Fatalf("Fatal error while opening file\n%+v\n", err)
		}
		fmt.Println(fd.Name())
		htmllinkparser.ParseHTMLandExecute(fd)
		fd.Close()
	}
}
