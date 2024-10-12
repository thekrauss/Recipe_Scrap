package main

import (
	"log"

	"github.com/gocolly/colly"
)

func extraireListeRecette(url string) []map[string]string {

	var listeResultats []map[string]string
	c := colly.NewCollector()

	//les éléments de la liste des recettes
	c.OnHTML("div#recettes ul", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {

			a := el.DOM.Find("a")
			titre := nettoyerTexte(a.Find("strong").Text())
			urlRecette := a.AttrOr("href", "")
			imgSrc := a.Find("img").AttrOr("src", "")

			recette := map[string]string{
				"titre":     titre,
				"url":       urlRecette,
				"url_image": imgSrc,
			}

			listeResultats = append(listeResultats, recette)
		})

	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return listeResultats

}
