// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func createBackup(f *os.File) (string, error) {
	var err error

	_, err = f.Seek(0, 0)

	if err != nil {
		return "", err
	}

	name := f.Name()
	outputName := name[:strings.Index(name, ".tpl")] + "_backup" + ".tpl"

	outputFile, err := os.Create(outputName)

	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	reader := bufio.NewReader(f)
	writer := bufio.NewWriter(outputFile)

	var buf int

	for {
		line, err := reader.ReadString('\n')

		// fmt.Print(line)

		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		n, err := writer.WriteString(line)
		if err != nil {
			return "", err
		}

		buf += n

		if buf > 1024 {
			writer.Flush()
		}
	}

	writer.Flush()

	return outputName, nil
}
