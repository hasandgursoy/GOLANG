package main

import "fmt"

func main() {

	name := "hasan"
	surName := "gursoy"
	age := 23
	const isMarried = false
	isStudying := "graduate"
	isHandsome := true

	if name == "hasan" && surName == "gursoy" {

		if age < 20 {

			fmt.Println("İşine bak yaşın küçük.")

		} else {
			fmt.Println("Hoş geldin 23 yaşındaki hasan kardeş")

			if isMarried == true {

				fmt.Print("Hızlısın baya slow down")

			} else {

				fmt.Println("Aferin evlilik aptallıkdır.")

				if isStudying == "graduate" {

					fmt.Println("güzel en azından okumuşsun.")

					if isHandsome {

						fmt.Println("Tartışmaya açık bir konu olmadığı için sormuyorum bile sen : Yakışıklısın.")
					}
				} else {

					fmt.Println("Tebrik ediyorum okul vakit kaybıdır insanı cahilleştirir.")
				}
			}

		}

	}

	// if yapısının içinde bir değişken kullanıcaksak değişkeni if bloğunun içinde atayabiliyoruz.
	// bir statement da birden fazla işlem yapıyosak ";" ile ayırmalıyız

	if x := 25; x < 0 {
		fmt.Println(x, "negatif sayıdır.")
		fmt.Println(x, " karaktersizdir.")
	} else if x == 0 {

		fmt.Println(x, " 0'a eşttir.")
	} else {
		fmt.Println("x pozitif sayıdır.")
	}

}
