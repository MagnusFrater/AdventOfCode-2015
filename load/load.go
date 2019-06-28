package load

import (
	"io/ioutil"
)

// InputFileAsString returns the entire contents of the input file as a string.
func InputFileAsString(day string) string {
	buf, err := ioutil.ReadFile("input" + day + ".txt")
	if err != nil {
		panic(err)
	}

	return string(buf)
}
