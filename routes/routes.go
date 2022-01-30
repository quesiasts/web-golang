package routes

import (
	"net/http"
	"web-golang/controller"
)

func CarregaRotas() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.NewProduct)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)
}
