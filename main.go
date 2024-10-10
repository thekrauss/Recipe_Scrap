package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func nettoyerTexte(t string) string {
	return strings.ReplaceAll(strings.TrimSpace(t), "\u00a0", " ")
}

func extraireDureeRecette(e *colly.HTMLElement, className string) string {
	span := e.DOM.Find("span." + className)
	time := span.Find("time").Text()
	return nettoyerTexte(time)
}

func extraireInfosRecette(url string) map[string]interface{} {
	recette := make(map[string]interface{})

	c := colly.NewCollector()

	c.OnHTML("h1", func(h *colly.HTMLElement) {
		titre := h.Text
		fmt.Println("Titre: ", titre)
		recette["titre"] = titre
	})

	c.OnHTML("p#recipe-infos", func(e *colly.HTMLElement) {
		dureePreparation := extraireDureeRecette(e, "duree_preparation")
		dureeCuisson := extraireDureeRecette(e, "duree_cuisson")
		dureeRepos := extraireDureeRecette(e, "duree_repos")

		methodCuissonBruit := e.DOM.Find("a").Text()
		methodCuisson := nettoyerTexte(methodCuissonBruit)

		infos := map[string]string{
			"duree_preparation": dureePreparation,
			"duree_cuisson":     dureeCuisson,
			"duree_repos":       dureeRepos,
			"methode_cuisson":   methodCuisson,
		}

		recette["infos"] = infos

		fmt.Println("Infos: ", infos)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return recette
}

func main() {
	url := "https://www.cuisine-libre.org/gateau-au-miel-de-litha"
	recette := extraireInfosRecette(url)

	fmt.Println("Recette:", recette)
}
