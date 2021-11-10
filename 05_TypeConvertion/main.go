package main

import "fmt"

func main() {

	x := 10
	y := 10.0

	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", y, y)

	// Type Conversion uygulama şekli = type(value) olay bu kadar kardeşim.

	fmt.Println(x + int(y))

	var a int8 = 10
	var b int16 = 10

	// Strong type harbi strong muş  int8 ile int16 yı toplamadı la bu dil :D

	fmt.Println(a + int8(b))

	fmt.Println(int64(a) + int64(b))

	// Küçük olan veri tipini büyük olana dönüştürmek her zaman daha sağlıklıdır.Biraz maliyetlidir ancak güvenlidir.

	/*
	   q := 10
	   w := "10" 
	*/

	// fmt.Println(q+int(w)) type conversion yöntemi ile string ifade int'e dönüşmez.

	


}
