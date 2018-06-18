package main

import "fmt"

type Rect struct {
	Width int
	Height int
}

func main(){

	fmt.Println(Rect{7,8});
	var temp Rect;
	temp.Width = 11;
	temp.Height = 12;
	fmt.Println(temp);
}
