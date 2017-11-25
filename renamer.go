package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	globInfoUrl = "https://golang.org/pkg/path/filepath/#Match"
)

var (
	atCharactersRegex = regexp.MustCompile("@+")
)

func main() {
	var glob = flag.String("p", "", "Glob file pattern")
	var sheme = flag.String("n", "", "New naming sheme")
	flag.Parse()

	if *glob == "" || *sheme == "" {
		log.Println("-p and -n are required flags for this program")
		os.Exit(1)
	}

	if !strings.Contains(*sheme, "@") {
		log.Println(
			"Your new file name din't contain @ characters, it needs those to know wher to put in the count numbers",
		)
		os.Exit(2)
	}

	matches, err := filepath.Glob(*glob)
	if err != nil {
		panic(err)
	}

	if len(matches) == 0 {
		log.Printf(
			"Your pattern (%s) din't match any files, please see this for info on how to write it: %s\n",
			*glob,
			globInfoUrl,
		)
	}

	for i, fileName := range matches {
		newName := generateNewName(*sheme, i+1)
		fmt.Printf("Renaming %s to %s: ", fileName, newName)
		err = os.Rename(fileName, newName)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("OK")
		}
	}
}

func generateNewName(sheme string, counter int) string {
	atCharacters := atCharactersRegex.FindString(sheme)
	paddedCount := lpad(strconv.Itoa(counter), "0", len(atCharacters))

	return strings.Replace(sheme, atCharacters, paddedCount, -1)
}

func lpad(stringToPad string, padding string, paddingLenght int) string {
	for i := len(stringToPad); i < paddingLenght; i++ {
		stringToPad = padding + stringToPad
	}
	return stringToPad
}
