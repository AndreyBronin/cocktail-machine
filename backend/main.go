package main

import (
  "fmt"
)

type Ingredient struct {
  Name string
}


func main(){
  coke := Ingredient{"Coca-cola"}
  rum := Ingredient{"Rum"}
  whisky := Ingredient{"Whisky"}

  fmt.Println(coke, rum, whisky)
}