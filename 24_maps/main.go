package main

import (
	"fmt"
	
)

func main() {

	myMap := map[string]int{

		"Ahmet":   40,
		"Ayşe":    17,
		"Selim":   14,
		"Mustafa": 30,
	}

	fmt.Println(myMap)
	fmt.Println(myMap["Ahmet"],myMap["Selim"])

	liste1 := []string {}
	liste2 := []int {}

	for key,val := range myMap {

		fmt.Println(key,val,"yaşındadır.")
		liste1 = append(liste1, key)
		liste2 = append(liste2, val)
		
	} 
	
	fmt.Println("liste1 :",liste1)
	fmt.Println("liste2 :",liste2)


	isMaried := map[string]bool{
		"Ahmet":true,
		"Ayşe":false,
		"Mahmut":false,
	}

	fmt.Println(isMaried)
	
	// Slice'larda oldugu gibi maplerdede make metodu kullanılabilir.

	/* newMap := make(map[string]int)

	fmt.Println(newMap) */

	stutendsGrades := map[string]int{
		"arin":80,
		"ahmet":29,
		"selim":72,
		"ayşe":77,
		"çınar":0,
	}

	fmt.Println(stutendsGrades["ayşe"])
	
	// map'de değer yoksa 0 döner ancak bu sorun oluşturabilir.
	fmt.Println(stutendsGrades["asdasd"])
	// Yukarıdaki durumu yönetmek lazım 

	value, ok := stutendsGrades["arin"]
// (deger , var,yok)
	fmt.Println(value,ok)


	_, oks :=stutendsGrades["elis"]
	fmt.Println(oks)
	
	fmt.Println(stutendsGrades)

	// Map'e eleman ekleme

	stutendsGrades["mahmut"] = 54
	stutendsGrades["ela"] = 23

	// Map eleman silme

	delete(stutendsGrades,"mahmut")
	fmt.Println(stutendsGrades)

	fmt.Println(len(stutendsGrades))

	// maps ==> pass by reference

	// bir mapde değişiklik yaparsan kendisiyle alakası olan herşey değişir.


	// Son bir for-range örnegi yapalım

	for k,v := range stutendsGrades{

		fmt.Println(k)
		fmt.Println(v)

	}

	


}