package main

import "fmt"
//begin interface
type Supe interface{
    fly() string
}
//end interface

//begin implementation
type DCHero struct{
    Name string
}
func (h *DCHero)fly() string {
    return  h.Name + " is flying fast."
}
type MarvelHero struct{
    Name string
}
func (h *MarvelHero)fly() string {
    return  h.Name + " is flying slow."
}
//end implementation
//begin main
func main(){
  var powerwondermanwoman Supe

  powerwondermanwoman = &MarvelHero{
    Name: "mega_it",
  }
  fmt.Println(powerwondermanwoman.fly())

  powerwondermanwoman = &DCHero{
    Name: "megapro_it",
  }
  fmt.Println(powerwondermanwoman.fly())
}
//end main
