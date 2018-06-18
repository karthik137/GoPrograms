
package main

import "fmt"

func main(){

	stringSlice := make([]string, 2, 2);
	fmt.Println(stringSlice);
	stringSlice[1] = "hello";

	fmt.Println(stringSlice);
}
