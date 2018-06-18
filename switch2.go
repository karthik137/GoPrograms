package main

import "fmt"

func main(){

	length := 10;
	switch {
		case length<=7:
			fmt.Println("short");
		case length>=10:
			fmt.Println("Long");
	}
}
