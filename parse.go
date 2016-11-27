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

	parse := func(line string, lf bool) string {
		if line == "" {
			return line
		}

		var anyMatched bool
		var matched bool

		line, matched = parseInlineObject(line)
		anyMatched = anyMatched || matched

		line, matched = parseLeftBracket(line)
		anyMatched = anyMatched || matched

		line, matched = parseRightBracket(line)
		anyMatched = anyMatched || matched

		if anyMatched && lf {
			line += "\n"
		}

		return line
	}

	for {
		var isCommentLine bool
		var comment string

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

		cm, isCommentLine := parseCommentLine(line)

		if isCommentLine {
			line = cm[0]
			comment = cm[1] + "\n"
		}

		if !insideLiteralTag {
			insideLiteralTag = startOfLiteralTag(line)
		}

		if insideLiteralTag {
			insideLiteralTag = !endOfLiteralTag(line)
			writer.WriteString(line + comment)
			continue
		}

		if isRegExp(line) {
			slice, _ := parseRegExp(line)

			line = parse(slice[0], false) + slice[1] + parse(slice[2], false)

			if !isCommentLine {
				line += "\n"
			}
		} else {
			line = parse(line, !isCommentLine)
		}

		if insideScriptTag {
			insideScriptTag = !endOfScriptTag(line)
		}

		writer.WriteString(line + comment)
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
		var isCommentLine bool
		var comment string

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

		cm, isCommentLine := parseCommentLine(line)

		if isCommentLine {
			line = cm[0]
			comment = cm[1] + "\n"
		}

		if !insideLiteralTag {
			insideLiteralTag = startOfLiteralTag(line)
		}

		if insideLiteralTag {
			insideLiteralTag = !endOfLiteralTag(line)
			writer.WriteString(line + comment)
			continue
		}

		line, _ = parseLDelim(line)
		line, _ = parseRDelim(line)

		if insideScriptTag {
			insideScriptTag = !endOfScriptTag(line)
		}

		writer.WriteString(line + comment)
	}

	writer.Flush()

	return nil
}
