// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
)

var inputArg = flag.String("i", "", "Input file path")
var outputArg = flag.String("o", "", "Output file path (if not provied will overwrite input file)")
var bracketArg = flag.Bool("b", false, "Parse brackets into {delim}")
var delimArg = flag.Bool("d", false, "Parse {delim} into brackets")
var rmArg = flag.Bool("rm", false, "Remove backup file after parse")
var owArg = flag.Bool("ow", false, "Overwrite backup file if already exist")

func main() {
	flag.Parse()

	backupSuffix := "_backup"
	inputPath := *inputArg
	outputPath := *outputArg
	removeBackup := *rmArg
	overWrite := *owArg
	bracket := *bracketArg
	delim := *delimArg

	if !bracket && !delim {
		fmt.Println("Must choose an type of action delim or bracket parse")
		return
	} else if bracket && delim {
		fmt.Println("Must choose between delim or bracket parse, not both")
		return
	}

	if outputPath == "" {
		outputPath = inputPath
	}

	backup, err := createBackup(inputPath, backupSuffix, overWrite)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error ocurred during backup creation: %s", err))
		return
	}

	inputFile, err := os.Open(backup)
	defer inputFile.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error ocurred while trying to read backup file: %s", err))
		return
	}

	outputFile, err := os.Create(outputPath)
	defer inputFile.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error ocurred creating output file: %s", err))
		return
	}

	if bracket {
		err = parseBrackets(inputFile, outputFile)
	} else if delim {
		err = parseDelims(inputFile, outputFile)
	}

	if err != nil {
		t := "bracket"

		if delim {
			t = "delim"
		}

		fmt.Println(fmt.Sprintf("Error during %s parse operation: %s", t, err))
		return
	}

	if removeBackup {
		err = os.Remove(inputFile.Name())

		if err != nil {
			fmt.Println(fmt.Sprintf("Error ocurred trying to remove backup file %s", err))
		}
	}

}
