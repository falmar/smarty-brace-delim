// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

// ------------ LDELIM

var lDelims = []string{
	`funcion () {ldelim}`,
	`const myObject = {ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim}`,
	`call({ldelim}`,
	`{rdelim}, {ldelim}`,
	`let array = [{ldelim}`,
	`myObject:{ldelim}`,
	`const strangeObject = {ldelim}maybe: {ldelim}it: {ldelim}wont: "{ldelim}work: ?"`,
	`inline_call({ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim})`,
	`object = {ldelim}left: ["{lrdelim}", "{rdelim}"], right: {ldelim}"{rdelim}", "{ldelim}"{rdelim}{rdelim}`,
}

var expLDelims = []string{
	`funcion () {`,
	`const myObject = {hello: "world", myObject:{one: 1, two: [2, 2]{rdelim}{rdelim}`,
	`call({`,
	`{rdelim}, {`,
	`let array = [{`,
	`myObject:{`,
	`const strangeObject = {maybe: {it: {wont: "{ldelim}work: ?"`,
	`inline_call({hello: "world", myObject:{one: 1, two: [2, 2]{rdelim}{rdelim})`,
	`object = {left: ["{lrdelim}", "{rdelim}"], right: {"{rdelim}", "{ldelim}"{rdelim}{rdelim}`,
}

var nonLDelims = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`let some = 0`,
	`{rdelim}`,
	`hello: "world"`,
	`world: "hello"`,
	`{rdelim})`,
	`hello: "world",`,
	`one: 1,`,
	`two: [2, 2]`,
	`{rdelim}`,
	`{rdelim}]`,
	`{rdelim}, maybe: ""{rdelim}, did: "not"{rdelim}, work: "entirely"{rdelim}`,
	`</script>`,
	`</body>`,
	`{rdelim}, "{ldelim}"`,
	`{rdelim}, '{ldelim}'`,
	"{rdelim}, `{ldelim}`",
	`console.log('{rdelim}')`,
	`console.log("{ldelim}")`,
	"object.call('{rdelim}', \"{ldelim}\", `{ldelim} & {rdelim}`)",
}

func TestParseLDelimMatch(t *testing.T) {
	for i, line := range lDelims {
		nl, matched, _ := parseLDelim(line)

		if !matched {
			t.Fatalf("Should parse: %s", line)
		}

		if nl != expLDelims[i] {
			t.Fatalf("Expected ldelim parsed: %s; got: %s", expLDelims[i], nl)
		}
	}
}

func TestParseLDelimNotMatch(t *testing.T) {
	for _, line := range nonLDelims {
		nl, matched, _ := parseLDelim(line)

		if matched || nl != line {
			t.Fatalf("Should not match or perform change on string: %s", line)
		}
	}
}

// ------------ LDELIM

var rDelims = []string{
	`const myObject = {ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim}`,
	`{rdelim}`,
	`{rdelim}, {ldelim}`,
	`{rdelim})`,
	`{rdelim}`,
	`{rdelim}]`,
	`{rdelim}, maybe: ""{rdelim}, did: "{rdelim}not"{rdelim}, work: "entirely"{rdelim}`,
	`inline_call({ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim})`,
	`object = {ldelim}left: ["{lrdelim}", "{rdelim}"], right: {ldelim}"{rdelim}", "{ldelim}"{rdelim}{rdelim}`,
}

var expRDelims = []string{
	`const myObject = {ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]}}`,
	`}`,
	`}, {ldelim}`,
	`})`,
	`}`,
	`}]`,
	`}, maybe: ""}, did: "{rdelim}not"}, work: "entirely"}`,
	`inline_call({ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]}})`,
	`object = {ldelim}left: ["{lrdelim}", "{rdelim}"], right: {ldelim}"{rdelim}", "{ldelim}"}}`,
}

var nonRDelims = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`funcion () {ldelim}`,
	`let some = 0`,
	`call({ldelim}`,
	`hello: "world"`,
	`world: "hello"`,
	`let array = [{ldelim}`,
	`hello: "world",`,
	`myObject:{ldelim}`,
	`one: 1,`,
	`two: [2, 2]`,
	`const strangeObject = {ldelim}maybe: {ldelim}it: {ldelim}wont: {ldelim}work: "?"`,
	`</script>`,
	`</body>`,
	`"</body>{rdelim}const strangeObject", {ldelim}`,
	`'asdasd{rdelim}, '{ldelim}`,
	"`asdas{rdelim}asdas`, {ldelim}",
	`console.log('{rdelim}')`,
	`console.log("{ldelim}")`,
	"object.call('{rdelim}', \"{ldelim}\", `{ldelim} & {rdelim}`)",
}

func TestParseRDelimMatch(t *testing.T) {
	for i, line := range rDelims {
		nl, matched, _ := parseRDelim(line)

		if !matched {
			t.Fatalf("Should parse: %s", line)
		}

		if nl != expRDelims[i] {
			t.Fatalf("Expected ldelim parsed: %s; got: %s", expRDelims[i], nl)
		}
	}
}

func TestParseRDelimNotMatch(t *testing.T) {
	for _, line := range nonRDelims {
		nl, matched, _ := parseRDelim(line)

		if matched || nl != line {
			t.Fatalf("Should not match or perform change on string: %s; %s", line, nl)
		}
	}
}
