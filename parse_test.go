// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestParseBracket(t *testing.T) {
	input := "files/simple_bracket.tpl"
	output := "files/simple_bracket_parsed.tpl"
	exp := "files/simple_delim.tpl"

	inputFile, err := os.Open(input)
	defer inputFile.Close()
	if err != nil {
		t.Fatalf("Error opening input file: %s", err)
	}

	outPutFile, err := os.Create(output)
	defer outPutFile.Close()
	if err != nil {
		t.Fatalf("Error opening output file: %s", err)
	}

	err = parseBrackets(inputFile, outPutFile)

	if err != nil {
		t.Fatalf("Error during parse: %s", err)
	}

	expFile, err := os.Open(exp)
	defer expFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	outStat, err := outPutFile.Stat()
	if err != nil {
		t.Fatal(err)
	}

	expStat, err := expFile.Stat()
	if err != nil {
		t.Fatal(err)
	}

	if outStat.Size() != expStat.Size() {
		t.Fatalf("Expected filesize: %d; got: %d", expStat.Size(), outStat.Size())
	}

	outPutFile.Seek(0, 0)

	outReader := bufio.NewReaderSize(outPutFile, 1024)
	expReader := bufio.NewReaderSize(expFile, 1024)

	match := true

	for {
		outLine, err := outReader.ReadString('\n')

		if err != nil && err != io.EOF {
			t.Fatal(err)
		}

		expLine, err := expReader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}

		if outLine != expLine {
			t.Logf("Lines does not match \n%s\n%s", outLine, expLine)
			match = false
			break
		}
	}

	if !match {
		t.Fatal("Expected file and output file did not match")
	}
}
