package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Téléchargement et sauvegarde d'une image depuis une URL
func telechargerEtSauvegarderImage(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors du téléchargement de l'image :", err)
		return
	}
	defer response.Body.Close()

	filename := url[strings.LastIndex(url, "/")+1:]
	indexPointInterrogation := strings.Index(filename, "?")
	if indexPointInterrogation != -1 {
		filename = filename[:indexPointInterrogation]
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde de l'image :", err)
	}
}
