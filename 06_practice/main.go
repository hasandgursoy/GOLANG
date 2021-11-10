

// 1 : int x ,float64 y type conversion sample
/* package main

import "fmt"

func main()  {

	var x int = 75
	var y float64 = float64(x)

	fmt.Println(float64(x)+y)


*/

// 2 : multiple assing sample x,y = y,x
/* package main

import "fmt"

func main()  {

	x := 5
	y :=10

	fmt.Println("x : ",x,"Y : ",y)

	x,y = y,x

	fmt.Println("x : ",x,"Y : ",y)




} */

// 3 : non English variable names
/* package main

import "fmt"

func main()  {

	yaş := 40
	güneşDoğdumu := true

	fmt.Println("Yaşım : ",yaş,"günes doğdu mu ? : ",güneşDoğdumu)

	// hadi biraz abartalım madem utf-8 duyarlı bir dil


	哈桑 := "hasan"

	fmt.Println(哈桑)

} */

// 4 : shadowing kavramı ? gölgeleme

/* package main

import "fmt"

func main()  {

	x := 5

	if true {

		x = 10 // yukarıdaki x le burdaki x aynı şey yeniden init edip değer atamadık direk değerine müdahele ettik
		x++		// ancak gidip tekrar init edip atama yapsaydık assagıdaki print metodunda 5 yazmaya devam edicekdi
		fmt.Println(x+1)
	}

	fmt.Println(x)

}

*/

// 5 : 40 as a string
package main

import (
	"fmt"
	"strconv"
)

func main()  {
	
	x:= 40

	// strconv.Itoa() :D
	s := strconv.Itoa(x) // int ifadeyi string ifadeye çevirmek için kullanıyoruz.
	

	fmt.Printf("%v %T ",x,x )
	println()
	fmt.Printf("%v %T ",s,s )


	




}

