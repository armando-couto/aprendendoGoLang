package main

import (
	"awesomeProject/banco/contas"
	"fmt"
)
func PagarBoleto(conta VerificarConta, valorBoleto float64){
	conta.Sacar(valorBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}


func main()  {

	exemplo := contas.ContaCorrente{}
	fmt.Println(exemplo)
	exemplo.Depositar(500)
	fmt.Println(exemplo.ObterSaldo())

	exemplo2 := contas.ContaPoupanca{}
	fmt.Println(exemplo2)
	exemplo2.Depositar(250)
	fmt.Println(exemplo2.ObterSaldo())

	PagarBoleto(&exemplo2,100)
	fmt.Println(exemplo2.ObterSaldo())

}
