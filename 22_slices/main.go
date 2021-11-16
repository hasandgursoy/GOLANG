package main

import "fmt"

func main() {

	myArr := [3]int{1, 2, 3}
	fmt.Println(myArr)

	// Arrayi yukarıdaki gibide yazabiliriz yada

	// Bu şekilde eleman sayısı belirtmeden de yazabiliriz.
	myArr2 := [...]int{1, 2, 3}
	fmt.Println(myArr2)

	// Slice tanımlarken eleman sayısı tayin etmiyoruz.

	var mySlc []int
	fmt.Println(mySlc)

	/*  mySlc = make([]int, 4)
	fmt.Println(mySlc)

	Bu şekilde de slice tekrar array'e dönüşür.
	*/
	mySlc = append(mySlc, 5, 10, 20)
	fmt.Println(mySlc)

	// Önemli bir durum var açıklanması gereken Golang'deki Arrayler diğer dillerdeki arrayler gibi değildir.
	// Diğer dillerde Arrayler referans'ı kopyalar ancak Golang'de değerleri kopyalar referansa dokunmaz.
	// Ancak Slice Metot'da durum farklıdır. Slice'da referansı kopyalar.

	// Arrayler = Pass By Value
	// Slice = Pass By Reference
	

}
