package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "expandenv"
	app.Usage = "Expand environment variables in input file"
	app.Version = "0.0.1"
	app.Author = "Valentin Hăloiu"
	app.Email = "vially.ichb@gmail.com"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			log.SetFlags(0)
			log.Fatal("Usage: expandenv FILE")
		}

		buf, err := ioutil.ReadFile(c.Args()[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(os.ExpandEnv(string(buf)))
	}

	app.Run(os.Args)
}
