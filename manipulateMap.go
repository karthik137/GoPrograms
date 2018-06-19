package main

import "fmt"

var map2 = make(map[string]int);

func main(){

	var map1 = make(map[string]int);
	map1["one"] = 1;
	map2["two"] = 2;
	fmt.Println(map1,map2);

	value, ok := map1["one"]

	fmt.Println(value, ok);
}
