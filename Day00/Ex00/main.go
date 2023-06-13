package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	meanVar, medianVar, modeVar, sdVar bool
)

func init() {

	if len(os.Args) == 1 {
		flag.BoolVar(&meanVar, "Mean", true, "-Mean=true/false")
		flag.BoolVar(&medianVar, "Median", true, "-Median=true/false")
		flag.BoolVar(&modeVar, "Mode", true, "-Mode=true/false")
		flag.BoolVar(&sdVar, "SD", true, "-SD=true/false")
	} else {
		flag.BoolVar(&meanVar, "Mean", false, "-Mean=true/false")
		flag.BoolVar(&medianVar, "Median", false, "-Median=true/false")
		flag.BoolVar(&modeVar, "Mode", false, "-Mode=true/false")
		flag.BoolVar(&sdVar, "SD", false, "-SD=true/false")
	}
}

func main() {

	flag.Parse()

	in, err := ReadInput(bufio.NewScanner(os.Stdin))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Anscombe(in))
}
