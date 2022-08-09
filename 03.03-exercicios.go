/*Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.*/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	for num := 33; num <= 122; num++ {
		x := strconv.Itoa(num)
		fmt.Printf("\n%v %t", x, x)
	}
}
-------------------------------------------------------------------------------------------

/*Põe na tela: todos os números de 1 a 10000.*/
package main

import (
	"fmt"
)

func main() {

	for num := 1; num <= 10000; num++ {
		fmt.Printf("\n%v", num)
	}
}
-------------------------------------------------------------------------------------------
/*Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
Por exemplo: 65 U+0041 'A' U+0041 'A' U+0041 'A' 66 U+0042 'B' U+0042 'B' U+0042 'B' ...e por aí vai.*/
package main

import (
	"fmt"
)

func main() {

	for letra := 65; letra <= 90; letra++ {
		fmt.Printf("\n%v", letra)
		for rep := 1; rep <= 3; rep++ {
			fmt.Printf("\t%#U\n", letra)

		}
	}
}
-------------------------------------------------------------------------------------------

/*Crie um loop utilizando a sintaxe: for condition {}
Utilize-o para demonstrar os anos desde que você nasceu.*/

package main

import "fmt"

func main() {
	for ano := 1992; ano <= 2022; ano++ {
		fmt.Printf("\n%v", ano)
	}

}
-------------------------------------------------------------------------------------------
/*Demonstre o resto da divisão por 4 de todos os números entre 10 e 100*/

package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		a := i % 4

		fmt.Printf("\n%v / 4 sobra %v", i, a)

	}
}
-------------------------------------------------------------------------------------------

/*Crie um programa que demonstre o funcionamento da declaração if.*/
package main

import "fmt"

func main() {
	x := 10

	if x == 10 {
		fmt.Printf("O número é %v\n", x)
	}
}

-------------------------------------------------------------------------------------------

/*Utilizando a solução anterior, adicione as opções else if e else.*/

package main

import "fmt"

func main() {
	x := 11

	if x%2 == 0 {
		fmt.Printf("O número %v é par!\n", x)
	} else {
		fmt.Printf("O número %v é impar!\n", x)
	}
}

-------------------------------------------------------------------------------------------
/*Crie um programa que utilize a declaração switch, sem switch statement especificado.
Utilizando a solução anterior, adicione as opções else if e else.*/


package main

import "fmt"

func main() {
	x := 0

	switch {
	case x > 0:
		fmt.Printf("%v é positivo\n", x)
	case x < 0:
		fmt.Printf("%v é negativo\n", x)
	case x == 0:
		fmt.Printf("%v é zero\n", x)

	}
}
-------------------------------------------------------------------------------------------
/*Crie um programa que utilize a declaração switch, 
onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".*/


package main

import "fmt"

func main() {
	esporteFavorito := ""

	switch esporteFavorito {
	case "futebol":
		fmt.Println("eu prefito futebol")
	case "basquete":
		fmt.Println("eu prefito basquete")
	case "volei":
		fmt.Println("eu prefito volei")
	case "golf":
		fmt.Println("eu prefito golf")
	case "crossfit":
		fmt.Println("eu prefito crossfit")
	default:
		fmt.Println("Digite seu esporte favorito")
	}
}









