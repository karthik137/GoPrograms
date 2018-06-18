package main;

import "fmt"

func main(){

	sliceArray := []int{1,2,3,4,5}

	fmt.Println("slice ==", sliceArray);

	for index := 0; index < len(sliceArray); index++ {
		fmt.Println(index, sliceArray[index]);
	}

}
