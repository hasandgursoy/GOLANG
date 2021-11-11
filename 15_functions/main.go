package main

import "fmt"

func main() {

	sum(5, 10)

	nameWriter("hasan")

	// Fonksiyon isimlendirmede eğer paket dışı bir işlemde kullanacaksak büyük harfle başlayacak.
	// Değikenler içinde aynı durum geçerli paket dışı kullanımda büyük harf kullanılacak. no camelCase

	


}

// Moduler Programing
// Readable Code
// From Complex to Simple

// func funcName (params) return type {code}

// retrun type olan bir function
func sum(x, y int) int {

	fmt.Println(x+y)
	return x + y

}

// return type olmayan function da tanımlayalım

func nameWriter(name string)  {
	
	fmt.Printf("Selamlar %v",name)
}