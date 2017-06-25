package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type YamlItem struct {
	ID       int    `yaml:id`
	Singular string `yaml:singular`
	Plural   string `yaml:plural`
}

type YamlLocation struct {
	ID     string `yaml:id`
	Name   string `yaml:name`
	Prefix string `yaml:prefix`
}

func main() {
	fmt.Println("Creating items.go ...")
	writeItems()
	fmt.Println("Creating locations.go ...")
	writeLocations()
	fmt.Println("Done.")
}

func writeItems() {
	content, err := ioutil.ReadFile("items.yaml")
	if err != nil {
		panic(err)
	}

	items := make([]YamlItem, 0)

	if err := yaml.Unmarshal(content, &items); err != nil {
		panic("Could not unmarshal YAML.")
	}

	f, err := os.Create("../items.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("// This file has been generated.\n")
	f.WriteString("\n")
	f.WriteString("package data\n")
	f.WriteString("\n")
	f.WriteString("var Items = []Item{\n")

	for _, item := range items {
		f.WriteString(fmt.Sprintf("\tItem{ID: %d, Singular: %#v, Plural: %#v},\n", item.ID, item.Singular, item.Plural))
	}

	f.WriteString("}\n")
}

func writeLocations() {
	content, err := ioutil.ReadFile("locations.yaml")
	if err != nil {
		panic(err)
	}

	locations := make([]YamlLocation, 0)

	if err := yaml.Unmarshal(content, &locations); err != nil {
		panic("Could not unmarshal YAML.")
	}

	f, err := os.Create("../locations.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("// This file has been generated.\n")
	f.WriteString("\n")
	f.WriteString("package data\n")
	f.WriteString("\n")
	f.WriteString("var Locations = []Location{\n")

	for _, loc := range locations {
		f.WriteString(fmt.Sprintf("\tLocation{ID: %#v, Name: %#v, Prefix: %#v},\n", loc.ID, loc.Name, loc.Prefix))
	}

	f.WriteString("}\n")
}
