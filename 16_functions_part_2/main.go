package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	var puan int

	fmt.Println("Lütfen Sınav sonucunu giriniz : ")
	fmt.Scan(&puan)

	fmt.Println(quizResult(puan))

	showTheTime()

	takeData()

	deger1,deger2:=bolme(104,10)

	fmt.Println(deger1,deger2)

	fmt.Println("Değer giriniz ")
	//reader := bufio.NewReader(os.Stdin)
	//value,_ := reader.ReadString('\n')

	strVeri :="154"
	result ,_:= strconv.Atoi(strVeri)
	fmt.Println(result)
	fmt.Printf("%d  %T eğerini giridiniz !",result,result)


}

// Multiple Return

func quizResult(point int) string {

	if point < 50 {
		return "Kaldınız !! puan : C seviyesinde"
	} else if point >= 50 && point < 75 {

		return "Geçtiniz !! puan : B seviyesinde"
	} else if point >= 75 && point <= 100 {
		return "GEÇTİNİZ : A+ seviyesi."
	} else {
		return "Saçma sapan bir değer girdiniz "
	}

}

func showTheTime() time.Time {

	fmt.Println(time.Now())
	return time.Now()
	

}

// bufio ve os kullanarak klavye verisini almak


func takeData()   {

	fmt.Print("Lütfen Yaşınızı giriniz: ")
	reader := 	bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // blank identifier = _

	// blank identifier hata mesajını atlamak için yaptık boş bir tarafa al hatayı dedik resmen :D 
	

	result,_ := strconv.Atoi((value))
	fmt.Println(result)

	


} 


func bolme (bolunen int , bolen int ) (bolum int , kalan int){


	bolum = bolunen / bolen
	kalan = bolunen % bolen

	return bolum, kalan
	
}