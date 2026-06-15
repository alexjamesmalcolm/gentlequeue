package main

import (
	"os"

	"go.followtheprocess.codes/cli"
)

func main() {
	var title string
	_, err := cli.New("add",
		cli.Flag(&title, "title", 't', "Title of prayer"),
	)
	if err != nil {
		os.Exit(1)
	}
}
