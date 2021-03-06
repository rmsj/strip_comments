package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := `
// code stats with comments
// that goes to multi line
class Incrementer {
  int count = 0; // keeps a count
  string test = "http:/\/test.com"
  string test2 = 'http://test.com'
  /*
   * This method increments the count
   */
  public void inc() {
    // increment the counter here
    count++;
  } /* end of method */
}`
	result, _ := StripComments(s)

	fmt.Println(strings.TrimSpace(result))
}

// StripComments read a string that contains a source code of a java class, function
// and removes the single line and multiline comments
func StripComments(s string) (string, error) {

	// to save the result as we build it char by char
	var result string
	var singleLine string
	// gets one byte at a time from the string
	var charCode byte
	// error when reading the byte - EOF
	var err error

	var isSingleLineComment bool
	var isMultilineComment bool
	var isStringLiteral bool

	// we need an io.Reader
	source := strings.NewReader(s)

	// loop through every character in string
	for {
		// read the string byte by byte - try to get the first one
		charCode, err = source.ReadByte()

		// end of the string.
		if err != nil {
			break
		}

		// byte to string so we can check for comments
		char := string(charCode)

		switch char {
		case "\"", "'":
			isStringLiteral = !isStringLiteral
		}

		// check for the start of a comment - either multi line of single line
		if !isStringLiteral && char == "/" {
			nextChar, err := source.ReadByte()
			if err != nil {
				break
			}

			switch string(nextChar) {
			case "/":
				isSingleLineComment = true
			case "*":
				isMultilineComment = true
			}
		}

		// for single line comments, we keep reading until we find a line break
		if isSingleLineComment {
			for isSingleLineComment {
				slChar, err := source.ReadByte()
				if err != nil {
					break
				}
				if string(slChar) == "\n" || string(slChar) == "\r" {
					char = string(slChar)
					isSingleLineComment = false
				}
			}
		}

		// for multiple line comments, we keep reading until we find the "close" of the comment
		if isMultilineComment {
			for isMultilineComment {
				mlcChar, err := source.ReadByte()
				if err != nil {
					break
				}
				if string(mlcChar) == "*" {
					nextMlcChar, err := source.ReadByte()
					if err != nil {
						break
					}
					if string(nextMlcChar) == "/" {
						char = ""
						isMultilineComment = false
					}
				}
			}
		}

		// build each line or on line break, add each line to result
		switch char {
		case "\n", "\r":
			//if len(strings.TrimSpace(singleLine)) > 0 {
			result += singleLine + char
			singleLine = ""
			//}
		default:
			singleLine += char
		}
	}

	// we have something, even if it is empty and reached the end of the string
	if err == io.EOF {
		err = nil
	}

	// the remainder of chars, if any
	if len(singleLine) > 0 {
		result += singleLine
	}

	return result, err
}
