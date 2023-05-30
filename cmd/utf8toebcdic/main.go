package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/indece-official/go-ebcdic"
)

// Variables set during build
var (
	ProjectName  string
	ProjectURL   string
	BuildVersion string
	BuildDate    string
)

var (
	flagInputFile  = flag.String("i", "-", "Input utf8 file (use '-' for stdin)")
	flagCodepage   = flag.String("c", "037", "EBCDIC-Codepage to use (supported: 037 / 273 / 500 / 1140 / 1141 / 1148)")
	flagOutputFile = flag.String("o", "-", "Output ebcdic file (use '-' for stdout)")
	flagVersion    = flag.Bool("version", false, "Print version and quit with exit code 0")
)

func printOwnVersion() {
	fmt.Printf("%s %s (Build %s)\n", ProjectName, BuildVersion, BuildDate)
	fmt.Printf("\n")
	fmt.Printf("%s\n", ProjectURL)
	fmt.Printf("\n")
	fmt.Printf("Copyright 2023 by indece UG (haftungsbeschränkt)\n")
}

func main() {
	var err error

	flag.Parse()

	if *flagVersion {
		printOwnVersion()
		os.Exit(0)
	}

	if *flagInputFile == "" {
		fmt.Fprintf(os.Stderr, "Missing input file\n")
		os.Exit(1)
	}

	if *flagOutputFile == "" {
		fmt.Fprintf(os.Stderr, "Missing output file\n")
		os.Exit(1)
	}

	codePage := 0

	switch strings.ToLower(*flagCodepage) {
	case "37", "037", "ebcdic037":
		codePage = ebcdic.EBCDIC037
	case "273", "ebcdic273":
		codePage = ebcdic.EBCDIC273
	case "500", "ebcdic500":
		codePage = ebcdic.EBCDIC500
	case "1140", "ebcdic1140":
		codePage = ebcdic.EBCDIC1140
	case "1141", "ebcdic1141":
		codePage = ebcdic.EBCDIC1141
	case "1148", "ebcdic1148":
		codePage = ebcdic.EBCDIC1148
	default:
		fmt.Fprintf(os.Stderr, "Unsupported codepage %s\n", *flagCodepage)
		os.Exit(1)
	}

	var reader io.ReadCloser

	if *flagInputFile == "-" {
		reader = os.Stdin
	} else {
		reader, err = os.Open(*flagInputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %s\n", err)
		}
	}

	defer reader.Close()

	var writer io.WriteCloser

	if *flagOutputFile == "-" {
		writer = os.Stdout
	} else {
		writer, err = os.Create(*flagOutputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening output file: %s\n", err)
		}
	}

	defer writer.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		inputData := scanner.Text()

		outputData, err := ebcdic.Encode(inputData, codePage)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding data: %s\n", err)
		}

		_, err = writer.Write(outputData)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to output file: %s\n", err)
		}
	}
}
