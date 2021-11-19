package main

import "fmt"

type employee struct { // underlying type

	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee  // OOP' deki gibi miras alıyor.
	hasDegree bool
}

// IS A Relaiton --> Klasik OOP
// HAS A Relation

func main() {

	// Structlar pass by value'dur

	m1 := manager{
		employee: employee{
			name:      "Ayşe",
			age:       28,
			isMarried: false,
		},
		hasDegree: true,
	}

	// İkinci tanımlama şekli

	m2:=manager{}
	m2.name = "Hazal"
	m2.age = 29
	m2.isMarried=false
	m2.hasDegree=true

	fmt.Println(m1)
	fmt.Println(m2)

	// Anonim Struct
	// Sadece bir tane patron oldgunu düşünelim bir daha type oluşturmaya gerek yok :D
	theBoos := struct {

		name string 
		money bool
	}{name:"The BOSS",money: true}

	fmt.Println(theBoos)

}