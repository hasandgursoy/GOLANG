package main

import (
	"fmt"
)

func main() {

	switch grade := 3; grade {

	case 5: // if grade == 5 {fmt.Println("Pek iyi")}
		fmt.Println("Pek iyi ")

	case 4:
		fmt.Println("İyi")

	case 3:
		fmt.Println("Orta")
		y := 100
		fmt.Println(y)

	case 2:
		fmt.Println("Geçer")

	case 1:

		fmt.Println("Başarısız")

	default:
		fmt.Println("Geçersiz Not")
	}

}
