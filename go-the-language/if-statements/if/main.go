package main

import "fmt"

type language int

const  (
  german language = iota
  english
  iceland
)

func main(){
  var lang language
  lang = german
  //begin only
  if lang == german {
    fmt.Println("Tick/Trick/Track")
  }
  //end only
  //begin stmt
  if lang:= f(); lang == english {
    fmt.Println("Huey, Dewey, and Louie")
  }
  //end stmt

  //begin else
  if lang == iceland {
    fmt.Println("Ripp/Rapp/Rupp")
  } else if lang == english {
    fmt.Println("Huey, Dewey, and Louie")
  } else if lang > iceland {
    fmt.Println("I don´t know")
  }
  //end else

  //begin scope
  if lang2:= f(); lang2 == german {
    // do something
  }else{
    fmt.Println(lang2) //is defined
  }
  //end scope

  //begin func
  if f() == german {
    fmt.Println("Tick/Trick/Track")
  }else{
    fmt.Println("Huey, Dewey, and Louie") //is defined
  }
  //end func

  //begin ternary
  var a int

  if test {
    a = 1
  } else {
    a = 2
  }
  //end ternary
}

//begin func

func f() language{
  return english
}
//end func
