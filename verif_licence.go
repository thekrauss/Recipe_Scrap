package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func verifierLicence(p *colly.HTMLElement) bool {
	var licenceLinkText string
	p.DOM.Find("footer#license a[rel='license']").Each(func(i int, s *goquery.Selection) {
		licenceLinkText = nettoyerTexte(s.Text())
		fmt.Println("Texte du lien de la licence extrait :", licenceLinkText)
	})

	licenceValid := strings.Contains(strings.ToLower(licenceLinkText), "cc0") || strings.Contains(strings.ToLower(licenceLinkText), "domaine public")

	if !licenceValid {
		fmt.Println("La licence n'est pas CC0 ou Domaine public")
		return false
	}

	return true
}
