// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

// ------------ INLINE OBJECT

var inlineObject = []string{
	`var car = {type:"Fiat", model:"500", color:"white"};`,
	`var person = {firstName: '{$firstName}', lastName: "Doe", age: {json_decode($age)}, eyeColor: {$var}};`,
	`const myObject = {hello: "world", myObject:{one: 1, two: [2, 2]}}`,
	`inline_call({hello: "world", myObject:{one: 1, two: [2, 2]}})`,
}

var expInlineObject = []string{
	`var car = {ldelim}type:"Fiat", model:"500", color:"white"{rdelim};`,
	`var person = {ldelim}firstName: '{$firstName}', lastName: "Doe", age: {json_decode($age)}, eyeColor: {$var}{rdelim};`,
	`const myObject = {ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim}`,
	`inline_call({ldelim}hello: "world", myObject:{ldelim}one: 1, two: [2, 2]{rdelim}{rdelim})`,
}

var nonInlineObject = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`funcion () {`,
	`}`,
	`call({`,
	`hello: "world"`,
	`}, {`,
	`world: "hello"`,
	`})`,
	`let array = [{`,
	`hello: "world",`,
	`myObject:{`,
	`one: 1,`,
	`two: [2, 2]`,
	`}`,
	`}]`,
	`const strangeObject = {maybe: {it: {wont: {work: "?"`,
	`}, maybe: ""}, not: "did"}, work: "entirely"}`,
	`</script>`,
	`</body>`,
}

func TestParseInlineObjectMatch(t *testing.T) {
	for i, line := range inlineObject {
		nl, matched := parseInlineObject(line)

		if !matched {
			t.Fatalf("Should parse")
		}

		if nl != expInlineObject[i] {
			t.Fatalf("Expected inline object parsed: %s; got: %s", expInlineObject[i], nl)
		}
	}
}

func TestParseInlineObjectNoMatch(t *testing.T) {
	for _, line := range nonInlineObject {
		_, matched := parseInlineObject(line)

		if matched {
			t.Fatalf("Should not match")
		}
	}
}
