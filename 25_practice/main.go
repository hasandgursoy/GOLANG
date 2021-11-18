package main

import "fmt"

func main() {

	// 1-) 5 string elemandan oluşan bir array ve 5 int elemandan oluşan slice oluşturup index numaralarıyla beraber gösteriniz.

	fiveArray := [5]string{
		"tahran", "istanbul", "ankara", "napoli", "nepal",
	}

	for index, value := range fiveArray {

		fmt.Println(index,value)

	}

	fmt.Println("----------------------------")

	fiveSlc := []int {1,2,3,4,5}

	for i,v := range fiveSlc {
		fmt.Println(i,v)
	}

	// 2-) [0,1,2,3,4,5,6,7,8] slice dan [0,1,2,3], [4,5,6], [2,3,6,7] slicelarını oluşturunuz.

	baseSlc := []int {0,1,2,3,4,5,6,7,8}

	nSlc1 := baseSlc[:4]
	nSlc2 := baseSlc[4:7]
	nSlc3 := append(baseSlc[2:4],baseSlc[6:8]...) // 3 nokta operatoru ile slice 'ı parçaladık.


	fmt.Println(nSlc1,nSlc2,nSlc3)

	// 3-) slicelar için copy metodunu ve assign (=) ile farkını açıklayınız.

	/* mySlc := []int {1,2,3}
	mySlc2 := make([]int ,2)

	fmt.Println(mySlc)
	fmt.Println(mySlc2)
	
	copy(mySlc2,mySlc) // ilk slice 'a ikinci slice'ı kopyala
	fmt.Println(mySlc2)
	
	mySlc[0]=100
	fmt.Println(mySlc)
	fmt.Println(mySlc2) */

	// copy metodunda referance temel almaz yeni bir değer oluşturur.
	// birazdan aşşağıda assign metodunu yazıcaz orda referans temel alınır.

	mySlc := []int{1,2,3}
	mySlc2 := make([]int,2)

	fmt.Println(mySlc)
	fmt.Println(mySlc2)

	mySlc2 = mySlc

	fmt.Println(mySlc)
	fmt.Println(mySlc2)

	mySlc[0] = 100

	fmt.Println(mySlc)
	fmt.Println(mySlc2) // bunda da değer değişti çünkü assign ile yaptık işlemi.


	// 4-) map gösterimi ile yazar ve yazarlara ait kitapları for range ile gösteriniz.

	writersBooks := map[string][]string{

		"Stephan King" : {"Medyum","IT","Ruhlar Dükkanı"},
		"Noah Harari": {"21.yüzil 21 tavsiye","Sapiens","Homo Deus"},
		"Sabahattin Ali": {"Kuyucaklı Yusuf","Kürk Mantolu Madonna","Değirmen"},

	}

	fmt.Println(writersBooks)
	fmt.Println(writersBooks["Sabahattin Ali"])

	writersBooks["Sabahattin Ali"] = append(writersBooks["Sabahattin Ali"], "selamlar","hoşçakalın")

	fmt.Println(writersBooks["Sabahattin Ali"])
	fmt.Println(writersBooks["Sabahattin Ali"][0])

	for key,val := range writersBooks{
		
		fmt.Println("Yazarımız : ",key)
		fmt.Println("Bazı kitapları aşşağıdadır :")
		
		for i,v := range val{
			fmt.Println(i+1,v)
		}

	}

	


	



}