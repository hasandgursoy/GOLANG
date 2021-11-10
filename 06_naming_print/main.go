package main

// print işlemleri ve degisken isimlendirme


import "fmt"

// var Y = 100 //Bu şu anlama gelir bu değişkene diğer paketlerden de ulaşılabilir. Büyük harfle yazdık. dikkat !

func main() {

	// print - println - printf

	fmt.Print("Merhaba") // aynı satırda durur assagı inmez
	fmt.Println("Merhaba") // 1 satır assağı yazdırır.
	fmt.Printf("Merhaba")
	fmt.Println("")

	//name := "hasan"

	/* fmt.Print(name)
	fmt.Println(name)
	fmt.Printf(name) */

	/* fmt.Print("Benim Adım ",name)
	fmt.Println("")
	fmt.Println("Benim adım ",name)
	fmt.Printf("Benim Adım %v %T %X",name,name,name) */

	/* x := 100
	y := 20
	z := 30

	fmt.Printf("%b %d %o ",x,y,z) */
	
	//name, age := "Arin",5

	/* fmt.Print("Benim Adım ", name,", ve ben ",age, "yaşındayım.")
	fmt.Println("")
	fmt.Println("Benim adım,",name,", ve ben",age,"yasındayım.") */
	// println de degiken yazdırınca bosluk bırakmıyoruz stringin içinde çünkü gerek yok kendi bırakıyor. ama normal printde bunu yapmıyor.

	//fmt.Printf("Benim Adım %v ve ben %v yaşındayım.",name,age)


	// VISIBILITY

/* 	// Go da isimlendirme camelCase 'dir ve anlamlı olmalıdır.
	var coin string
	var customerName string

	// kısaltmalar büyük harfle yazılır

	var URL  //Url değil
	var HTTP // http değil */

	// benForDongusuDegiskeninIsmiyim böye bir isim kullanılmaz.
	
	

}