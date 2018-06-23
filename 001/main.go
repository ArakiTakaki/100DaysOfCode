package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	config(app)
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "cat, c",
			Usage: "Echo with cat",
		},
		cli.BoolFlag{
			Name:  "dog, d",
			Usage: "Echo with cat",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "cat",
			Usage:  "Echo with cat",
			Action: test2,
		},
		{
			Name:   "dog",
			Usage:  "Echo with dog",
			Action: test3,
		},
	}

	app.Run(os.Args)
}

func test1(c *cli.Context) error {
	if c.Bool("cat") {
		fmt.Println(c.Args().Get(0) + "にゃご")
		return nil
	}
	if c.Bool("dog") {
		fmt.Println(c.Args().Get(0) + "がおー")
		return nil

	}
	fmt.Println(c.Args().Get(0))
	return nil
}

func test2(c *cli.Context) error {
	fmt.Println(c.Args().Get(0) + "だにゃん♡")
	return nil
}

func test3(c *cli.Context) error {
	fmt.Println(c.Args().Get(0) + "わん☆")
	return nil
}
