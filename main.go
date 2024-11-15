package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Pokenea struct {
	ID              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Altura          float64 `json:"altura"`
	Habilidad       string  `json:"habilidad"`
	ImagenURL       string  `json:"imagen_url"`
	FraseFilosofica string  `json:"frase_filosofica"`
}

var pokeneas = []Pokenea{
	{1, "Jhon", 1.70, "Comer", "https://storage.cloud.google.com/pokeneas_kokc/FQMBEcZXwAcK7wv.png", "Vivir es como vivir"},
	{2, "Brayan", 1.60, "Jugar lol", "https://storage.cloud.google.com/pokeneas_kokc/FWG_2560x1600.jpg", "La violencia nunca es buena, mata la alma y la envenena"},
	{3, "La britani", 4.5, "Perrear hasta el piso", "https://storage.cloud.google.com/pokeneas_kokc/FyIFPo3WIAEP9ai.png", "no dejes para mañana lo que podes hacer hoy"},
	{4, "El sin papá", 1.3, "Tener PTSD", "https://storage.cloud.google.com/pokeneas_kokc/FyRCRGiX0AA427o.png", "skibiditoilet"},
	{5, "La aspiradora", 1.9, "Buffo con polvo blanco", "https://storage.cloud.google.com/pokeneas_kokc/FzHfwZqWAAAN7ux.png", "gyat"},
	{6, "Chorro", 2.5, "Drunken fist", "https://storage.cloud.google.com/pokeneas_kokc/GPoJkYkaAAAIFQg.png", "fiufiufiufiufui"},
	{7, "El pepas", 2.001, "Nulificacion del dolor", "https://storage.cloud.google.com/pokeneas_kokc/GTW2RQzXgAAkrjl.jpg", "me lo mecatee en cositas"},
	{8, "San juan", 0.4, "Vuelo rapido", "https://storage.cloud.google.com/pokeneas_kokc/GY913ogXkAAf2O_.png", "brainrot"},
}

func main() {
	http.HandleFunc("/pokenea/info", func(w http.ResponseWriter, r *http.Request) {
		randomIndex := rand.Intn(len(pokeneas))
		randomPokenea := pokeneas[randomIndex]

		containerID := os.Getenv("HOSTNAME")
		response := struct {
			Pokenea
			ContainerID string `json:"container_id"`
		}{
			Pokenea:     randomPokenea,
			ContainerID: containerID,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/pokenea/image", func (w http.ResponseWriter, r *http.Request) {
		randomIndex := rand.Intn(len(pokeneas))
		randomPokenea := pokeneas[randomIndex]

		containerID := os.Getenv("HOSTNAME")
		response := fmt.Sprintf(`
			<html>
			<meta charset="UTF-8">
			<body>
				<h1>Frase Filosófica: %s</h1>
				<img width="300" src="%s" alt="Imagen de %s" />
				<p>Container ID: %s</p>
			</body>
			</html>`,
			randomPokenea.FraseFilosofica,
			randomPokenea.ImagenURL,
			randomPokenea.Nombre,
			containerID,
		)

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, response)
	})

	log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
