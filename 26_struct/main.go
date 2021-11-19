package main

import "fmt"

type calisanlar struct {
	name string
	age int
	isMarried bool
	kids []string
}

func main() {

	var employee struct { // structure

		name      string
		age       int
		isMarried bool
	}

	/* fmt.Println(employee)
	fmt.Printf("%#v \n",employee) */
	fmt.Println(employee.age)

	employee.name = "Hasan"
	employee.age = 23
	employee.isMarried =true

	fmt.Printf("%#v \n",employee)
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.isMarried)

	// tamam güzel bir tane çalışan oluşturduk ancak 2.çalışanı da oluşturmak için böyle amele gibi kod mu yazıcaz. ? tabiki hayır
	// var ile değil type olarak oluşturcaz struct'ımızı ve sorun çözülcek

	type employ struct { // olay type ile oluşturmak bu kadar basit

		name string
		age int
		isMarried bool
	}

	var e1 employ 
	e1.name ="hasan"
	e1.age = 23
	e1.isMarried=false

	var e2 employ
	e2.name ="Zehra"
	e2.age=87
	e2.isMarried=false

	fmt.Printf("%#v \n",e1)
	fmt.Printf("%#v \n",e2)

	// type ları genelde func main dışında tanımlarız kısa not :>

	e3 := calisanlar {
		name : "Cengo",
		age :40,
		isMarried: false,
		kids: []string {
			"aslan",
			"kaplan",
			"inek",

		},

	}

	yazdirici(e3)


}

func yazdirici(employ calisanlar)  {
	
	fmt.Println("İsim :",employ.name)
	fmt.Println("Yaş :",employ.age)
	fmt.Println("Evli mi ? :",employ.isMarried)
	fmt.Println("Çocuklar :",employ.kids)
}