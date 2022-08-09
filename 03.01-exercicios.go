/*Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".

42
"James Bond"
true
Agora demonstre os valores nestas variáveis, com:

Uma única declaração print
Múltiplas declarações print*/
package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
------------------------------------------------------------------------------------------------------
/*Use var para declarar três variáveis. Elas devem ter package-level scope. 
Não atribua valores a estas variáveis. 
Utilize os seguintes identificadores e tipos para estas variáveis:
Identificador "x" deverá ter tipo int
Identificador "y" deverá ter tipo string
Identificador "z" deverá ter tipo bool
Na função main:
Demonstre os valores de cada identificador
O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?*/

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
	fmt.Printf("%v %T\n", z, z)
}
------------------------------------------------------------------------------------------------------

/*Utilizando a solução do exercício anterior:
Em package-level scope, atribua os seguintes valores às variáveis:
para "x" atribua 42
para "y" atribua "James Bond"
para "z" atribua true
Na função main:
Use fmt.Sprintf para atribuir todos esses valores a uma única variável. 
Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
Demonstre a variável "s".*/


package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v \n%v \n%v", x, y, z)
	fmt.Println(s)
}
------------------------------------------------------------------------------------------------------
/*Crie um tipo. O tipo subjacente deve ser int.
Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
Na função main:
Demonstre o valor da variável "x"
Demonstre o tipo da variável "x"
Atribua 42 à variável "x" utilizando o operador "="
Demonstre o valor da variável "x"*/

package main

import "fmt"

type subjacente int

var x subjacente

func main() {
	fmt.Printf("%v \n%T", x, x)
	x = 42
	fmt.Printf("\n%v", x)

}
------------------------------------------------------------------------------------------------------
/*Utilizando a solução do exercício anterior:
Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". 
O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
Na função main:
Isto já deve estar feito:
Demonstre o valor da variável "x"
Demonstre o tipo da variável "x"
Atribua 42 à variável "x" utilizando o operador "="
Demonstre o valor da variável "x"
Agora faça tambem:
Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
Demonstre o valor de "y"
Demonstre o tipo de "y"*/


package main

import "fmt"

type subjacente int

var x subjacente
var y int

func main() {
	fmt.Printf("x: %v \n%T", x, x)
	x = 42
	y = int(x)
	fmt.Printf("\ny: %v \n%T", y, y)

}
------------------------------------------------------------------------------------------------------













