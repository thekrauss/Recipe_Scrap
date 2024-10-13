package main

import (
	"encoding/json"
	"log"
)

const (
	JSON_FILENAME         = "recette.json"
	JSON_EXCLUDE_FILENAME = "recettes_a_exclure.json"
)

func main() {
	/*
		url_recette := "https://www.cuisine-libre.org/gateau-au-chocolat-granuleux"
		recette := extraireInfosRecette(url_recette)

	*/

	//les recettes déjà sauvegardées
	listeRecettesSauvegardees, err := chargerFichierJSON(JSON_FILENAME)
	if err != nil {
		log.Fatalf("Error loading saved recipes listeRecettesSauvegardees: %v", err)
	}

	//les URLs de recettes à exclure
	urlsRecettesAExclureSauvegardees, err := chargerFichierJSON(JSON_EXCLUDE_FILENAME)
	if err != nil {
		log.Fatalf("Error loading URLs to exclude urlsRecettesAExclureSauvegardees: %v", err)
	}

	// récupérer les URL déjà exclues
	urlsRecettesAExclure := []string{}
	for _, r := range urlsRecettesAExclureSauvegardees {
		urlsRecettesAExclure = append(urlsRecettesAExclure, r["url"].(string))
	}

	// recupere la liste des recettes
	urlListeRecette := "https://www.cuisine-libre.org/boulangerie-et-patisserie?mots%5B%5D=83&lang=&max=10"
	listRecette := extraireListeRecette(urlListeRecette)

	var nouvellesRecettes []map[string]interface{}
	for _, recette := range listRecette {
		if !recetteExiste(recette, listeRecettesSauvegardees) {
			nouvellesRecettes = append(nouvellesRecettes, recette)
		}
	}

	// ajouter les nouvelles recettes aux recettes sauvegardées
	listeRecettesSauvegardees = append(listeRecettesSauvegardees, nouvellesRecettes...)

	jsonData, err := json.MarshalIndent(listeRecettesSauvegardees, "", "  ")
	if err != nil {
		log.Fatalf("Erreur lors de la sérialisation en JSON: %v", err)
	}

	// Imprimer les données JSON console
	log.Println(string(jsonData))

	// sauvegarde la liste mise à jour dans le fichier JSON
	err = sauvegarderFichierJSON(JSON_FILENAME, listeRecettesSauvegardees)
	if err != nil {
		log.Fatalf("Error saving recipes: %v", err)
	}
	log.Printf("Data was successfully saved to file %s!\n", JSON_FILENAME)

	// sauvegarde les nouvelles URLs à exclure
	err = sauvegarderFichierJSON(JSON_EXCLUDE_FILENAME, []map[string]interface{}{})
	if err != nil {
		log.Fatalf("Error saving URLs to exclude: %v", err)
	}

	log.Printf("URLs to exclude have been updated successfully!")
}
