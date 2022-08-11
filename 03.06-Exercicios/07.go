package main

import "fmt"

func main ()  {
	sexo := "f" // m - masculino ou f - faminino
	altura := 1.6 //  ente 0.0 e 2.5  colocar em meltros

	if float64(altura) > 0.0 && float64(altura) < 2.5 {
		switch {
		case sexo == "m":
			peso := (72.7 * float64(altura)) - 58
			fmt.Printf("O seu peso ideal é: %.2fkg \n", peso)
		case sexo == "f":
			peso := (62.1 * float64(altura)) - 44.7
			fmt.Printf("O seu peso ideal é: %.2fkg \n", peso)
		default:
			fmt.Println("Coloque o sexo Corretamente")
		}
	} else{
		fmt.Println("Digite a altura corretamente")
	}
}
