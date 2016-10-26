package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jawher/mow.cli"
)

func main() {
	app := cli.App("expandenv", "Expand environment variables in input file")
	app.Spec = "[-b] FILE"
	app.Version("v version", "0.0.2")

	var (
		encode = app.BoolOpt("b base64", false, "base64 encode values")
		src    = app.StringArg("FILE", "", "input file")
	)

	app.Action = func() {
		buf, err := ioutil.ReadFile(*src)
		if err != nil {
			log.Fatal(err)
		}

		fileData := string(buf)
		var mapper func(string) string
		if !*encode {
			mapper = os.Getenv
		} else {
			mapper = envBase64Encoder
		}
		fmt.Print(os.Expand(fileData, mapper))
	}

	app.Run(os.Args)
}

func envBase64Encoder(key string) string {
	return base64.StdEncoding.EncodeToString([]byte(os.Getenv(key)))
}
