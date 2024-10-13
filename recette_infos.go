package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func extraireInfosRecette(url string) map[string]interface{} {
	recette := make(map[string]interface{})

	c := colly.NewCollector()

	c.OnHTML("footer#license", func(l *colly.HTMLElement) {
		licenceValid := verifierLicence(l)
		if licenceValid {
			fmt.Println("Licence valide")
		} else {
			fmt.Println("La licence n'est pas CC0 ou Domaine public")
		}
	})

	// Extraction du titre de la recette
	c.OnHTML("h1", func(h *colly.HTMLElement) {
		titre := nettoyerTexte(h.DOM.Contents().Nodes[0].Data)
		fmt.Println("Titre: ", titre)
		recette["titre"] = titre
	})

	// Extraction des informations de préparation, cuisson et repos
	c.OnHTML("p#recipe-infos", func(e *colly.HTMLElement) {
		dureePreparation := extraireDureeRecette(e, "duree_preparation")
		dureeCuisson := extraireDureeRecette(e, "duree_cuisson")
		dureeRepos := extraireDureeRecette(e, "duree_repos")

		methodCuissonBruit := e.DOM.Find("a").Text()
		methodeCuisson := nettoyerTexte(methodCuissonBruit)

		infos := map[string]string{
			"duree_preparation": dureePreparation,
			"duree_cuisson":     dureeCuisson,
			"duree_repos":       dureeRepos,
			"methode_cuisson":   methodeCuisson,
		}

		recette["infos"] = infos
		fmt.Println("Infos: ", infos)
	})

	// Extraction de la préparation
	c.OnHTML("div#preparation", func(p *colly.HTMLElement) {
		preparations := extrairePreparation(p)
		recette["preparations"] = preparations
		fmt.Println("Préparations: ", preparations)
	})

	// Extraction des ingrédients
	c.OnHTML("div#ingredients", func(e *colly.HTMLElement) {
		ingredients := extraireIngredients(e)
		recette["ingredients"] = ingredients
		fmt.Println("Ingrédients:", ingredients)
	})

	// Lancer la visite de la page
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return recette
}

// supprime les espaces et caractères non imprimables
func nettoyerTexte(t string) string {
	// Remplace les espaces non imprimables (comme &nbsp;) par des espaces normaux
	t = strings.ReplaceAll(t, "\u00a0", " ")
	// Supprime les espaces multiples par un seul espace
	t = strings.Join(strings.Fields(t), " ")
	return strings.TrimSpace(t)
}

// Extrait la durée (préparation, cuisson, etc.) à partir des éléments DOM
func extraireDureeRecette(e *colly.HTMLElement, className string) string {
	span := e.DOM.Find("span." + className)
	time := span.Find("time").Text()
	return nettoyerTexte(time)
}

// Extrait les ingrédients de la recette
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

// Extrait la préparation de la recette
func extrairePreparation(p *colly.HTMLElement) []string {
	var preparations []string

	p.DOM.Find("p#preparation").Each(func(i int, p *goquery.Selection) {
		preparation := nettoyerTexte(p.Text())
		preparations = append(preparations, preparation)
	})
	return preparations
}
