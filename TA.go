package main

import "fmt"

var (
 pita  string
 CC    byte
 EOP   bool
 index int
)

func START() {
 index = 0
 CC = pita[index]
 EOP = CC == '.'
}

func ADV() {
 index = index + 1
 CC = pita[index]
 EOP = CC == '.'

}

func min() {
 index -= 1
 CC = pita[index]
 EOP = CC == '.'
}

func main() {
 var nama string
 var kntl bool
 var cek1, cek2 string

 fmt.Scan(&pita)

 START()

 nama = ""
 cek1 += string(CC)
 for !EOP {
  nama += string(CC)
  ADV()
  if EOP {
   min()
   cek2 += string(CC)
   ADV()
  }
 }

 if cek1 == cek2 {
  kntl = true
 } else {
  kntl = false
 }

 fmt.Print(kntl)
}