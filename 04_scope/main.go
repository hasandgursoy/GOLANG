package main

import "fmt"

/* var packVar = "Package Scope" */

// kelime := "hasan" bu kod hata verir
// short and decleration package kısmında çalışmaz. (:=)
// paket düzeyinde bir değişken tanımlarsak hafızada sürekli olarak yer kaplar.
// func içinde tanımlarsak sadece blok çalıştığında çalışır.

func main() {

	if true {
		var blockvar = "Block Scope"
		fmt.Println(blockvar)
	}

	var funcVar string = "Func Scope"

	fmt.Println(funcVar)

	// Re Decleration 
	var name = "Sezar"
	name, surname := "Hasan","Bizans"

	// Şimdi durum şöyle name değerine tekrar atama yapmak istiyoruz diyelim.
	// Ancak bunu var name diyerek yapamayız yada name := şeklinde de yapamayız.
	// Bunu sadece 2 tane değeri tanımlarken tekrar yapabiliyoruz bunuda (:=) Short Decleration şeklinde yapabiliyoruz.
	
	fmt.Println(name,surname)

}
