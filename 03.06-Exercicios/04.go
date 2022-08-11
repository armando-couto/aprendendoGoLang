package main

import "fmt"

func main ()  {
	notas := [] float64{7.4, 8.2, 10, 9.5}
	soma := 0.0

	for _, num := range notas{
		soma += num
	}
	media := float64(soma) / float64(len(notas))
	fmt.Printf("%v\n", media)

}
