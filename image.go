package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func image(w http.ResponseWriter, r *http.Request) {
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
}