package main

import (
  "fmt"
)

func main() {
  test()
}

func mimicError(key string) {
//   return fmt.Errorf("Mimic Error : %s", key)
  panic("Mimic Error")
}

func test() {
  fmt.Println("Start Test")

//   err := mimicError("1")
	mimicError("1")

  fmt.Println("Meluncur pak")
  fmt.Println("Siap meluncur pak")

  defer func() {
    fmt.Println("Start Defer")
 
    // if err != nil {
    //   fmt.Println("Defer Error:", err)
    // }
  }()

  fmt.Println("Masih 5 bab")
  fmt.Println("Mata tinggal 5 watt")

  fmt.Println("End Test")
}