package env_parser

import (
	"bufio"
	"io"
	"strings"
)

// GetKeyValuePairs reads the contents of the given environment file and returns a map of key-value pairs.
func GetKeyValuePairs(envFile io.Reader) map[string]string {
	pairs := make(map[string]string)
	scanner := bufio.NewScanner(envFile)

	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == '#' {
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {
			pairs[parts[0]] = parts[1]
		}
	}

	return pairs
}
