package main

import (
	"github.com/urfave/cli"
	"github.com/ungerik/go-dry"
	"path/filepath"
	"log"
	"os"
	"sort"
	"github.com/spf13/cast"
)

func renameFiles(dir string) {
	if !dry.FileIsDir(dir) {
		return
	}

	abs, _ := filepath.Abs(dir)
	b := filepath.Base(abs)
	files, _ := dry.ListDirFiles(abs)
	sort.Strings(files)

	for idx, v := range files {
		oldFile := filepath.Join(abs, v)
		ext := filepath.Ext(v)
		newFile := filepath.Join(abs, b+`_`+cast.ToString(idx)+ext)
		//dry.FileTouch(newFile)
		err := os.Rename(oldFile, newFile)
		if err != nil {
			log.Println(err)
		}
	}

	subDirs, _ := dry.ListDirDirectories(abs)
	for _, sub := range subDirs {
		renameFiles(filepath.Join(abs, sub))
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
		log.Println("rename files in", path)
		//renameFiles(`D:\ffres\4.6大厅\4.6动画序列帧`)
		renameFiles(path)

		//subDirs, _ := dry.ListDirDirectories(path)
		//for _, sub := range subDirs {
		//	renameFiles(sub)
		//}

		//cmd := exec.Command("texturemerger")
		//cmd.Dir = h5Path
		//out, err := cmd.Output()
		//if err != nil {
		//	sugar.Fatal(err)
		//}

		return nil
	}

	app.Run(os.Args)
}
