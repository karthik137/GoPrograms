package main

import "fmt"

func firstFunc(a int){
	a = 0;
	fmt.Println("Value of a in firstFunc : ",a);
}


func secondFunc(a int) int {
	a = 10;
	return a;
}


func thirdFunc() (int,int){
	return 11,12;
}

func main(){
	//short hand decl of var a
	a := 10;
	firstFunc(a);
	fmt.Println("Value of a in main : ",a);
	firstFunc(a);

	b := secondFunc(20);
	fmt.Println("value of b is : ",b);

	c,d := thirdFunc();

	fmt.Println("value of c is ",c,"value of d is ",d);
}
