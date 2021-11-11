// addition +
// substruction -
// product *
// division /
// remainder %

package main

import "fmt"

func main()  {
	
  //x,y := 15,10

 /* fmt.Printf("%T ,%v \n",x,x)
	fmt.Printf("%T ,%v \n",y,y)
	fmt.Printf("%T ,%v \n",x+y,x+y)
	fmt.Printf("%T ,%v \n",x-y,x-y)
	fmt.Printf("%T ,%v \n",x*y,x*y)
	fmt.Printf("%T ,%v \n",x/y,x/y)  // sonucu int olarak döndürüyor bu dil kafayı yemelik yemin ediyorum.
	fmt.Printf("%T ,%v \n ",x % y,x % y)*/
	
	// bizim şuan bunu float olarak alabilmemiz için tip dönüşümü yapmak lazım yada direk float bir ifade ile işlem yapmak gerekiyor.


	/* z := 5.0 / 2
	fmt.Printf("%T , %v \n",z,z) */

	// Increment ++ , Decrement -- POSTFIX

	x :=10

	fmt.Println(x)

	x +=1

	fmt.Println(x)

	x++
	x--
	x--

	fmt.Println(x)
	
	// fmt.Println(x++)  bunu bu şekilde yazdıramayız çünkü increment ve decrement bir statement'dır önceden yapmamız lazımdı.








}