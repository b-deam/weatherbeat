package main

import (
	"os"

	"github.com/b-deam/weatherbeat/cmd"

	_ "github.com/b-deam/weatherbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
