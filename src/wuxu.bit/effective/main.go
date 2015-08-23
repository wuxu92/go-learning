package main

import (
  "fmt"
)

func main() {
  filename := "/home/wuxu/tmp/go-de.txt"
  content, err := Contents(filename)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(content)
  }


  deferRun()
  literal()
  arr :=[]int{1,2,3,4}
  newArr := arrayCopy(arr)

  // fmt.Println(arr, newArr)
  fmt.Printf("arr type: %T, newArr type: %T", arr, newArr)

  fmt.Println(twoDimenSlice())

  // mapOp()
  // constOp()
  // varOp()
  interfaceOp()
  addFuncOp()
}

