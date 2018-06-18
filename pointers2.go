package main

import "fmt"

func number(aPtr *int){
	*aPtr = 10;
}

func main() {
	//initializes bPtr location with int type its value will be zero
	bPtr := new (int);
	fmt.Println("value at bPtr : ",*bPtr);
	fmt.Println("value of bPtr : ",bPtr);

	number(bPtr);
	
	fmt.Println("value of bPtr : ",bPtr);
	fmt.Println("value at bPtr after number() : ",*bPtr);

}
