package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

/*
func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}
*/

func valIP(input string) bool {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	var bit []string = strings.Split(input, ".")
	if len(bit) < 4 {
		return false
	}
	for _, val := range bit {
		if i, _ := strconv.Atoi(val); i > 255 {
			return false
		}
	}
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.MatchString(input)
}
func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	for _, filename := range arguments[1:] {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			os.Exit(-1)
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
				break
			}
			r := valIP(line)
			if !r {
				fmt.Println("Incorrect IP Address: ", line)
				continue
			} else {
				fmt.Println("Correct IP aaddress: ", line)
			}

		}
	}
}
