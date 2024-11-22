package main

import (
	"os"
	"strings"

	"github.com/tushar2708/altcsv"
)

func parseEnvLine(line string) (string, string) {
	parts := strings.Split(line, "=")
	key := parts[0]
	value := strings.Join(parts[1:], "=")
	return key, value
}

func main() {
	csvWriter := altcsv.NewWriter(os.Stdout)
	csvWriter.AllQuotes = true
	defer csvWriter.Flush()
	csvWriter.Write([]string{"name", "value"})
	for _, line := range os.Environ() {
		key, value := parseEnvLine(line)
		err := csvWriter.Write([]string{key, value})
		if err != nil {
			panic(err)
		}
	}
}
