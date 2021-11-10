package main

import "fmt"

func main() {

	var name string = "Hasan"

	var age int = -255

	var isMarried bool = false

	fmt.Println(name,age, isMarried)

	secondName:= "dogacan"
	age2:=23
	isHandsome:=true

	fmt.Println(secondName,age2,"Yakışıklımı ? : ",isHandsome)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age) // type öğrenmek için %T kullanıyoruz.

	




}
