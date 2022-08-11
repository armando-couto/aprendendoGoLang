package main

import "fmt"

var HoraTrabalhada float64 = 50 // informe a quantidade de horas trabalhas no mes
var valorDaHora float64 = 120.0	// informe a valor da sua hora

func calSalarioBruto( hora, valor float64 ) (float64){
	return hora * valor
}

func impRenda() float64{
	return calSalarioBruto(HoraTrabalhada,valorDaHora) * 0.11
}

func inss() float64{
	return calSalarioBruto(HoraTrabalhada, valorDaHora) * 0.08
}

func sindicato() float64{
	return calSalarioBruto(HoraTrabalhada,valorDaHora) * 0.05
}

func desconto() float64{
	return calSalarioBruto(HoraTrabalhada,valorDaHora) - (impRenda() + inss() + sindicato() )
}


func main ()  {
	fmt.Printf("Salario Bruto R$: %.2f\n", calSalarioBruto(HoraTrabalhada, valorDaHora) )
	fmt.Printf("- IR R$: %.2f\n", impRenda() )
	fmt.Printf("- INSS R$: %.2f\n", inss() )
	fmt.Printf("- sindicato R$: %.2f\n", sindicato() )
	fmt.Printf("= Salario liquido R$: %.2f\n", desconto())

}
