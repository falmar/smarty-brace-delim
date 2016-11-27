// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

// ------------ LEFT BRACKET

var leftBracket = []string{
	"function () {",
	"call({",
	"myObject: { // random comment",
	"let array = [{",
	"{",
	"calling({my: {",
	"const single = {}",
	"{}",
}

var expLeftBracket = []string{
	"function () {ldelim}",
	"call({ldelim}",
	"myObject: {ldelim} // random comment",
	"let array = [{ldelim}",
	"{ldelim}",
	"calling({ldelim}my: {ldelim}",
	"const single = {ldelim}}",
	"{ldelim}}",
}

var nonLeftBracket = []string{
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

func TestParseLeftBracketMatch(t *testing.T) {
	for i, line := range leftBracket {
		nl, matched := parseLeftBracket(line)

		if !matched {
			t.Fatalf("Should match %s", line)
		}

		if nl != expLeftBracket[i] {
			t.Fatalf("Expected left bracket parsed: %s; got: %s", expLeftBracket[i], nl)
		}
	}
}

func TestParseLeftBracketNoMatch(t *testing.T) {
	for _, line := range nonLeftBracket {
		nl, matched := parseLeftBracket(line)

		if matched && nl != line {
			t.Fatalf("Should not match %s; %s", line, nl)
		}
	}
}

// ------------ RIGHT BRACKET

var rightBracket = []string{
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

var expRightBracket = []string{
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

var nonRightBracket = []string{
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

func TestParseRightBracket(t *testing.T) {
	for i, line := range rightBracket {
		nl, matched := parseRightBracket(line)

		if !matched {
			t.Fatalf("Should match: %s", line)
		}

		if nl != expRightBracket[i] {
			t.Fatalf("Expected right bracket parsed: %s; got: %s", expRightBracket[i], nl)
		}
	}
}

func TestParseRightBracketNoMatch(t *testing.T) {
	for _, line := range nonRightBracket {
		nl, matched := parseRightBracket(line)

		if matched && nl != line {
			t.Fatalf("Should not match %s; %s", line, nl)
		}
	}
}
