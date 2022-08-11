package main

import "fmt"

var pesoPescado float64 = 50.1 // informe o peso to total da pesca





func vaiTerMulta(peso float64){
	if  peso > 50.0 {
		resp := calculaMulta(peso)
		fmt.Printf("Voce tem que pagar uma multa no valor de R$:%.2f\n", resp )

	} else{
		fmt.Println("Não é necessario pagar a multa")
	}

}
func calculaMulta(valorPesado float64) float64{
	return (valorPesado - 50.0) * 4.0
}

func main ()  {
	vaiTerMulta(pesoPescado)
}
