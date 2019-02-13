package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func process(startpath string, dirMode os.FileMode, fileMode os.FileMode) error {
	return filepath.Walk(startpath, func(fspc string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		mode := info.Mode()
		if mode.IsDir() {
			fname := filepath.Base(fspc)
			if strings.HasPrefix(fname, "@") {
				if *verbose {
					fmt.Printf(" skip     %s\n", fspc)
				}
				return filepath.SkipDir
			}
			if mode != (dirMode + os.ModeDir) {
				os.Chmod(fspc, dirMode)
				if *verbose {
					fmt.Printf("chmod %#o %s\n", dirMode, fspc)
				}
			}
			return nil
		}
		if mode.IsRegular() {
			if mode != fileMode {
				os.Chmod(fspc, fileMode)
				if *verbose {
					fmt.Printf("chmod %#o %s\n", fileMode, fspc)
				}
			}
		}
		return nil
	})
}

var verbose = flag.Bool("v", false, "run in verbose mode")

func main() {

	fileOct := flag.String("files", "0644", "mode to set for ordinary files")
	dirOct := flag.String("dirs", "0755", "mode to set for directories")

	flag.Parse()

	fileMode, err := strconv.ParseInt(*fileOct, 8, 64)
	if err != nil || fileMode > 0777 {
		fmt.Fprintf(os.Stderr, "bad mode for files %s", fileMode)
	}
	dirMode, err := strconv.ParseInt(*dirOct, 8, 64)
	if err != nil || dirMode > 0777 {
		fmt.Fprintf(os.Stderr, "bad mode for directories %s", fileMode)
	}

	for _, arg := range flag.Args() {
		err := process(arg, os.FileMode(dirMode), os.FileMode(fileMode))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
