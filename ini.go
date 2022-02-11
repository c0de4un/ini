package ini

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS.PUBLIC
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Reader struct {
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// INTERFACES.PUBLIC
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type IReadListener interface {
	// Return false to stop parse
	OnParam(name string, value string) bool
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// METHODS.PUBLIC
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func NewReader() Reader {
	var reader Reader

	return reader
}

func (reader Reader) ReadAll(filePath string, listener IReadListener) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var paramName string
	var paramValue string
	var equalIndex int
	var line string
	var lineIndex int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) < 1 {
			continue
		}

		line = strings.ReplaceAll(line, " ", "")

		if strings.IndexRune(line, ';') == 0 {
			continue
		}

		equalIndex = strings.IndexRune(line, '=')
		if equalIndex < 1 {
			return fmt.Errorf("invalid ini format: line %d at %d", lineIndex, equalIndex)
		}

		if strings.IndexRune(line, ';') > 0 {
			return fmt.Errorf("invalid ini format, comment in param: line %d at %d", lineIndex, equalIndex)
		}

		paramName = line[0:equalIndex]
		paramValue = line[equalIndex+1:]

		if !listener.OnParam(paramName, paramValue) {
			break
		}

		lineIndex++
	}

	return nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
