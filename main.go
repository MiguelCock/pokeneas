package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/pokenea/info",info)

	http.HandleFunc("/pokenea/image", image)

	log.Println("Servidor iniciado en el puerto 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
