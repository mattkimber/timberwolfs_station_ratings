package main

import (
	"embed"
	"encoding/json"
	"io"
	"os"
	"strings"
	"text/template"
)

type CargoType struct {
	Name           string `json:"name"`
	Switch         string `json:"switch,omitempty"`
	DefaultSprites bool   `json:"defsprites,omitempty"`
}

type Economy struct {
	Name       string
	Condition  string      `json:"condition"`
	CargoTypes []CargoType `json:"cargotypes"`
}

type EconomyCollection struct {
	Name      string
	IsFirst   bool
	Economies []Economy
}

//go:embed templates/*
var templates embed.FS

func main() {
	doSpritesets()
	doEconomyCollections()
}

func doEconomyCollections() {
	economyFile, err := os.Create("pnml/economies.pnml")
	defer economyFile.Close()

	economyCollections, err := os.ReadDir("economies")
	if err != nil {
		panic(err)
	}

	for idx, economyCollection := range economyCollections {
		if economyCollection.Name() != "default" {
			doEconomyCollection(EconomyCollection{Name: economyCollection.Name(), IsFirst: idx == 0 || idx == 1 && economyCollections[0].Name() == "default"}, economyFile)
		}
	}

	doEconomyCollection(EconomyCollection{Name: "default"}, economyFile)
}

func doEconomyCollection(collection EconomyCollection, buf io.Writer) {
	economies, err := os.ReadDir("economies/" + collection.Name)
	if err != nil {
		panic(err)
	}

	for _, economy := range economies {
		var def Economy

		jsonData, err := os.ReadFile("economies/" + collection.Name + "/" + economy.Name())
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(jsonData, &def)
		if err != nil {
			panic(err)
		}

		def.Name = strings.ReplaceAll(economy.Name(), ".json", "")

		collection.Economies = append(collection.Economies, def)
	}

	econTemplate := GetTemplate("economy", template.FuncMap{"upper": strings.ToUpper})

	econTemplate.Execute(buf, collection)
}

func doSpritesets() {
	entries, err := os.ReadDir("2x")
	if err != nil {
		panic(err)
	}

	spritesets := GetTemplate("spriteset", nil)

	spritesetFile, err := os.Create("pnml/spritesets.pnml")
	defer spritesetFile.Close()

	for _, entry := range entries {
		spritesets.Execute(spritesetFile, strings.Split(entry.Name(), "_")[0])
	}
}

func GetTemplate(name string, funcMap template.FuncMap) *template.Template {
	data, err := templates.ReadFile("templates/" + name + ".tmpl")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New(name).Funcs(funcMap).Parse(string(data))
	if err != nil {
		panic(err)
	}

	return tmpl
}
