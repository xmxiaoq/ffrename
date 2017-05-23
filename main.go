package main

import (
	"github.com/urfave/cli"
	"github.com/ungerik/go-dry"
	"path/filepath"
	"log"
	"os"
	"strings"
)

func renameFiles(dir string) {
	if !dry.FileIsDir(dir) {
		return
	}

	abs, _ := filepath.Abs(dir)
	b := filepath.Base(abs)
	files, _ := dry.ListDirFiles(dir)
	for _, v := range files {
		if strings.HasPrefix(v, b) {
			continue
		}
		oldFile := filepath.Join(dir, v)
		newFile := filepath.Join(abs, b+`_`+v)
		//dry.FileTouch(newFile)
		err := os.Rename(oldFile, newFile)
		if err != nil {
			log.Println(err)
		}
	}

	subDirs, _ := dry.ListDirDirectories(dir)
	for _, sub := range subDirs {
		renameFiles(sub)
	}
}
func main() {
	var path string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path,p",
			Value:       ".",
			Usage:       "path for the greeting",
			Destination: &path,
		},
	}

	app.Action = func(c *cli.Context) error {
		renameFiles(path)

		return nil
	}

	app.Run(os.Args)
}
