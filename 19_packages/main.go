package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Lütfen not giriniz : ")

	reader := bufio.NewReader(os.Stdin)
	opener, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(opener)
	}


	fmt.Println(strings.Count("hasana","a"))
	fmt.Println(strings.ToUpper("hasan"))

	// Golang paketler üzerinden ilerleyen bir dildir şu ana kadar hep strandart library'i kullandık .
	// Bir sonraki dil de kendi paketlerimizi yazıcaz. Genel olarak buraya kadar import edilen paketler 
	// İşimizi hallediyordu ancak biraz eli kirletme vakti geldi :>
	// Paketler hayatı kolaylaştırır.
	// Paketler içinde fonksiyonlar barındırır.
	// Genel olarak bunları bilsek bu ders için yeterli dersler ilerledikçe olayları biraz daha heycanlaştırıcaz. Sabır.
	


}