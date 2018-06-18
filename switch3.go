package main

import "fmt"
import "time"

func main(){
	switch time.Now().Weekday() {
		case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
			fmt.Println("It is a weekday");
		default:
			fmt.Println("It is weekEnd");
	}
}
