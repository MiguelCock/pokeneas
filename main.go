package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/pokenea/info",info)

	http.HandleFunc("/pokenea/image", image)

	log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
