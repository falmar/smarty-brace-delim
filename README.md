# smarty-bracket-delim

[![Status](https://travis-ci.org/falmar/smarty-bracket-delim.svg?branch=master)](https://travis-ci.org/falmar/smarty-bracket-delim) [![Codecov](https://img.shields.io/codecov/c/github/falmar/smarty-bracket-delim.svg)](https://codecov.io/gh/falmar/smarty-bracket-delim)

Replace `{rdelim}` and `{ldelim}` Smarty brackets delimiters for real brackets `{` and `}` or viceversa from Smarty templates


## Install

`$ go get github.com/falmar/smarty-bracket-delim`

## Usage

```
$ smarty-bracket-delim -h
Usage of smarty-bracket-delim:
  -b	Parse brackets into {delim}
  -d	Parse {delim} into brackets
  -i string
    	Input file path
  -o string
    	Output file path (if not provied will overwrite input file)
  -ow
    	Overwrite backup file if already exist
  -rm
    	Remove backup file after parse
```

## Test

`$ go test github.com/falmar/smarty-bracket-delim -v -cover`


## TODO

- Take care not parse anything inside {php} {/php} tags within <script> tag
