# expandenv

Replace environment variables from input file

## Usage
`expandenv FILE`

## Example

```
$ cat input.txt
Hello $GREETER_NAME
$ export GREETER_NAME=John
$ expandenv input.txt > output.txt
$ cat output.txt
Hello John
```
