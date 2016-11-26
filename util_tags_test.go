// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

// ------------ SCRIPT TAGS
var openScriptTags = []string{
	`<script src="js/pace.min.js">`,
	`<script src="app.min.js" another="weird-tag">`,
	`  yo-!    <script no-tag>       // random comment`,
}

var closeScriptTags = []string{
	`anything </script> //  comment`,
	`</script>`,
	`  yo-!    </script>       // random comment`,
}

var nonScriptTags = []string{
	`<!doctype html>`,
	`<html class="no-js" lang="en">`,
	`<head>`,
	`<meta charset="utf-8" />`,
	`<meta http-equiv="x-ua-compatible" content="ie=edge">`,
	`<meta name="viewport" content="width=device-width, initial-scale=1.0" />`,
	`<title>React-App</title>`,
	`<link rel="stylesheet" href="/css/pace.css">`,
	`<link rel="stylesheet" href="/css/app.min.css">`,
	`</head>`,
	`<body>`,
	`<div id="app"></div>`,
	`<script src="js/pace.min.js"></script>`,
	`<script src="app.min.js"></script>`,
	`<script src="app.min.js"> function() {}</script>`,
	`</body>`,
	`</html>`,
}

func TestStartScriptTags(t *testing.T) {
	for _, line := range openScriptTags {
		if !startOfScriptTag(line) {
			t.Fatal("Should be script tag")
		}
	}
}

func TestStartScriptTagNonTags(t *testing.T) {
	for _, line := range nonScriptTags {
		if startOfScriptTag(line) {
			t.Fatal("Should not be script tag")
		}
	}
}

func TestEndScriptTags(t *testing.T) {
	for _, line := range closeScriptTags {
		if !endOfScriptTag(line) {
			t.Fatal("Should be script tag")
		}
	}
}

func TestEndScriptTagNonTags(t *testing.T) {
	for _, line := range nonScriptTags {
		if endOfScriptTag(line) {
			t.Fatal("Should not be script tag")
		}
	}
}
