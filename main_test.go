// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"os"
	"testing"
)

func getCommonFlags() map[string]interface{} {
	return map[string]interface{}{
		"backupSuffix": "_backup",
		"inputPath":    "",
		"outputPath":   "",
		"removeBackup": false,
		"overWrite":    false,
		"brace":      false,
		"delim":        false,
	}
}

func TestMainCheckParseOptionNone(t *testing.T) {
	cflags := getCommonFlags()
	expCode := 1
	expErr := "Must choose an type of action delim or brace parse"
	code, err := altMain(cflags)

	if code != expCode {
		t.Fatalf("Expected exit code: %d; got: %d", expCode, code)
	}

	if err == nil {
		t.Fatal("Expected error to not be nil")
	}

	if err.Error() != expErr {
		t.Fatalf("Expected error: %s; got: %s", expErr, err.Error())
	}
}

func TestMainCheckParseOptionBoth(t *testing.T) {
	cflags := getCommonFlags()
	cflags["delim"] = true
	cflags["brace"] = true

	expCode := 1
	expErr := "Must choose between delim or brace parse, not both"

	code, err := altMain(cflags)

	if code != expCode {
		t.Fatalf("Expected exit code: %d; got: %d", expCode, code)
	}

	if err == nil {
		t.Fatal("Expected error to not be nil")
	}

	if err.Error() != expErr {
		t.Fatalf("Expected error: %s; got: %s", expErr, err.Error())
	}
}

func TestMainDontOverwrite(t *testing.T) {
	cflags := getCommonFlags()
	cflags["delim"] = true
	cflags["overWrite"] = false
	cflags["inputPath"] = "files/simple_brace.tpl"
	cflags["outputPath"] = "files/simple_brace_tm1.tpl"
	cflags["backupSuffix"] = "_tm1_backup"

	expCode := 2
	expErr := "Error ocurred during backup creation: Backup file already exist"

	_, err := os.Create("files/simple_brace_tm1_backup.tpl")
	if err != nil {
		t.Fatal(err)
	}

	code, err := altMain(cflags)

	if code != expCode {
		t.Fatalf("Expected exit code: %d; got: %d", expCode, code)
	}

	if err == nil {
		t.Fatal("Expected error to not be nil")
	}

	if err.Error() != expErr {
		t.Fatalf("Expected error: %s; got: %s", expErr, err.Error())
	}
}

func TestMainOverwrite(t *testing.T) {
	cflags := getCommonFlags()
	cflags["delim"] = true
	cflags["overWrite"] = true
	cflags["inputPath"] = "files/simple_brace.tpl"
	cflags["outputPath"] = "files/simple_brace_tm2.tpl"
	cflags["backupSuffix"] = "_tm2_backup"

	expCode := 0

	_, err := os.Create("files/simple_brace_tm2_backup.tpl")

	if err != nil {
		t.Fatal(err)
	}

	code, err := altMain(cflags)

	if code != expCode {
		t.Fatalf("Expected exit code: %d; got: %d", expCode, code)
	}

	if err != nil {
		t.Fatalf("Expected error to be nil; got: %s", err)
	}
}

func TestMainDontRemoveBackup(t *testing.T) {
	cflags := getCommonFlags()
	cflags["brace"] = true
	cflags["overWrite"] = true
	cflags["removeBackup"] = false
	cflags["inputPath"] = "files/simple_delim.tpl"
	cflags["outputPath"] = "files/simple_delim_tm3.tpl"
	cflags["backupSuffix"] = "_tm3_backup"

	expCode := 0

	code, err := altMain(cflags)

	if code != expCode {
		t.Fatalf("Expected exit code: %d; got: %d", expCode, code)
	}

	if err != nil {
		t.Fatalf("Expected error to be nil; got: %s", err)
	}

	_, err = os.Stat("files/simple_delim_tm3_backup.tpl")

	if os.IsNotExist(err) {
		t.Fatal("Expected backup file to exist")
	}
}

func TestMainRemoveBackup(t *testing.T) {
	cflags := getCommonFlags()
	cflags["brace"] = true
	cflags["overWrite"] = true
	cflags["removeBackup"] = true
	cflags["inputPath"] = "files/simple_delim.tpl"
	cflags["outputPath"] = "files/simple_delim_tm4.tpl"
	cflags["backupSuffix"] = "_tm4_backup"

	expCode := 0

	code, err := altMain(cflags)

	if code != expCode {
		t.Fatalf("Expected exit code: %d; got: %d", expCode, code)
	}

	if err != nil {
		t.Fatalf("Expected error to be nil; got: %s", err)
	}

	_, err = os.Stat("files/simple_delim_tm4_backup.tpl")

	if !os.IsNotExist(err) {
		t.Fatal("Expected backup file to not exist")
	}
}
