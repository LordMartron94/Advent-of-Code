package shared

import (
	"bufio"
	"io"
)

func ReaderToRunes(reader io.Reader) ([]rune, error) {
	bReader := bufio.NewReader(reader)
	runes := make([]rune, 0)

	for {
		scannedRune, _, err := bReader.ReadRune()

		if err == io.EOF {
			break
		} else if err != nil {
			return runes, err
		}

		runes = append(runes, scannedRune)
	}

	return runes, nil
}
