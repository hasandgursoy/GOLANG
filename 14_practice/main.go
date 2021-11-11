package main

import "fmt"

// 1-)  1 ile 10 arasındaki sayıları if yapısıyla tek - çift olarak yazdırınız.

/* func main()  {

	for i := 1; i < 10; i++ {

		if i%2 ==0 {
			fmt.Println(i," çifttir")
		}else{
			fmt.Println(i," tektir")
		}
	}
} */

// 2-) for yapısını kullanarak Go'da olmayan while döngüsüne örnek veriniz.

/* func main()  {

	x := 10

	for x > 0  {

		fmt.Println(x)
		x--
	}
} */

// 3-) Switch fallthrough ifadesini açıklayınız.

// falltrough ifadesi : case'e bak doğru olabilir ancak sen yinde diğerlerinede bak anlamına gelir.
// normalde falltrough yazmasak diğer koşullara bakmadan switch bloğundan dışarı çıkar.
func main() {

/* 	switch x := 25; {

	case x < 20:
		fmt.Printf("%d küçüktür 20 \n", x)
		fallthrough

	case x < 50:
		fmt.Printf("%d küçüktür 50 \n", x)
		fallthrough

	case x < 100:
		fmt.Printf("%d küçüktür 100 \n", x)

	case x < 200:
		fmt.Printf("%d küçüktür 200 \n", x)

	} */

	// 5 -) 1 ile 50 arasındaki asal sayıları gösteren bir program yazınız.

	for i := 2; i < 50; i++ {
		deger := true
		for j := 2; j < i; j++ {

			if i%j == 0 {

				deger = false
			}
		}

		if deger  {
			fmt.Printf("%v is prime \n", i)
		}
	}

}
