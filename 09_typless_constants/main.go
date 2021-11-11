package main

import "fmt"

func main() {

	/* const x = 5

	fmt.Printf("%T , %v " ,x,x) */

	//const x int8 = 3

	const x = 3
	var y int16 = 12

	fmt.Printf("%T , %v " ,x,x)
	fmt.Printf("%T , %v " ,y,y)
	fmt.Printf("%T , %v " ,x+y,x+y)  // x 'e int8 tipini verip yazmıştım ve hata almıştım şimdi type kalıdırıcam o yüzden not aldım buraya
	



}