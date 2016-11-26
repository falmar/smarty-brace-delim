// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"os"
	"testing"
)

func TestBackup(t *testing.T) {
	path := "files/simple_bracket.tpl"
	expBackup := "files/simple_bracket_backup.tpl"

	backup, err := createBackup(path, "_backup")
	if err != nil {
		t.Fatalf("Error during backup: %s", err)
	}

	if backup != expBackup {
		t.Fatalf("Expected backup name: %s; got : %s", expBackup, backup)
	}

	inputFile, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}

	backupFile, err := os.Open(backup)
	defer backupFile.Close()
	if err != nil {
		t.Fatal(err)
	}

	inputStat, err := inputFile.Stat()
	if err != nil {
		t.Fatal(err)
	}

	backupStat, err := backupFile.Stat()
	if err != nil {
		t.Fatal(err)
	}

	if inputStat.Size() != backupStat.Size() {
		t.Fatalf("Expected filesize: %d; got: %d", inputStat.Size(), backupStat.Size())
	}
}
