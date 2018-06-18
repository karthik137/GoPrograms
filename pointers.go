package main

import "fmt"

func func1(aPtr *int){
	*aPtr = 20;
}

func main() {
	a := 10;
	fmt.Println("value of a : ",a);
	func1(&a);
	fmt.Println("value of a after function call : ",a);


	//pointer operations
	var bPtr *int;
	bPtr = &a;
	*bPtr++;
	fmt.Println("value of a : ",a);

	*bPtr = *bPtr + 4;
	fmt.Println("value of a : ",a);

	fmt.Println("memory address of value a is : ",&a);
	fmt.Println("memory address of bPtr is : ",bPtr);

	
	//increment address
	//bPtr = bPtr++;

	//fmt.Println("memory address of bPtr is : ",bPtr);
}
