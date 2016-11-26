// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bufio"
	"io"
)

// ----------------------- BRACKETS
func parseBrackets(inputFile io.Reader, outputFile io.Writer) error {
	reader := bufio.NewReaderSize(inputFile, 1024)
	writer := bufio.NewWriterSize(outputFile, 1024)

	var insideScriptTag bool

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if !insideScriptTag {
			insideScriptTag = startOfScriptTag(line)
		}

		if !insideScriptTag {
			writer.WriteString(line)
			continue
		}

		var matched bool

		line, matched = parseInlineObject(line)

		if matched {
			line = line + "\n"
		}

		line, matched = parseLeftBracket(line)

		if matched {
			line = line + "\n"
		}

		line, matched = parseRightBracket(line)

		if matched {
			line = line + "\n"
		}

		if insideScriptTag {
			insideScriptTag = !endOfScriptTag(line)
		}

		writer.WriteString(line)
	}

	writer.Flush()

	return nil
}
