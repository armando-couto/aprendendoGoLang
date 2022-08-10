/*Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
Nome
Sobrenome
Sabores favoritos de sorvete
Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.*/

package main

import "fmt"

func main() {
	type pessoa struct {
		nome      string
		sobrenome string
		sorvete   []string
	}
	pessoa1 := pessoa{nome: "Lucas", sobrenome: "macedo", sorvete: []string{"pistache", "chocolate"}}
	pessoa2 := pessoa{nome: "Mateus", sobrenome: "macedo", sorvete: []string{"morango", "Flocos"}}
	fmt.Println(pessoa1.nome)
	for _, v := range pessoa1.sorvete {
		fmt.Println(v)
	}
	fmt.Println(pessoa2.nome)
	for _, v := range pessoa2.sorvete {
		fmt.Println(v)
	}
}
----------------------------------------------------------------------------------------------------------------------
/*rie um novo tipo: veículo
O tipo subjacente deve ser struct
Deve conter os campos: portas, cor
Crie dois novos tipos: caminhonete e sedan
Os tipos subjacentes devem ser struct
Ambos devem conter "veículo" como struct embutido
O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
O tipo sedan deve conter um campo bool chamado "modeloLuxo"
Usando os structs veículo, caminhonete e sedan:
Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
Demonstre estes valores.
Demonstre um único campo de cada um dos dois*/


package main

import "fmt"

func main() {
	type veiculo struct {
		portas int
		cor    string
	}
	type caminhonete struct {
		veiculo
		tracaoNasQuatro bool
	}
	type sedan struct {
		veiculo
		modeloLuxo bool
	}
	carro1 := caminhonete{veiculo: veiculo{portas: 4, cor: "vermelho"}, tracaoNasQuatro: false}
	carro2 := sedan{veiculo: veiculo{portas: 2, cor: "azul"}, modeloLuxo: false}
	fmt.Println(carro1)
	fmt.Println(carro2)
	fmt.Println(carro1.portas)
	fmt.Println(carro2.modeloLuxo)

}
