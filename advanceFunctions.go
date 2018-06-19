package main

import "fmt"

func add(args ...int){

	for index,val := range args{
		fmt.Println(index,val);
	}
	//fmt.Println("Printing args : ", args);
	//return args;
}


func main() {
	add(1,2,3,4);
	//fmt.Println("Printing values : ",vals);
}
