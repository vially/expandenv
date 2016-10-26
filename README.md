# expandenv

Replace environment variables from input file

## Usage
`expandenv [--base64] FILE`

## Installation

Grab a binary from [Releases](https://github.com/vially/expandenv/releases) or build from source:
`go get github.com/vially/expandenv`

## Basic Example

```bash
$ cat input.txt
Hello $GREETER_NAME
$ export GREETER_NAME=John
$ expandenv input.txt > output.txt
$ cat output.txt
Hello John
```

## Base64 Encode Example

```bash
$ cat secrets.yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  password: $PASSWORD
$ export PASSWORD=s3cret
$ expandenv --base64 secrets.yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  password: czNjcmV0
```
