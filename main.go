package main

import (
	"github.com/Lateks/cotsworth/cmd"
	"os"
)

func main() {
	cmd.Execute(os.Args[1:])
}
