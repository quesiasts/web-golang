package main

import (
	"net/http"
	"web-golang/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
