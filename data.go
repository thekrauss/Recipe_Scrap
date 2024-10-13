package main

import (
	"encoding/json"
	"os"
)

func sauvegarderFichierJSON(filename string, data []map[string]interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encode := json.NewEncoder(file)
	encode.SetIndent("", " ")
	return encode.Encode(data)
}

func chargerFichierJSON(filename string) ([]map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// le fichier n'existe pas, on renvoie une liste vide
			return []map[string]interface{}{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var data []map[string]interface{}
	decode := json.NewDecoder(file)
	err = decode.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
