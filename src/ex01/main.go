package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
)

type flags struct {
	lines, characters, words bool
}

var fl flags

func init() {
	flag.BoolVar(&fl.lines, "l", false, "l flag")
	flag.BoolVar(&fl.characters, "m", false, "m flag")
	flag.BoolVar(&fl.words, "w", false, "w flag")
}

func main() {
	if err := initFlags(); err != nil {
		fmt.Println(err)
		return
	}
	wg := new(sync.WaitGroup)
	for _, filename := range flag.Args() {
		wg.Add(1)
		go printWithCount(filename, wg)
	}
	wg.Wait()
}

func initFlags() error {
	flag.Parse()
	if flag.NFlag() > 1 {
		return fmt.Errorf("only one can be specified at a time")
	} else if flag.NFlag() == 0 {
		fl.words = true
	}

	return nil
}

func printWithCount(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)

		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	if fl.words {
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			count++
		}
		fmt.Printf("%d\t%s\n", count, filename)
	} else if fl.lines {
		for scanner.Scan() {
			count++
		}
		fmt.Printf("%d\t%s\n", count, filename)
	} else if fl.characters {
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			count++
		}
		fmt.Printf("%d\t%s\n", count, filename)
	}
	if err = scanner.Err(); err != nil {
		fmt.Println(err)

		return
	}
}
