package main

import (
	"log"
	"os"

	dduper "github.com/twinsnes/dduper/src"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "dDuper",
		Usage: "Removes duplicate files in folders and subfolders",
		Version: "v0.0.1",
		Action: func(*cli.Context) error {
			return dduper.DDupeCurrentDir(true)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
