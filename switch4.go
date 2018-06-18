package main

import "fmt"
import "time"

func main(){

	time := time.Now()
	switch {
		case time.Hour() < 12:
			fmt.Println("Before noon");
		default:
			fmt.Println("afternoon");
	}
}
