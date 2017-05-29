# smarty-brace-delim

[![Status](https://travis-ci.org/falmar/smarty-brace-delim.svg?branch=master)](https://travis-ci.org/falmar/smarty-brace-delim) [![Codecov](https://img.shields.io/codecov/c/github/falmar/smarty-brace-delim.svg)](https://codecov.io/gh/falmar/smarty-brace-delim)

Replace `{rdelim}` and `{ldelim}` Smarty braces delimiters for real braces `{` and `}` or viceversa from Smarty templates


## Install

`$ go get github.com/falmar/smarty-brace-delim`

## Usage

```
$ smarty-brace-delim -h
Usage of smarty-brace-delim:
  -b	Parse braces into {delim}
  -d	Parse {delim} into braces
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

`$ go test github.com/falmar/smarty-brace-delim -v -cover`

## Examples

Using option `-b`

```
$ smarty-brace-delim -i path/to/file -o path/to/output_file -b
```

Will transform the [input file](https://github.com/falmar/smarty-brace-delim/blob/master/files/simple_brace.tpl) into [output file](https://github.com/falmar/smarty-brace-delim/blob/master/files/simple_delim.tpl)

Using the option `-d` will do the opposite

## TODO

- [x] Take care of fragments multiline comments eg. `function { {* comment *}   }`
