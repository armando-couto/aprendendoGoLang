/*Escreva um programa que mostre um número em decimal, binário, e hexadecimal.*/
package main

import "fmt"

func main() {
	a := 1

	fmt.Printf("%d \n%#x \n%b", a, a, a)

}
------------------------------------------------------------------------------

/*Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
==
!=
<=
<
=*/

package main

import "fmt"

func main() {
	a := 10 == 110
	b := 100 > 90
	c := 90 >= 10+90
	d := a != b
	e := 50 < 10
	f := 10 <= 9

	fmt.Printf("%v \n%v \n%v \n%v \n%v \n%v", a, b, c, d, e, f)

}

---------------------------------------------------------------------------------
/*Crie constantes tipadas e não-tipadas.
Demonstre seus valores.*/

package main

import "fmt"

func main() {
	const a int = 10
	const b = 20
	fmt.Printf("%v %v", a, b)
}
---------------------------------------------------------------------------------

/*Crie um programa que:
Atribua um valor int a uma variável
Demonstre este valor em decimal, binário e hexadecimal
Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
Demonstre esta outra variável em decimal, binário e hexadecimal*/

package main

import "fmt"

func main() {
	x := 100
	fmt.Printf("decimal: %d \nbinario: %b \nhexadecimal: %#x", x, x, x)
	y := 1 >> x
	fmt.Printf("decimal: %d \nbinario: %b \nhexadecimal: %#x", y, y, y)

}

--------------------------------------------------------------------------------------

/*Crie uma variável de tipo string utilizando uma raw string literal.
Demonstre-a.*/

package main

import "fmt"

func main() {
	x := `Lucas     MAcedo     Viana`

	fmt.Printf("%v", x)

}
-------------------------------------------------------------------------------------

/*Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
Demonstre estes valores.*/

package main

import "fmt"

const (
	_ = 2022 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Printf("%v \n%v \n%v \n%v", b, c, d, e)

}
-------------------------------------------------------------------------------------








