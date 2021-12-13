package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	drinkType        string
	isDrinkAvailable bool
	boissons         = []string{"Eau", "Cafe", "The"}
)

type Machine struct {
	Name   string
	Drinks int
}

const name string = "BlipBloop"
const baseStocks int = 10

func main() {
	blipbloop := Machine{
		Drinks: 30,
		Name:   "BlipBloop"}

	blipbloop.SayHello()
	PrintBoissons()
	stocks := make(map[string]int)
	stocks["Eau"] = baseStocks
	stocks["Cafe"] = baseStocks
	stocks["The"] = baseStocks
	for {
		fmt.Print("Quelle boisson voulez vous ? ")
		fmt.Printf("%v Eau | %v Café | %v Thé\n", stocks["Eau"], stocks["Cafe"], stocks["The"])
		blipbloop.GetInput(stocks)

		if blipbloop.Drinks > 0 && isDrinkAvailable == true {
			blipbloop.Serve(&blipbloop.Drinks)
			fmt.Printf("Voici votre %v, il me reste %v autres boissons\n", drinkType, blipbloop.Drinks)
		} else if blipbloop.Drinks > 0 && isDrinkAvailable == false {
			fmt.Printf("Il n'y a pas de boissons de ce type dans %v\n", blipbloop.Name)
		} else {
			fmt.Printf("Il n'y a plus de boissons disponible dans %v\n", blipbloop.Name)
		}

	}

}

/*func Serve(boissons *int) {
	*boissons--
}*/

func PrintBoissons() {
	fmt.Print("( ")
	s := boissons[:3]
	for index, value := range s {
		index++
		fmt.Printf("#%v %v, ", index, value)
	}
	fmt.Print(" )\n")
}

//Methodes
func (distributeur Machine) SayHello() {
	fmt.Printf("Bonjour je suis %v, je dispose de %v boissons parmi ceci ", distributeur.Name, distributeur.Drinks)
}

func (distributeur Machine) Serve(boissons *int) {
	*boissons--
}

func (distributeur Machine) GetInput(stocks map[string]int) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	drinkType = scanner.Text()
	switch drinkType {
	case "Cafe":
		drinkType = scanner.Text()
		if stocks["Cafe"] > 0 {
			isDrinkAvailable = true
			stocks["Cafe"]--
		} else {
			isDrinkAvailable = false
		}
		break
	case "The":
		drinkType = scanner.Text()
		if stocks["The"] > 0 {
			isDrinkAvailable = true
			stocks["The"]--
		} else {
			isDrinkAvailable = false
		}
		break
	case "Eau":
		drinkType = scanner.Text()
		if stocks["Eau"] > 0 {
			isDrinkAvailable = true
			stocks["Eau"]--
		} else {
			isDrinkAvailable = false
		}
		break
	case "*":
		fmt.Print("Good Bye")
		return
	default:
		isDrinkAvailable = false
		break
	}
}
