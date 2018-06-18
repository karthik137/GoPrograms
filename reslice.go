package main

import "fmt"

func main(){

	slice := []int{1,2,3,4,5}
	fmt.Println("slice[1:4]  : ",slice[1:]);

	fmt.Println("slice[:4]   : ",slice[:4]);

	fmt.Println("slice[2:4]  : ",slice[2:4]);

}
