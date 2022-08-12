package contas

import "awesomeProject/banco/clientes"

type ContaPoupanca struct {
	Titular clientes.Titular
	NumeroAgencia,NumeroConta,Operacao int
	saldo float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar{
		c.saldo -= valorDoSaque
		return "Saque autorizado"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar (valorDoDeposito float64) (string,float64) {
	if valorDoDeposito > 0{
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!",c.saldo
	} else{
		return "Esse valor não pode ser depositado!", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64  {
	return c.saldo
}
