package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type region struct {
	Region string
	Hotels []hotel
}

type Regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := Regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{
					Name:    "Southside",
					Address: "1800 SW Ave",
					City:    "Compton",
					Zip:     "95612",
					Region:  "southern",
				},
				hotel{
					Name:    "Southside 2",
					Address: "1900 SW Ave",
					City:    "Beverly Hills",
					Zip:     "90210",
					Region:  "southern",
				},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{
					Name:    "Centered",
					Address: "1800 W Ave",
					City:    "Central City",
					Zip:     "9xxxx",
					Region:  "central",
				},
				hotel{
					Name:    "Centered 2",
					Address: "1900 W Ave",
					City:    "Central City",
					Zip:     "9xxxx",
					Region:  "central",
				},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{
					Name:    "North Pole",
					Address: "1800 SW Ave",
					City:    "Northern City",
					Zip:     "9xxxx",
					Region:  "northern",
				},
				hotel{
					Name:    "North Pole",
					Address: "1900 SW Ave",
					City:    "Northern City",
					Zip:     "9xxxx",
					Region:  "northern",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
