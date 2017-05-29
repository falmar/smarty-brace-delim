// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

// ------------ LEFT BRACE

var leftBrace = []string{
	"function () {",
	"call({",
	"myObject: { // random comment",
	"let array = [{",
	"{",
	"calling({my: {",
	"const single = {}",
	"{}",
}

var expLeftBrace = []string{
	"function () {ldelim}",
	"call({ldelim}",
	"myObject: {ldelim} // random comment",
	"let array = [{ldelim}",
	"{ldelim}",
	"calling({ldelim}my: {ldelim}",
	"const single = {ldelim}}",
	"{ldelim}}",
}

var nonLeftBrace = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`const myObject = {hello: "world", myObject:{one: 1, two: [2, 2]}}`,
	`}`,
	`hello: "world"`,
	`world: "hello"`,
	`})`,
	`hello: "world",`,
	`one: 1,`,
	`two: [2, 2]`,
	`}`,
	`}]`,
	`</script>`,
	`</body>`,
}

func TestParseLeftBraceMatch(t *testing.T) {
	for i, line := range leftBrace {
		nl, matched := parseLeftBrace(line)

		if !matched {
			t.Fatalf("Should match %s", line)
		}

		if nl != expLeftBrace[i] {
			t.Fatalf("Expected left brace parsed: %s; got: %s", expLeftBrace[i], nl)
		}
	}
}

func TestParseLeftBraceNoMatch(t *testing.T) {
	for _, line := range nonLeftBrace {
		nl, matched := parseLeftBrace(line)

		if matched && nl != line {
			t.Fatalf("Should not match %s; %s", line, nl)
		}
	}
}

// ------------ RIGHT BRACE

var rightBrace = []string{
	`everthing }`,
	`}, { should`,
	`be }) good`,
	`or }] not?`,
	`... maybe }] }?`,
	`}, {ldelim}`,
	`}, maybe: ""}, did: "not"}, work: {"entirely"}`,
	`}, maybe: ""}, did: "not"}, work: {"entir}ely"}`,
	"const single = {}",
	"{}",
}

var expRightBrace = []string{
	`everthing {rdelim}`,
	`{rdelim}, { should`,
	`be {rdelim}) good`,
	`or {rdelim}] not?`,
	`... maybe {rdelim}] {rdelim}?`,
	`{rdelim}, {ldelim}`,
	`{rdelim}, maybe: ""{rdelim}, did: "not"{rdelim}, work: {"entirely"}`,
	`{rdelim}, maybe: ""{rdelim}, did: "not"{rdelim}, work: {"entir}ely"{rdelim}`,
	"const single = {{rdelim}",
	"{{rdelim}",
}

var nonRightBrace = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`funcion () {`,
	`const myObject = {hello: "world", myObject:{one: 1, two: [2, 2]}`,
	`call({`,
	`hello: "world"`,
	`, {`,
	`world: "hello"`,
	`)`,
	`let array = [{`,
	`hello: "world",`,
	`myObject:{`,
	`one: 1,`,
	`two: [2, 2] `,
	`]`,
	`inline_call({hello: "world", myObject:{one: 1, two: [2, 2]})`,
	`</script>`,
	`</body>`,
}

func TestParseRightBrace(t *testing.T) {
	for i, line := range rightBrace {
		nl, matched := parseRightBrace(line)

		if !matched {
			t.Fatalf("Should match: %s", line)
		}

		if nl != expRightBrace[i] {
			t.Fatalf("Expected right brace parsed: %s; got: %s", expRightBrace[i], nl)
		}
	}
}

func TestParseRightBraceNoMatch(t *testing.T) {
	for _, line := range nonRightBrace {
		nl, matched := parseRightBrace(line)

		if matched && nl != line {
			t.Fatalf("Should not match %s; %s", line, nl)
		}
	}
}
