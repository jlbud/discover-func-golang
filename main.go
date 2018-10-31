package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		name := ""
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}


	app.Run(os.Args)
}
