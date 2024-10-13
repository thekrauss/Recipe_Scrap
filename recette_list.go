package main

import (
	"log"

	"github.com/gocolly/colly"
)

func extraireListeRecette(url string) []map[string]interface{} {

	var listeResultats []map[string]interface{}
	c := colly.NewCollector()

	//les éléments de la liste des recettes
	c.OnHTML("div#recettes ul", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {

			a := el.DOM.Find("a")
			titre := nettoyerTexte(a.Find("strong").Text())
			urlRecette := a.AttrOr("href", "")
			imgSrc := a.Find("img").AttrOr("src", "")
			/*
				if !strings.HasPrefix(urlRecette, "/") {
					urlRecette = "/" + urlRecette
				}
				if !strings.HasPrefix(imgSrc, "/") {
					imgSrc = "/" + imgSrc
				}
			*/

			urlRecette = el.Request.AbsoluteURL(urlRecette)
			imgSrc = el.Request.AbsoluteURL(imgSrc)

			recette := extraireInfosRecette(urlRecette)

			if recette != nil {

				resultat := map[string]interface{}{
					"titre":     titre,
					"url":       urlRecette,
					"url_image": imgSrc,
					"recette":   recette,
				}

				listeResultats = append(listeResultats, resultat)
			}

		})

	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return listeResultats

}
