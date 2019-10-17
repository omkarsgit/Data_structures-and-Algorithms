package main

import (
  "fmt"
)

func main() {
  //Created an array of numbers
  numbers := [6]int{2, 3, 5, 7, 11, 13}

  //scanning input from the user
  var input int
  fmt.Println("Enter an integer value: ")
  _, err := fmt.Scanf("%d", &input)

  //Checking for errors
  if err != nil {
    fmt.Println(err)
  }

  //Iterating through all elements in array numbers
  for i := 0; i < 6; i++ {
    //Checking if the input is present in that array
    if numbers[i] == input {
      fmt.Printf("Found %d\n", input)
      return
    }
  }
  fmt.Printf("%d not found.\n", input)
  return
}
