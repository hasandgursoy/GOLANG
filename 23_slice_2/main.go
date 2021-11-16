package main

import "fmt"

func main() {

	// Bir slice'ı array üzerinden de oluşturabiliyoruz buna underArray'deniyor.

	underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(underArray)

	mySlc := underArray[2:5] // başlangıç dahil bitiş dahil değil
	fmt.Println(mySlc)

	mySlc2 := underArray[:6]
	fmt.Println(mySlc2)

	mySlc3 := underArray[0:]
	fmt.Println(mySlc3)

	mySlc4 := underArray[:]
	fmt.Println(mySlc4)

	// Burda garip bir olay var aynı underArray'den türetilen Array'lerin herhangi birinde değişiklik yaparsan 
	// Bütün arrayler değişir. :S 
	mySlc[0] =100
	
	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	fmt.Println(mySlc3)
	fmt.Println(mySlc4)


	fmt.Println("-----------  Slice İşlemleri --------------")

	slc1 := []int{}
	fmt.Println(slc1)

	slc1 = append(slc1, 3,4,5,6,7,12,3)
	fmt.Println(slc1)

	fmt.Println("--------- Append ile iki slice aynı anda ekleme ------------")

	a := []int{1,2,3}
	b := []int{1,2}

	a = append(a, b...) // yanına 3 tane nokta koyunca array'e array ekelenebiliyor.
	fmt.Println(a)

	fmt.Println("--------- ilk iki eleman silme ---------")
	
	c := []int {1,2,3,4,5,6,8}
	fmt.Println(c)
	c = c[2:]
	fmt.Println(c)

	fmt.Println("--------- son iki eleman silme -----------")

	c = c[:len(c)-2]
	fmt.Println(c)

	var cs []int

	fmt.Printf(" %#v",cs) // tipini vs de gösteriyor bu şekilde güzelmiş.

	fmt.Println()

	// son kez make metodu ile bir slice yapalım

	benimSliceim := make([]int,0)
	fmt.Printf(" %#v",benimSliceim)


}