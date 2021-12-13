package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	nbDrink          int    = 50
	courtoisie       string = "Bonjour je suis " + name
	drinkType        string
	isDrinkAvailable bool
)

const name string = "BlipBloop"

func main() {
	//fmt.Printf("%v il me reste %v boissons\n", courtoisie, nbDrink)
	//nbDrink = nbDrink - 49
	//fmt.Printf("Vous êtes vraiment groumand ! Il ne me reste que %v boisson", nbDrink)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%v je dispose de %v boissons (Cafe/The/Eau)\n", courtoisie, nbDrink)
	for {
		fmt.Print("Quelle boisson vouelez vous ? ")
		scanner.Scan()
		drinkType = scanner.Text()
		switch drinkType {
		case "Cafe":
			drinkType = scanner.Text()
			isDrinkAvailable = true
			break
		case "The":
			drinkType = scanner.Text()
			isDrinkAvailable = true
			break
		case "Eau":
			drinkType = scanner.Text()
			isDrinkAvailable = true
			break
		case "*":
			fmt.Print("Good Bye")
			return
		default:
			isDrinkAvailable = false
			break
		}
		if nbDrink > 0 && isDrinkAvailable == true {
			Serve(&nbDrink)
			fmt.Printf("Voici votre %v, il m'en reste %v\n", drinkType, nbDrink)
		} else if nbDrink > 0 && isDrinkAvailable == false {
			fmt.Printf("Il n'y a pas de boissons de ce type dans %v\n", name)
		} else {
			fmt.Printf("Il n'y a plus de boissons disponible dans %v\n", name)
		}

	}

}

func Serve(boissons *int) {
	*boissons--
}
