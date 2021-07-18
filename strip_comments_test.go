package main

import (
	"strings"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestStripComments
func TestStripComments(t *testing.T) {

	// what we are testing
	tests := []string{
		`
// comment at the beginning
public class AddTwoNumbers {
   /*
    * This method adds two numbers
    */
   public static void main(String[] args) {
      /* single line with multiline style */
      int num1 = 5, num2 = 15, sum;
      sum = num1 + num2;
	  // we print the result
      System.out.println("Sum of these numbers: "+sum);
   }
}`,
		`
import java.util.Scanner;
// CheckEvenOdd checks for even and odd
class CheckEvenOdd
{
  /*
    * The main, master, rule of all functions
    */
  public static void main(String args[])
  {
    int num; // comment on same line of code
    System.out.println("Enter an Integer number:");

    //The input provided by user is stored in num
    Scanner input = new Scanner(System.in);
    num = input.nextInt(); /* the next int... */

    // If number is divisible by 2 then it's an even number
    // else odd number
    if ( num % 2 == 0 )
        System.out.println("Entered number is even");
     else // comment on where it should not be...
        System.out.println("Entered number is odd");
  }
}`,
	}

	// and what we expect
	expected := []string{
		`
public class AddTwoNumbers {
   
   public static void main(String[] args) {
      
      int num1 = 5, num2 = 15, sum;
      sum = num1 + num2;
          
      System.out.println("Sum of these numbers: "+sum);
   }
}`,
		`
import java.util.Scanner;

class CheckEvenOdd
{
  
  public static void main(String args[])
  {
    int num; 
    System.out.println("Enter an Integer number:");

    
    Scanner input = new Scanner(System.in);
    num = input.nextInt(); 

    
    
    if ( num % 2 == 0 )
        System.out.println("Entered number is even");
     else 
        System.out.println("Entered number is odd");
  }
}
`,
	}

	t.Log("Given the need to test we can remove all comments")
	{
		for testID, test := range tests {
			t.Logf("\tTest %d:\tWhen checking the java source", testID)
			{
				response, err := stripComments(test)
				if err != nil {
					t.Fatalf("\t%s\tTest %d:\tShould be able to parse the given string : %v", failed, testID, err)
				}

				if strings.Compare(expected[testID], response) != 0 {
					t.Logf("\t%s\tTest %d:\tShould remove all comments.", succeed, testID)
				} else {
					t.Errorf("\t%s\tTest %d:\tShould remove all comments : %t-%t", failed, testID, test, response)
				}
			}
		}
	}
}
