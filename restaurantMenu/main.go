package main

import (
	"log"
	"os"
	"text/template"
)

type menuItem struct {
	Name        string
	Price       float64
	Description string
}

type meal struct {
	Meal  string
	Items []menuItem
}

// Meals is a list of meals for menu
type menu []meal

type restaurant struct {
	Name string
	Menu menu
}

// Restaurants holds menus for different ret
type Restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := Restaurants{
		restaurant{
			Name: "Luigi's",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Items: []menuItem{
						menuItem{
							Name:        "Sunrise Scrambler",
							Price:       8.95,
							Description: "Scrambled eggs with bacon",
						},
						menuItem{
							Name:        "Blueberry Crepes",
							Price:       6.95,
							Description: "Blueberry crepes with a side of fruit",
						},
					},
				},
				meal{
					Meal: "Lunch",
					Items: []menuItem{
						menuItem{
							Name:        "BLT",
							Price:       6.95,
							Description: "BLT on white bread with chips",
						},
						menuItem{
							Name:        "Summer Salad",
							Price:       5.95,
							Description: "Summer salad with choice of dressing",
						},
					},
				},
				meal{
					Meal: "Dinner",
					Items: []menuItem{
						menuItem{
							Name:        "Prime Rib",
							Price:       13.95,
							Description: "16 oz prime rib with choice of 2 sides",
						},
						menuItem{
							Name:        "Tofurky platter",
							Price:       6.95,
							Description: "Tofu turkey with choice of 2 sides",
						},
					},
				},
			},
		},
		restaurant{
			Name: "Retro Diner",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Items: []menuItem{
						menuItem{
							Name:        "Denver Omlette",
							Price:       8.95,
							Description: "Three eggs with bacon, onions and other",
						},
						menuItem{
							Name:        "Short Stack",
							Price:       6.95,
							Description: "Three pancakes",
						},
					},
				},
				meal{
					Meal: "Lunch",
					Items: []menuItem{
						menuItem{
							Name:        "BLT",
							Price:       6.95,
							Description: "BLT on white bread with chips",
						},
						menuItem{
							Name:        "Triple Cheeseburnger",
							Price:       5.95,
							Description: "It's a burger with three patties",
						},
					},
				},
				meal{
					Meal: "Dinner",
					Items: []menuItem{
						menuItem{
							Name:        "Philly Cheesesteak",
							Price:       13.95,
							Description: "Philly cheesesteak sandwich with extra onions",
						},
						menuItem{
							Name:        "The Widowmaker",
							Price:       56.95,
							Description: "As many burgers as you can eat",
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
