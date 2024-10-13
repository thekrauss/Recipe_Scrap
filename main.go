package main

import (
	"fmt"
)

func main() {
	/*
		url_recette := "https://www.cuisine-libre.org/gateau-au-chocolat-granuleux"
		recette := extraireInfosRecette(url_recette)
	*/

	url_list_recette := "https://www.cuisine-libre.org/boulangerie-et-patisserie?mots%5B%5D=83&lang=&max=10"
	list_recette := extraireListeRecette(url_list_recette)

	// afficher les recettes extraites
	for _, recette := range list_recette {
		fmt.Printf("Titre: %s\nURL: %s\nImage: %\nReccete: %s\n\n", recette["titre"], recette["url"], recette["url_image"], recette)
	}

}
