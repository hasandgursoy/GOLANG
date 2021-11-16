package main

import (
	"fmt"
	"math"
)

func main() {

	city1 := "istanbul"
	city2 := "roma"
	city3 := "tahran"
	city4 := "belgrad"

	

	cities := [4] string {city1,city2,city3,city4}

	fmt.Println(cities)
	fmt.Println(cities[0]) // zero based index
	fmt.Println(cities[1])
	fmt.Println(cities[3])
	
	fmt.Println(len(cities))

	// arrayler'e değer atamasak bile içine kendi değerini atar bu int için 0 string için "" olabilir.

	var myArr [5]int
	fmt.Println(myArr)

	myArr[0] = 100
	fmt.Println(myArr)
	myArr[len(myArr)-1] = 200 // son eleman :-

	fmt.Println(myArr)


	// ---------------------------------
	// Golang'de arraylerde array işareti içindeki sayı da veritipi dahilindedir.
	// Bu şu demek [5]int != [4]int dikkat edin uzunluğu aynı değildir demiyorum veri tipi farklıdır .

	var myArr1 [5]int
	var myArr2 [5]int

	/* if myArr1 == myArr2 {
		
		fmt.Println("Mesaj")
	} */

	fmt.Println(myArr1,myArr2)

	// For Range :D pythondan tanıdık geliyodur kulağa

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}

	println()

	cities[0] = "ANKARA"

	for i := 0; i < len(cities); i++ {
		fmt.Println(i, cities[i])
	}

	arr10 :=[10]int {0,1,2,3,4,5,6,7,8,9}

	arr10 = arrSquare(arr10)
	fmt.Println(arr10)


	// FOR RANGE

	for  index,number := range cities {

		fmt.Println(index," " ,number)
	}



}

func arrSquare(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		
		arr[i] = int(math.Pow(float64(arr[i]),2))

	}	

	
	return arr

}