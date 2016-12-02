// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bufio"
	"io"
)

// ----------------------- DELIMS
func parseDelims(inputFile io.Reader, outputFile io.Writer) error {
	reader := bufio.NewReaderSize(inputFile, 1024)
	writer := bufio.NewWriterSize(outputFile, 1024)

	var insideScriptTag bool
	var insideLiteralTag bool
	var insidePHPTag bool
	var insideMultilineComment bool
	var cm []string
	var mlm []string

	for {
		var isCommentLine bool
		var firstML bool
		var comment string
		var leftComment string
		var rightComment string

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

		if !insideMultilineComment {
			mlm, insideMultilineComment = parseMultilineCommentStart(line, false)

			if !insideMultilineComment {
				mlm, insideMultilineComment = parseMultilineCommentStart(line, true)
			}

			if insideMultilineComment {
				firstML = true
				line = mlm[0]
				rightComment = mlm[1] + "\n"
			}
		}

		if insideMultilineComment && !firstML {
			var endMultiline bool

			mlm, endMultiline = parseMultilineCommentEnd(line, false)

			if !endMultiline {
				mlm, endMultiline = parseMultilineCommentEnd(line, true)
			}

			if endMultiline {
				insideMultilineComment = false
				leftComment = mlm[0]
				line = mlm[1] + "\n"
			} else {
				writer.WriteString(line)
				continue
			}
		}

		cm, isCommentLine = parseCommentLine(line)

		if isCommentLine {
			line = cm[0]
			comment = cm[1] + "\n"
		}

		if !insidePHPTag {
			insidePHPTag = startOfPHPTag(line)
		}

		if insidePHPTag {
			insidePHPTag = !endOfPHPTag(line)
			writer.WriteString(leftComment + line + comment + rightComment)
			continue
		}

		if !insideLiteralTag {
			insideLiteralTag = startOfLiteralTag(line)
		}

		if insideLiteralTag {
			insideLiteralTag = !endOfLiteralTag(line)
			writer.WriteString(leftComment + line + comment + rightComment)
			continue
		}

		line, _, split := parseLDelim(line)
		if split {
			line = line + "\n"
		}

		line, _, split = parseRDelim(line)
		if split {
			line = line + "\n"
		}

		if insideScriptTag {
			insideScriptTag = !endOfScriptTag(line)
		}

		writer.WriteString(leftComment + line + comment + rightComment)
	}

	writer.Flush()

	return nil
}
