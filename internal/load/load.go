package load

import (
	"io/ioutil"
	"log"
	"os"
)

// InputFileAsString returns the entire contents of the input file as a string.
func inputFileAsString(inputFilePath string) string {
	buf, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}

	return string(buf)
}

// Input returns the entire contents of the command args input file as a string.
func Input(day string) string {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./day0" + day + " path/to/input")
	}

	return inputFileAsString(os.Args[1])
}
