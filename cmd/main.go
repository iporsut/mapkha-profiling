package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"

	m "github.com/veer66/mapkha"
)

var dixPath, cpuProfileFile string

func init() {
	flag.StringVar(&dixPath, "dix", "", "Dictionary path")
	flag.StringVar(&cpuProfileFile, "cpupprof", "cpu.pprof", "CPU PPROF file")
}

func main() {
	flag.Parse()

	f, err := os.Create(cpuProfileFile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	var dict *m.Dict
	if dixPath == "" {
		dict, err = m.LoadDefaultDict()
	} else {
		dict, err = m.LoadDict(dixPath)
	}
	if err != nil {
		log.Fatal(err)
	}

	wordcut := m.NewWordcut(dict)
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("could not read input:", err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))
	var countWord int
	for scanner.Scan() {
		segment := wordcut.Segment(scanner.Text())
		countWord += len(segment)
	}
	fmt.Println(countWord)
}
