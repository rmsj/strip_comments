package main

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	javaCode = iota
	firstSlash
	singleLineComment
	multiLineCommentStart
	multiLineCommentEnd
	stringLiteral
)

// strCache keeps a cache of non-comment characters to be written back to result
var strCache string

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
	result, _ := stripComments(s)

	fmt.Println(result)
}

func stripComments(s string) (string, error) {

	// sourceCodeType identifies in what "type" of code we are currently going through
	sourceCodeType := javaCode
	var result string
	// gets one byte at a time from the string
	var charCode byte
	// error when reading the byte - EOF
	var err error

	// we need an io.Reader
	source := strings.NewReader(s)
	reader := bufio.NewReader(source)
	//buf := make([]byte, len(s))

	// loop through every character in string
	for {

		// read the string byte by byte - try to get the first one
		charCode, err = reader.ReadByte()

		// end of the string.
		if err != nil {

			// Flush the reset of the bytes we have.
			//output.WriteByte(buf[:end])
			break
		}

		// byte to string so we can check for comments
		char := string(charCode)

		switch sourceCodeType {

		// normal source code
		case javaCode:
			switch char {
			case "\"", "'":
				// we are starting to read a string literal on the source code
				sourceCodeType = stringLiteral
				result += char
			case "/":
				// we found a slash outside a string literal, check for single or multiline comment
				sourceCodeType = firstSlash
				strCache += char
			default:
				result += fmt.Sprintf("%s%s", strCache, char)
				strCache = ""
			}
			strCache = ""
		// on the loop iteration for firstSlash we are looking for single or multi line comments
		case firstSlash:
			switch char {
			case "/":
				sourceCodeType = singleLineComment
				strCache = ""
			case "*":
				sourceCodeType = multiLineCommentStart
				strCache = ""
			default:
				sourceCodeType = javaCode
				result += fmt.Sprintf("%s%s", strCache, char)
				strCache = ""
			}
		case singleLineComment:
			// once this is true, we keep on it until line break
			if char == "\n" || char == "\r" {
				sourceCodeType = javaCode
				result += char
			}
		case multiLineCommentStart:
			// for multiline comments, we look for another star character and send
			// the status to star - to look for a closing "/" for the comment
			if char == "*" {
				sourceCodeType = multiLineCommentEnd
			}
		case multiLineCommentEnd:
			switch char {
			// if after a star character we hit a closing / it means the end of the multiline comment
			case "/":
				sourceCodeType = javaCode
			default:
				// otherwise we keep going - we are still inside the multiline character
				sourceCodeType = multiLineCommentStart
			}
			// anything inside a " or a '
		case stringLiteral:
			switch char {
			// we are closing the string literal
			case "\"", "'":
				sourceCodeType = javaCode
			}
			result += char
		}
	}

	return result, nil
}
