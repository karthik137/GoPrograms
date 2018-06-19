package main

import "fmt"

type Asset struct{
	assetName string
}

func (asset *Asset) setAsset() string{
	asset.assetName = "myCar1";
	return asset.assetName;
}

func main(){
	var assetVar Asset;
	asset := assetVar.setAsset();
	fmt.Println("assetName : ",asset);
	fmt.Println("assetName : ",assetVar.assetName);
}
