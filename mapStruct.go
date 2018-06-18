package main

import "fmt"

type Asset struct{
	assetName string
}

func main() {
	var m = map[string]Asset{
		"car": Asset{"maruti 900"},
		"bike": Asset{"Hayabuza"},
	}

	fmt.Println(m["car"].assetName);
}
