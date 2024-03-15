package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	fileName := flag.String("f", "", "The filename to read")
	date := flag.String("d", "", "The date to be parsed")
	flag.Parse()
	// Check if a filename was provided with the -f flag
	if *fileName == "" {
		log.Fatal("Please provide a filename using the -f flag.")
		return
	}
	if *date == "" {
		log.Fatal("Please provide the date -d flag")
		return
	}
	fmt.Println("date", *date)
}
