/*tente acessar todos os itens de uma slice sem utilizar range.*/

package main

import "fmt"

func main() {
	sabores := []string{"portuguesa", "frango", "calabresa", "peperoni"}
	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
}
----------------------------------------------------------------------------

/*Usando uma literal composta:
Crie um array que suporte 5 valores to tipo int
Atribua valores aos seus índices
Utilize range e demonstre os valores do array.
Utilizando format printing, demonstre o tipo do array.*/

package main

import "fmt"

func main() {
	num := [5]int{10, 20, 30, 40, 50}
	
	for i := 0; i < len(num); i++ {
		fmt.Println(i, num[i])
	}
	fmt.Printf("%T\n", num)
}
-----------------------------------------------------------------------------
/*Usando uma literal composta:
Crie uma slice de tipo int
Atribua 10 valores a ela
Utilize range para demonstrar todos estes valores.
E utilize format printing para demonstrar seu tipo.*/


package main

import "fmt"

func main() {
	num := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i := 0; i < len(num); i++ {
		fmt.Println(i, num[i])
	}
	fmt.Printf("%T\n", num)
}
----------------------------------------------------------------------------------
/*Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
Do quinto ao último item do slice (incluindo o último item!)
Do segundo ao sétimo item do slice (incluindo o sétimo item!)
Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item*/


package main

import "fmt"

func main() {
	num := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf(" primeiro ao  terceiro: %v\n", num[:3])
	fmt.Printf(" quinto ao ultimo: %v\n", num[4:])
	fmt.Printf(" segundo ao setimo: %v\n", num[1:7])
	fmt.Printf(" terceio ao penultimo: %v\n", num[2:9])
	fmt.Printf(" Penultimo Item usando a len: %v\n", num[2:len(num)-1])

}
----------------------------------------------------------------------------------
/*Começando com a seguinte slice:
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
Anexe a ela o valor 52;
Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
Demonstre a slice;
Anexe a ela a seguinte slice:
y := []int{56, 57, 58, 59, 60}
Demonstre a slice x.*/

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}
----------------------------------------------------------------------------------

/*Comece com essa slice:
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
Utilizando slicing e append, crie uma slice y que contenha os valores:
[42, 43, 44, 48, 49, 50, 51]*/
package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(append(x[:3], x[6:]...))
	fmt.Println(x)

}
----------------------------------------------------------------------------------
/*Crie uma slice usando make que possa conter todos os estados do Brasil.
Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", 
"Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", 
"Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", 
"São Paulo", "Sergipe", "Tocantins"
Demonstre o len e cap da slice.
Demonstre todos os valores da slice sem utilizar range.*/

package main

import "fmt"

func main() {
	estados := make([]string, 26, 26)
	estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo",
		"Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná",
		"Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima",
		"Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}
	fmt.Println(len(estados), cap(estados))
	for i := 0; i <= len(estados); i++ {
		fmt.Println(estados[i])
	}
}
----------------------------------------------------------------------------------
/*Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
"Nome", "Sobrenome", "Hobby favorito"
Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.*/
package main

import "fmt"

func main() {
	pessoas := [][]string{
		[]string{
			"jose",
			"marcos",
			"jogar",
		},
		[]string{
			"gabriel",
			"pereira",
			"malhar",
		},
		[]string{
			"everton",
			"henrique",
			"correr",
		},
	}
	for _, v := range pessoas {
		fmt.Println(v)
	}
}
----------------------------------------------------------------------------------
/*Crie um map com key tipo string e value tipo []string.
Key deve conter nomes no formato sobrenome_nome
Value deve conter os hobbies favoritos da pessoa
Demonstre todos esses valores e seus índices.*/

package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"lucas_Maceo":   []string{"Estudar", "treinar"},
		"Mateus_Macedo": []string{"Correr", "Pedalar"},
	}

	for k, v := range pessoas {
		fmt.Printf(k)
		for i, hobby := range v {
			fmt.Println("\t", i, " - ", hobby)
		}
	}
}
----------------------------------------------------------------------------------
/*Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.*/
package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"lucas_Maceo":   []string{"Estudar", "treinar"},
		"Mateus_Macedo": []string{"Correr", "Pedalar"},
	}
	pessoas["Rafaela_Nasser"] = []string{"dormir", "correr"}
	
	for k, v := range pessoas {
		fmt.Printf(k)
		for i, hobby := range v {
			fmt.Println("\t", i, " - ", hobby)
		}
	}
}
----------------------------------------------------------------------------------
/*Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.*/
package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"lucas_Macedo":  []string{"Estudar", "treinar"},
		"Mateus_Macedo": []string{"Correr", "Pedalar"},
	}
	pessoas["Rafaela_Nasser"] = []string{"dormir", "correr"}
	delete(pessoas, "lucas_Macedo")
	for k, v := range pessoas {
		fmt.Printf(k)
		for i, hobby := range v {
			fmt.Println("\t", i, " - ", hobby)
		}
	}
}



