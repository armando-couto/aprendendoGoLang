/*tente acessar todos os itens de uma slice sem utilizar range.*/

package main

import "fmt"

func main() {
	sabores := []string{"portuguesa", "frango", "calabresa", "peperoni"}
	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
}
-----------------------------------------------------------------------------
