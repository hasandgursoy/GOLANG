package main

import (
	"fmt"
	"strings"
	//"strconv"
)

// struct --> underlying type
// employee --> Defined type, Named type


type mile float64
type kilometer float64
type kelime string
func main()  {
	
	var sayi mile = 3.2
	var sayi2 float64 = 3.5
	fmt.Println(sayi)
	fmt.Printf("%T %v",sayi,sayi)

	// kendi oluşturduğumuz tiple herhangi float değer toplayalım
	// Golang type güvenliği çok abartmış ciddiye almış bir dil sürekli tür dönüşümü yapmak zorundayız.
	fmt.Println(mile(sayi2) + (sayi))

	//deger,_ := strconv.ParseFloat("15",64)  deneme yaptım :D bu dersle alakası yok
	//deger := strconv.FormatFloat(sayi2,'f',6,64) // 1. kısım = sayi,2.kısım = tür,3.kısım = kaç basamklı devam etsin , 4.kısım = veritipi
	//fmt.Printf("%T %v",deger,deger)

	var m1 mile = 3.2
	var k1 kilometer = 7.8

	fmt.Printf("%T, %v",k1,k1)
	fmt.Printf("%T, %v",m1,m1)

	var s1 kelime = "hasan"

	fmt.Println(strings.ToUpper(string(s1)))

	t1 := mile(10)

	yol := toKilometer(t1)

	fmt.Println(yol)

	kilometre := toMile(yol)

	fmt.Println(kilometre)


}

func toKilometer (m mile)kilometer {
	return kilometer(m*1.6)
}

func toMile (k kilometer) mile {
	return mile(k * 0.62)

}