package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type flags struct {
	sl, d, f bool
	ext      string
	filepath string
}

var fl flags

func init() {
	flag.BoolVar(&fl.sl, "sl", false, "sl flag")
	flag.BoolVar(&fl.d, "d", false, "d flag")
	flag.BoolVar(&fl.f, "f", false, "f flag")
	flag.StringVar(&fl.ext, "ext", "", "ext flag")
}

func main() {
	flag.Parse()
	if err := initFlags(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fl)
	err := filepath.Walk(fl.filepath, walker)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func initFlags() error {
	args := flag.Args()
	if len(args) > 0 {
		fl.filepath = args[0]
	} else {
		return fmt.Errorf("missing folder path")
	}
	if (!fl.f || fl.d || fl.sl) && fl.ext != "" {
		return fmt.Errorf("-ext works ONLY when -f is specified")
	}
	if !fl.sl && !fl.d && !fl.f {
		fl.sl = true
		fl.d = true
		fl.f = true
	}
	return nil
}

func walker(path string, file os.FileInfo, err error) error {
	if os.IsPermission(err) {
		return filepath.SkipDir
	} else if err != nil {
		return err
	}
	if strings.HasPrefix(file.Name(), ".") {
		if file.IsDir() {
			return filepath.SkipDir
		}
		return nil
	}
	if fl.ext != "" {
		if file.Mode().IsRegular() {
			fileext := strings.TrimLeft(strings.ToLower(filepath.Ext(path)), ".")
			if fileext == fl.ext {
				fmt.Println(path)
			}
		}
	} else {
		if fl.d && file.Mode().IsDir() {
			fmt.Println(path)
		}
		if fl.f && file.Mode().IsRegular() {
			fmt.Println(path)
		}
		if fl.sl && file.Mode()&os.ModeSymlink != 0 {
			symlinkWork(path)
		}
	}
	return nil
}

func symlinkWork(path string) {
	link, err := os.Readlink(path)
	if err == nil {
		_, err2 := os.Stat(path)
		err = err2
	}
	if err != nil {
		fmt.Printf("%s -> [broken]\n", path)
	} else {
		fmt.Printf("%s -> %s\n", path, link)
	}
}
