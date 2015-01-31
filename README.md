# expandenv

Replace environment variables from input file

## Usage
`expandenv FILE`

## Installation

Grab a binary from [Releases](https://github.com/vially/expandenv/releases) or build from source:
`go get github.com/vially/expandenv`

## Example

```bash
$ cat input.txt
Hello $GREETER_NAME
$ export GREETER_NAME=John
$ expandenv input.txt > output.txt
$ cat output.txt
Hello John
```
