package main

import "fmt"

type Asset struct {
	assetName string
}

func main(){
	r := Asset{"Maruti 900"}
	var pointer *Asset;

	pointer = &r;

	fmt.Println(pointer.assetName);
}
