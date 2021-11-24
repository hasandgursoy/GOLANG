package main

import "fmt"

func main() {

	mySlc := []int{1, 10, 100}

	double(mySlc) // burda değer değişecek
	fmt.Println(mySlc)

	myArr := [3]int{1, 10, 100}

	double2(myArr) // burda array olduğu için değer değişmeyecek
	fmt.Println(myArr)

	x := 5 

	fmt.Println(x)
	double3(&x) // pointer methoda pointer değer vericeksin.
	fmt.Println(x)



}

func double(slice []int) {

	for i, _ := range slice {
		
		slice[i] *= 2


	}
	fmt.Println(slice)
}

func double2(slice [3]int) {

	for i, _ := range slice {
		
		slice[i] *= 2


	}
	fmt.Println(slice)
}

func double3(num *int){ // Bu artık bir pointer method 

	*num *=2
	fmt.Println(*num)

}
