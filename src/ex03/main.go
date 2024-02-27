package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

const err1 = "missing folder"
const err2 = "missing log files"
const err3 = "missing log files for rotate"

var flagA bool
var logFiles []string
var archivePath string

func init() {
	flag.BoolVar(&flagA, "a", false, "a flag for directory")
}

func main() {
	if err := initFlags(); err != nil {
		println(err.Error())

		return
	}
	rotate()
}

func initFlags() error {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 && flagA {

		return fmt.Errorf(err1)
	}
	if len(args) < 2 && flagA {

		return fmt.Errorf(err2)
	}
	if len(args) < 1 {

		return fmt.Errorf(err3)
	}

	if flagA {
		archivePath = "." + args[0] + "/"
		logFiles = args[1:]
	} else {
		archivePath = "./"
		logFiles = args[0:]
	}
	return nil
}

func rotate() {
	var wg sync.WaitGroup
	for _, elem := range logFiles {
		wg.Add(1)
		go archiveCreator(elem, &wg)
	}
	wg.Wait()
}

func archiveCreator(logfileName string, wg *sync.WaitGroup) {
	defer wg.Done()
	src := logfileName
	file, err := os.Stat(logfileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := os.MkdirAll(archivePath, 0755); err != nil && !os.IsExist(err) {
		fmt.Println(err)
		return
	}
	parts := strings.Split(file.Name(), ".")
	ModTime := strconv.FormatInt(file.ModTime().Unix(), 10)
	dst := archivePath + parts[0] + ModTime + ".tar.gz"

	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstFile.Close()

	gw := gzip.NewWriter(dstFile)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	header := &tar.Header{
		Name: file.Name(),
		Mode: int64(file.Mode()),
		Size: file.Size(),
	}

	if err := tw.WriteHeader(header); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := io.Copy(tw, srcFile); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(logfileName, " archived")
}
