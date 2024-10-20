package main

import (
	"encoding/json"
	"os"
	"text/template"
)

type Year struct {
	From int
	To   int
}

type Analytics struct {
	GoogleTag string
}

type Hero struct {
	Title     string
	Paragraph string

	CtaButton  string
	CtaButton2 string
}

type Feature struct {
	Title     string
	Paragraph string
	ImageURL  string
}

type Price struct {
	Currency     string
	Amount       int
	BillingCycle string

	Features []string
}

type Pricing struct {
	Enabled bool

	Title     string
	Paragraph string

	Prices []Price
}

type Cta struct {
	Enabled bool

	Title  string
	Button string
}

type Link struct {
	URL   string
	Title string
}

type Footer struct {
	Links  []Link
	Social []Feature
}

type FormItem struct {
	Question string
	Type     string
	Name     string
}

type Config struct {
	Title     string
	Language  string
	Copyright string
	Analytics Analytics

	Pricing Pricing

	Hero     Hero
	Features []Feature
	Tools    []Feature
	Cta      Cta

	Footer Footer

	Year Year

	Form []FormItem
}

func main() {

	var config Config
	// "GoogleTag": "G-4FHLVKD6PN"

	configFile, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}

	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		panic(err)
	}

	tmpl, err := template.New("index.html").ParseFiles("./templates/index.html")
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("./out/index.html")

	if err := tmpl.Execute(f, config); err != nil {
		panic(err)
	}
}
