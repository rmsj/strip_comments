package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	s := `
// code stats with comments
// that goes to multi line
class Incrementer {
  int count = 0; // keeps a count
  string test = "http://test.com"
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

	var result string

	// we need an io.Reader
	source := strings.NewReader(s)
	reader := bufio.NewReader(source)

	// try to get the first character
	charCode, err := reader.ReadByte()

	if err != nil {
		log.Fatal("Cannot read first character of string")
	}

	// loop through every character in string
	for err == nil {
		// byte to string so we can check for comments
		char := string(charCode)

		result += char

		charCode, err = reader.ReadByte()
	}

	return result, nil
}
