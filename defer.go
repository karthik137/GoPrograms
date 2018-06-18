package main
import "fmt"

func funA(){
	fmt.Println("This is function A ");
}

func funB(){
	fmt.Println("This is function B ");
}

func main(){
	defer fmt.Println("world");
	fmt.Println("hello");
	defer funA();
	defer funB();
}
