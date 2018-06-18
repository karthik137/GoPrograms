package main

import "fmt"

func main(){
	x := 4;
	fmt.Println("This is : ",x);
	switch x {
		case 3:
			fmt.Println("This is case 3");
		case 5:
			fmt.Println("This is case 4");
		default:
			fmt.Println("This is case 5");
	}
}
