package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {

	result , err:=evenNum(11)
	
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Girdiğiniz Sayı : ",result)
	}

	kareKok, err := sRoot(-62)

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(kareKok)
	}


	file , err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Dosyamız", file)
	}


}

func evenNum(num int) (int, error) {

	// error = bir interfacedir ileride görücez ve açıklayacaz

	if num%2 != 0 {

		return 0, errors.New("hata : Çift Sayı Girmediniz")
		
		// hata string'i Büyük harfle başlamamalı ve noktalama işaretleri ile bitmemeli
	}

	return num, nil

	// nil = başlangıç değeri olmayan ifadelere verilen addır.
	// nil 'in amacı yukarıda error da dönmesini istediğimiz için
	// ikinci kısımda error tipinde bir dönüş olmasını istiyoruz ancak 
	// hatayı if'in içinde yakaladık ama bizden hala 2 tane dönüş bekliyor.
	// int ve error o zaman error a bir ifade gidecek ancak bu hata olamaz çünkü 
	// hata if de yakalandı o zamanda boş bir ifade dönmesi gerekiyor.
	// nil in görevi bu kadar umarım anlatabilmişimdir. Biliyorum Try and Cath'den sonra
	// saçmalık olarak geliyor ancak böyle bir yapı kurmuş muhteşem 3'lü.


	
}

func sRoot(num float64) (float64,error){

	if num < 0 {
		return 0 , errors.New("error : Negatif Sayıların Karekökü Alınamaz")
	}else {

		 
		return math.Sqrt(float64(num)) , nil
	}

	

}

