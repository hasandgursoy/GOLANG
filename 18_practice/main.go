// 1-) İki rakam arasında toplama, çıkarma ve çarpma
// işleminin yapıldığı bir fonksiyon yazınız. package main

// 2-) Kullanıcı tarafından girilen nota göre geçtiniz
// veya kaldınız geri dönüşü yazdırınız.

// 3-) 1 ile yüz arasındaki bir sayıyı tahmin etme uygulaması
// yazınız. Toplam tahmin hakkınız 10 olsun.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"math/rand"
)

func main()  {
	
	sum,dif,multip:=calculation(105,47)

	fmt.Println("Toplam : ",sum)
	fmt.Println("Fark : ",dif)
	fmt.Println("Çarpım : ",multip)

	// 2. uygulama

	fmt.Println("Lütfen notunuzu giriniz : ")

	hesaplayici, _ := notHesapla()

	if hesaplayici < 50  {
		fmt.Println("Kaldınız")
	}else {
		fmt.Println("Geçtiniz ")
	}

	// 3. uygulama

	target := numRand(1,100)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalışın")

	for attemps := 0 ; attemps < 10; attemps++ {

		fmt.Println(10-attemps," hakkınız kaldı :")

		sayi , err:= reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}

		sayi = strings.TrimSpace(sayi)

		num,err := strconv.Atoi(sayi)

		if err != nil {
			fmt.Println(err)
		}else {
			if num == target && (num >1 && num <100)   {
				
				fmt.Println("Tebrikler bildiniz hemde ",attemps,". hak da <3")
				break

			}else if   !(num >1 && num < 100){
				fmt.Println("Girdiğiniz sayı 1-100 aralığının dışındadır. ")
			}else {

				fmt.Println("Girdiğiniz sayı tutmadı")

				if attemps != 9 {
					if num < target {
						fmt.Println("Sayıyı yükseltin")
					}else {
						fmt.Println("Sayıyı düşürün")
					}
				}else {
					fmt.Println("GameOver")
				}
			}
				
			
		}


	}

}


// 1. uygulama

func calculation (n1  , n2 float64) (sum float64, dif float64, multip float64){

	sum = n1 + n2
	multip = n1*n2
	dif = n1 - n2

	return sum,dif,multip



} 

// 2.Uygulama

func notHesapla() (int,error)  {
	
	reader := bufio.NewReader(os.Stdin)

	input,err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	input  = strings.TrimSpace(input) // başındaki ve sonundaki boşlukları siler.

	num ,err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
	}

	return num ,nil


}

// 3. Uygulama

func numRand(min,max int)int {

	
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min


}

