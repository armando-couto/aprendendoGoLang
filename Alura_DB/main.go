package main

import (
	"net/http"

	"awesomeProject/bancoDeDadosGo/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}