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
	var insideLiteralTag bool

	parse := func(line string) string {
		if line == "" {
			return line
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

		return line
	}

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

		if !insideLiteralTag {
			insideLiteralTag = startOfLiteralTag(line)
		}

		if insideLiteralTag {
			insideLiteralTag = !endOfLiteralTag(line)
			writer.WriteString(line)
			continue
		}

		if isRegExp(line) {
			slice, _ := parseRegExp(line)

			line = parse(slice[0]) + slice[1] + parse(slice[2]) + "\n"
		} else {
			line = parse(line)
		}

		if insideScriptTag {
			insideScriptTag = !endOfScriptTag(line)
		}

		writer.WriteString(line)
	}

	writer.Flush()

	return nil
}

// ----------------------- BRACKETS
func parseDelims(inputFile io.Reader, outputFile io.Writer) error {
	reader := bufio.NewReaderSize(inputFile, 1024)
	writer := bufio.NewWriterSize(outputFile, 1024)

	var insideScriptTag bool
	var insideLiteralTag bool

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

		if !insideLiteralTag {
			insideLiteralTag = startOfLiteralTag(line)
		}

		if insideLiteralTag {
			insideLiteralTag = !endOfLiteralTag(line)
			writer.WriteString(line)
			continue
		}

		line, _ = parseLDelim(line)
		line, _ = parseRDelim(line)

		if insideScriptTag {
			insideScriptTag = !endOfScriptTag(line)
		}

		writer.WriteString(line)
	}

	writer.Flush()

	return nil
}
