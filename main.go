//Installing Golang https://golang.org/ is the site from where to download Golang.

package main 
import "fmt"
//The main package is the default golang package while the fmt package allows one to format text and print to console.

func main(){
 
  //In Golang there are Two ways to declare a Variable 
  var name string ="RayM"
  fmt.Println(name)
// RayM is printed out
number:=42
  fmt.Println(number)
  //42 is printed out


  
  //format specifiers are used in Golang to print out the value of a Variable in a formatted way ; Hence use of fmt.Printf and not fmt.Println

//There are a number of different format specifiers that can be used with the % format. Some of the most common ones are:

//%d - Prints an integer in decimal format.
//%f - Prints a floating-point number in decimal format.
//%s - Prints a string.
//%T - Prints the type of the value being printed.
//%+v - Prints the value in a more verbose format, including the field names of structs.
//%#v - Prints the value in a Go syntax representation

fmt.Println("hello world \nhello world")
  //the code prints out "hello world" twice in separate lines

// Format a decimal integer as a string.
fmt.Printf("%d\n", 12345)
  //prints out 12345

// Format a floating-point number as a string with a precision of 3.
fmt.Printf("%.3f\n", 12.345678)
  //prints out 12.346

// Format a string as a string.
fmt.Printf("%s\n", "Hello, world!")
  //prints out Hello, world!

// Format a boolean value as a string.
fmt.Printf("%t\n", true)
  //prints out true

//Practice Exercise. Evaluate the following.
fmt.Printf("The value is: %d\n", 123)
fmt.Printf("The value is: %f\n", 1.2345)
fmt.Printf("The value is: %s\n", "Hello, world!")
fmt.Printf("The type of the value is: %T\n", 123)
fmt.Printf("The value in a verbose format is: %+v\n", struct {
  Name string
  Age int
}{Name: "John Doe", Age: 30})
fmt.Printf("The value in a Go syntax representation is: %#v\n", struct {
  Name string
  Age int
}{Name: "John Doe", Age: 30})



//Mathematical Concepts in Golang.
  // The main rule is that all the variables must be of the same type , That is, An integer and a float can not be computed together, one must be changed to match the other
  
  //Example 1, Integers
  var Father int = 45
  var son int = 15
  fmt.Printf("son was born when father was %d\n",Father-son)
  //Example 2, Float values 
  nitrogen := 77.6
  oxygen:= 21.9
  fmt.Printf("%g\n", nitrogen + oxygen)
  //Example 3, mixed values
  Tea := 22
  Coffee := 10.2
  fmt.Printf("%g\n", float64(Tea)+Coffee)
  //example 1 reads out "son was       born when father was 30"
  //example 2 prints out 99.5
  //example 3 prints out 32.2

  

//AND,OR,NOT statements
 // "||" is used to represent OR statement, one of the expressions must be true for the OR statement to be valid
 // "&&" used to show AND statement, All of the expressions must be true for the AND statement to be valid
  // "!" used to represent NOT i.e the opposite


//IF,ELSE IF,ELSE
  //if-Executes a code block if condition evaluates to true
  //else if-Executes a code block only when previous criterion (if statement ) evaluated to false
  //else- Executes code block when none of the previous conditions are met


age:= 25
salary:= 100000
engaged:= true
  if (age>35 || engaged==true ) && salary >= 100000{
    fmt.Println("You can vie for Presidency")
  } else if age>=18 && salary>= 75000{
    fmt.Println("you can vie for governor")
  }else {
    fmt.Println("You can't run for any office")
  }
//This code block will print the line "you can vie for presidency" since the or expression in brackets is evaluated to true and the and expression is also evaluated to true.
  //However if the age variable was equals to 17 both the if and else if conditions would not have been met forcing the code block to print the else condition "You can't run for any office"
  


// SWITCH statements
//switch statements are used to print out statements that accomodate specific instances, unlike if ,else if and else statements 
//In using switch statements in golang the variable and case must be of the same type , for example if one is an integer the other must also be an integer 


year:=2000
  switch year{
    case 1:1985
      fmt.Println("pass")
    case 2:1995
      fmt.Println("second pass")
    case 3:2000
      fmt.Println("Perfect")
    default:
      fmt.Println("No suitable candidate")
  //}
  //The code will print out the word perfect since the year variable matches with case 3. If the variable had no matching case it would print out the default.




// For Loops
  //Golang uses For loops and Has No While loops
  //For loops are used for easier iteration of a code block
//The basic for loop structure is:
  //for variable; condition; x+= 10
    //fmt.Println(x)
for x:=10 ; x<40 ; x+= 10  {
  fmt.Println(x)}
  // This will print out 10,20,30
//BREAK and CONTINUE statements
  //Break statement will immediately exit the loop and move to the next code block, while a Continue statement will execute the continuing expressions of the loop
for i := 1; i <= 10; i++ {
 if i == 5 {
   break
}
  fmt.Println(i)
    }
//the statement will print out 1 2 3 4
  // the code prints numbers 1 to 4 before being broken

for k := 1; k <= 10; k++ {
 if k == 5 {
   continue
        }
   fmt.Println(k)
    }
//the statement will print out 1 2 3 4 6 7 8 9 10
  //the code prints 1 to 4 from where the continue expression is met hence jumping back to the code and printing from 6 to 10






}


