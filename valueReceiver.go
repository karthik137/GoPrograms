package main

import "fmt"

type cars string

func (myCar cars) setCar() {
	myCar = "maruti";
	fmt.Println(myCar);
}


func (myCar cars) getCar() {
	fmt.Println(myCar);
	//return myCar;
}

func main(){
	cars.setCar("");
}
