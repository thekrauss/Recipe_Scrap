package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func extraireInfosRecette(urlRecette string) map[string]interface{} {
	recette := make(map[string]interface{})

	c := colly.NewCollector()

	//  titre
	var titre string
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		titre = nettoyerTexte(e.Text)
	})

	//  informations de la recette (durée, cuisson, repos)
	var dureePreparation, dureeCuisson, dureeRepos, methodeCuisson string
	c.OnHTML("p#recipe-infos", func(e *colly.HTMLElement) {
		dureePreparation = extraireDureeRecette(e, "duree_preparation")
		dureeCuisson = extraireDureeRecette(e, "duree_cuisson")
		dureeRepos = extraireDureeRecette(e, "duree_repos")
		methodeCuisson = e.DOM.Find("a").Text()
	})

	//  ingrédients
	var ingredients []string
	c.OnHTML("div#ingredients li.ingredient", func(e *colly.HTMLElement) {
		ingredient := nettoyerTexte(e.Text)
		if e.DOM.Find("i").Length() == 0 {
			ingredients = append(ingredients, ingredient)
		}
	})

	//  des étapes de préparation
	var etapes []string
	c.OnHTML("div#preparation p, div#preparation li", func(e *colly.HTMLElement) {
		etapes = append(etapes, nettoyerTexte(e.Text))
	})

	//  l'URL
	err := c.Visit(urlRecette)
	if err != nil {
		log.Printf("Failed to visit %s: %v", urlRecette, err)
		return nil
	}

	// map recette
	recette = map[string]interface{}{
		"titre": titre,
		"infos": map[string]string{
			"duree_preparation": dureePreparation,
			"duree_cuisson":     dureeCuisson,
			"duree_repos":       dureeRepos,
			"methode_cuisson":   nettoyerTexte(methodeCuisson),
		},
		"ingredients": ingredients,
		"etapes":      etapes,
	}

	return recette
}

// supprime les espaces et caractères non imprimables
func nettoyerTexte(t string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(t, "\xa0", " "), "\n", ""))
}

// a durée (préparation, cuisson, etc.)
func extraireDureeRecette(e *colly.HTMLElement, className string) string {
	span := e.DOM.Find("span." + className)
	time := span.Find("time").Text()
	return nettoyerTexte(time)
}

// ingrédients de la recette
func extraireIngredients(e *colly.HTMLElement) []string {
	var ingredients []string

	e.DOM.Find("li.ingredient").Each(func(_ int, s *goquery.Selection) {
		if s.Find("i").Length() == 0 {
			ingredient := nettoyerTexte(s.Text())
			ingredients = append(ingredients, ingredient)
		}
	})

	return ingredients
}

// préparation de la recette
func extrairePreparation(p *colly.HTMLElement) []string {
	var preparations []string

	p.DOM.Find("p#preparation").Each(func(i int, p *goquery.Selection) {
		preparation := nettoyerTexte(p.Text())
		preparations = append(preparations, preparation)
	})
	return preparations
}
