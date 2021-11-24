package main

import "fmt"

func main() {

	name := 22

	fmt.Println(name)
	fmt.Println(&name) // name değişkenin adresini gösteriyor & bu işaret koyunca.
	// & --> address operator

	fmt.Printf("%T , %v \n",name,name)
	fmt.Printf("%T , %v \n",&name,&name)

	// var y int = &name bunu bu şekilde atayamıyoruz kısa şekilde atayacağız.

	y := &name

	fmt.Printf("%T %v \n",y,y)
	

	z := 25

	fmt.Println(z) // 25
	fmt.Println(&z) // 25 'in adresi
	fmt.Println(*&z) // dereferencig tekrar 25 yazdırcak
	fmt.Println(&*&z) // bu seferde adresi göstercek tekrar

	fmt.Println(3 * 5)
	
	// Tekrar edelim 
	// & --> bellekdeki adresi gösterir
	// * --> bellekdeki değeri gösterir

	/* x1 := 10
	x2 := x1
	fmt.Println(x1 ,x2)
	x1 = 5
	fmt.Println(x1 ,x2) */

	x1 := 10
	x2 := &x1
	fmt.Println(x1,x2)
	fmt.Println(x1,*x2)

	*x2 = 3 // adresdeki değeri 3'e eşitle dersek x1 de 3 e eşit olur.

	fmt.Println(x1,*x2)


	d1 := [4]int {1,10,100,1000}
	d2 := d1

	fmt.Println(d1,d2)
	
	d2[0]=3
	fmt.Println(d1,d2) // Daha önce arraylerin pass by value oldugunu söylemiştik

	slc1 := []int{1,10,100,1000}
	slc2 := slc1

	fmt.Println(slc1,slc2)

	slc1[0] = 2
	
	fmt.Println("slc2 :",slc2) // slice Lar pass by referance :D 

	






	







}