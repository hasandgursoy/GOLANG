package main

import "fmt"

// 1-) Underlying Type 'int' olacak şekilde kendi veri tipimizi oluşturalım

type myType int




func main()  {
	
	x := 10
	fmt.Printf("%T , %v \n",x,x)

	// f değişkeni oluşturalım ama tipi myType olsun.

	f := myType(15)
	fmt.Printf("%T %v \n",f,f)

	// 2-) Başlangıç değeri 10 olan bir x değişkenin bulunduğu adresi
	// y değişkenine atayıp y değişkeninin x değerini 20 yapınız.

	y := &x
	*y = 20
	fmt.Println(x,y)

	r1 := rectangle{
		a: 10,
		b :10,
	}

	r1.display()
	fmt.Println("Alanı :",r1.area())
	fmt.Println("Çevresi : ",r1.circumference())
	
	families := showFamily()

	for index,deger := range families{

		fmt.Printf("%v , %T \n",families[index],deger)
	}



}


// 3-) Underlying Type struct olan Rectangle type oluşturun
	// display , area , circumference metodlarını yazınız.

	type rectangle struct {
		a,b int

	}

// Aşşağıda farklı bir method yöntemi kullanıcaz nesneye . koydukdan sonra kendini göstermesi için
// aşşağıdaki yolu izlicez.

func (r rectangle) display() {

	fmt.Println("Kenar Uzunlukları : ",r.a, " ,",r.b," olan dikdörtgen")


}

func (r rectangle) area() int{

	return r.a * r.b
}

func (r rectangle) circumference() int {

	return 2 * (r.a+r.b)

}

// 4 -) family aile bireyleri şeklinde veri tipi oluşturunuz. Aile bireylerinin hepisini farklı
// şekilde tanımlayınız. Sonrasında for döngüsünde yazdırınız. field, age ,name, isMarried.

type family struct {
	name string
	age int 
	isMarried bool

}

func showFamily() []family{

	family1 := family {
		name : "hasan",
		age : 23,
		isMarried: false,

	}

	family2 := family {
		name :"Elis",
		age : 3,

	}

	family3 := family {
		
		"gurcan",40,true,
	}

	var family4 family

	family4.name="selin"
	family4.age =40
	family4.isMarried=true


	return []family {family1,family2,family3,family4}


}