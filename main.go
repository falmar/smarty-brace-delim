// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"errors"
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

	args := map[string]interface{}{
		"backupSuffix": "_backup",
		"inputPath":    *inputArg,
		"outputPath":   *outputArg,
		"removeBackup": *rmArg,
		"overWrite":    *owArg,
		"bracket":      *bracketArg,
		"delim":        *delimArg,
	}

	code, err := altMain(args)

	if err != nil {
		fmt.Println(err)
	}

	os.Exit(code)
}

func altMain(args map[string]interface{}) (int, error) {
	backupSuffix := args["backupSuffix"].(string)
	inputPath := args["inputPath"].(string)
	outputPath := args["outputPath"].(string)
	removeBackup := args["removeBackup"].(bool)
	overWrite := args["overWrite"].(bool)
	bracket := args["bracket"].(bool)
	delim := args["delim"].(bool)

	if !bracket && !delim {
		return 1, errors.New("Must choose an type of action delim or bracket parse")
	} else if bracket && delim {

		return 1, errors.New("Must choose between delim or bracket parse, not both")
	}

	if outputPath == "" {
		outputPath = inputPath
	}

	backup, err := createBackup(inputPath, backupSuffix, overWrite)
	if err != nil {
		return 2, fmt.Errorf("Error ocurred during backup creation: %s", err)
	}

	inputFile, err := os.Open(backup)
	defer inputFile.Close()
	if err != nil {
		return 3, fmt.Errorf("Error ocurred while trying to read backup file: %s", err)
	}

	outputFile, err := os.Create(outputPath)
	defer inputFile.Close()
	if err != nil {
		return 4, fmt.Errorf("Error ocurred creating output file: %s", err)
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

		return 5, fmt.Errorf("Error during %s parse operation: %s", t, err)
	}

	if removeBackup {
		err = os.Remove(inputFile.Name())

		if err != nil {
			return 6, fmt.Errorf("Error ocurred trying to remove backup file %s", err)
		}
	}

	return 0, nil
}
