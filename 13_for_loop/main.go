package main

import "fmt"

func main() {

/* 	for i := 1; i <= 10; i++ {

		fmt.Println(i)

	} */

	// Golang'de while gibi döngüler yok onun yerine sadece for var herşeyi for ile yapıyoruz.

	// şimdi garip bir tanımlama yapalım

	/* i := 0 // değeri dışarıda tanımladık :D garip değil mi ?

	for ; i<=10 ;i +=5 {

		fmt.Println(i)
	} */

/* 	for {
		fmt.Println("Benim Adım Hasan")

		// Sonsuz Döngü çalıştırmamak lazım :D 
		// Çalıştırırsanız cntrl + c yapıp çıkabilirsiniz
	} */

	/* for i := 0; true ; i += 5 {
		
		// Buda sonsuz döngüye sokar true koyduk çünkü şarta :D 

		fmt.Println(i)

	} */

	
	// while tarzı bir kullanım yapıcaz şimdi for herşeye yarıyor abi :D 

/* 	i := 10

	for i >= 0 {

		fmt.Println(i)
		i--

	} */

	
	for i := 0 ; i <=10; i+=1 {
		
		if i%3 ==0 {
			
			continue
		}else if i ==5 {
			break
		}

		fmt.Println(i)

	}


}