# Scraper de Recettes - Cuisine Libre

Ce projet est un scraper développé en Go pour extraire et sauvegarder des recettes de cuisine à partir du site [Cuisine Libre](https://www.cuisine-libre.org). Le projet utilise la bibliothèque `gocolly` pour le scraping et la bibliothèque `encoding/json` pour sérialiser les données extraites en JSON.

## Fonctionnalités

- **Extraction des recettes** : Le programme récupère les titres, les images, les informations (durée de préparation, cuisson, repos), les ingrédients et les étapes de préparation des recettes à partir de la page de liste de recettes.
- **Vérification de la licence** : Avant de sauvegarder une recette, le programme vérifie que la licence associée est CC0 ou Domaine Public. Les recettes qui ne respectent pas ces conditions ne sont pas sauvegardées.
- **Gestion des doublons** : Les recettes déjà extraites et sauvegardées dans un fichier JSON ne sont pas ré-extraites. Cela permet d'éviter les doublons lors des exécutions successives du programme.
- **Sauvegarde des recettes** : Les recettes sont sauvegardées dans un fichier `recette.json` au format JSON, avec toutes les informations extraites. Les URL des recettes non valides (licence incorrecte) sont enregistrées dans un fichier `recettes_a_exclure.json`.


### Étapes d'installation

1. Clonez le projet dans votre répertoire local :

```bash
git clone https://github.com/ton-utilisateur/recette-scraper.git
cd recette-scraper
```

2. Installez les dépendances Go nécessaires, notamment `gocolly` :

```bash
go get github.com/gocolly/colly
```

3. Compilez et exécutez le projet :

```bash
go run main.go
```

## Utilisation

Lorsque vous exécutez le programme, il extrait les recettes de la page spécifiée, les sérialise en JSON, les affiche dans le terminal et les enregistre dans un fichier `recette.json`.

### Exemple de commande :

```bash
go run main.go
```

### Affichage des recettes en JSON dans le terminal :

Les recettes extraites seront affichées au format JSON directement dans le terminal, et le fichier `recette.json` sera mis à jour.

### Éviter les doublons :

Le programme vérifie automatiquement si une recette existe déjà dans `recette.json`. Si une recette est déjà présente, elle ne sera pas re-scrapée.

### Gestion des licences :

Si une recette ne possède pas une licence CC0 ou Domaine Public, elle est exclue et l'URL de cette recette est enregistrée dans le fichier `recettes_a_exclure.json`.

## Fichiers JSON générés

- **recette.json** : Contient toutes les recettes valides extraites.
- **recettes_a_exclure.json** : Contient les URLs des recettes dont la licence n'est pas valide (CC0 ou Domaine Public).

## Structure des données JSON

Les données des recettes sont sauvegardées dans le fichier `recette.json` au format suivant :

```json
[
  {
    "titre": "Tarte aux pommes",
    "url": "https://www.cuisine-libre.org/tarte-aux-pommes",
    "url_image": "https://www.cuisine-libre.org/images/tarte-aux-pommes.jpg",
    "recette": {
      "titre": "Tarte aux pommes",
      "infos": {
        "duree_preparation": "30 minutes",
        "duree_cuisson": "45 minutes",
        "duree_repos": "",
        "methode_cuisson": "Four"
      },
      "ingredients": [
        "500g de pommes",
        "200g de farine",
        "100g de beurre",
        "50g de sucre"
      ],
      "etapes": [
        "Préparer la pâte.",
        "Étaler les pommes sur la pâte.",
        "Cuire au four."
      ],
      "licence": {
        "valide": true,
        "message": "La licence est CC0 ou Domaine public"
      }
    }
  }
]
```

## Améliorations futures

- **Pagination** : Ajouter une gestion de la pagination pour scraper plusieurs pages de recettes.
- **Filtres avancés** : Ajouter des filtres pour sélectionner les types de recettes à scraper (exemple : pâtisseries, plats principaux).
- **Interface utilisateur** : Intégrer une interface web pour rendre le scraping plus interactif.

- ##Contributeurs** : N'hésitez pas à proposer des améliorations via des Pull Requests !
## Licence

Ce projet est sous licence MIT
