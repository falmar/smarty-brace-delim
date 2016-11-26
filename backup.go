// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func createBackup(path, suffix string, overwrite bool) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

	name := f.Name()
	outputName := name[:strings.Index(name, ".tpl")] + suffix + ".tpl"

	_, err = os.Stat(outputName)
	if !os.IsNotExist(err) && !overwrite {
		return "", errors.New("Backup file already exist")
	}

	outputFile, err := os.Create(outputName)

	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	reader := bufio.NewReaderSize(f, 1024)
	writer := bufio.NewWriterSize(outputFile, 1024)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		_, err = writer.WriteString(line)
		if err != nil {
			return "", err
		}
	}

	writer.Flush()

	return outputName, nil
}
