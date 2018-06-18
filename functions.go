package main

import "fmt"

func firstFunc(a int){
	a = 0;
	fmt.Println("Value of a in firstFunc : ",a);
}

func main(){
	//short hand decl of var a
	a := 10;
	firstFunc(a);
	fmt.Println("Value of a in main : ",a);
	firstFunc(a);
}
