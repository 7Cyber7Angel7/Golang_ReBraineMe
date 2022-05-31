package main

import (
  "fmt"
  "strconv"
)

func main()  {
  a := "104"
  b := 35

  a1, _ := strconv.Atoi(a) // преобразовываем тип String в тип Int
  b1 := strconv.Itoa(b) // преобразовываем тип Int в тип String

  fmt.Println(a1) //    int - 104
  fmt.Println(b1) // string - "35"
}
