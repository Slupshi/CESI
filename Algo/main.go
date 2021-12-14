package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
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

	blipbloop.RunMachine()

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
		distributeur.StopMachine()
	case "m":
		distributeur.MachineMaintenance(stocks)
		break
	default:
		isDrinkAvailable = false
		break
	}
}

func (distributeur Machine) RunMachine() {
	distributeur.SayHello()
	PrintBoissons()
	stocks := make(map[string]int)
	stocks["Eau"] = baseStocks
	stocks["Cafe"] = baseStocks
	stocks["The"] = baseStocks
	for {
		fmt.Print("Quelle boisson voulez vous ? ")
		fmt.Printf("%v Eau | %v Café | %v Thé\n", stocks["Eau"], stocks["Cafe"], stocks["The"])
		fmt.Print("Appuyez sur \"m\" pour entrer en mode maintenance\n")
		distributeur.GetInput(stocks)

		if distributeur.Drinks > 0 && isDrinkAvailable == true {
			distributeur.Serve(&distributeur.Drinks)
			commandeTime := time.Now().Format(time.Kitchen)
			fmt.Printf("Voici votre %v de %v, il me reste %v autres boissons\n", drinkType, commandeTime, distributeur.Drinks)
		} else if distributeur.Drinks > 0 && isDrinkAvailable == false {
			fmt.Printf("Il n'y a pas de boissons de ce type dans %v\n", distributeur.Name)
		} else {
			fmt.Printf("Il n'y a plus de boissons disponible dans %v\n", distributeur.Name)
		}

	}
}

func (distributeur Machine) StopMachine() {
	fmt.Print("Good Bye")
	os.Exit(0)
}

func (distributeur Machine) MachineMaintenance(stocks map[string]int) {
	for {
		fmt.Print("De quelle boisson voulez vous vous occupez ?\n")
		for index, value := range boissons {
			index++
			fmt.Printf("%v, ", value)
		}
		fmt.Print("\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		drinkType = scanner.Text()
		switch drinkType {
		case "Cafe":
			drinkType = scanner.Text()
			//distributeur.ManageCoffee(stocks)
			distributeur.ManageDrinks(stocks, "Cafe")
			break
		case "The":
			drinkType = scanner.Text()
			//distributeur.ManageTea(stocks)
			distributeur.ManageDrinks(stocks, "The")
			break
		case "Eau":
			drinkType = scanner.Text()
			//distributeur.ManageWater(stocks)
			distributeur.ManageDrinks(stocks, "Eau")
			break
		default:
			return
		}
	}

}

func (distributeur *Machine) ManageDrinks(stocks map[string]int, drinks string) {
	for {
		fmt.Print("Que voulez vous faire ? Add or Remove ?\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		manageType := scanner.Text()
		if manageType == "Add" || manageType == "add" {
			fmt.Print("Combien de boissons voulez vous ajouter ? \n")
			scanner.Scan()
			howMuch := scanner.Text()
			howMany, err := strconv.Atoi(howMuch)
			if howMany < 0 {
				fmt.Printf("%v Error : Action Impossible\n", err)
				break
			}
			stocks[drinks] = stocks[drinks] + howMany
			fmt.Printf("%v drinks has been added.\n", howMuch)
			fmt.Printf("Il y a maintenant %v boissons de ce type dans le distributeur %v\n", stocks[drinks], distributeur.Name)
			distributeur.Drinks = distributeur.Drinks + howMany
			fmt.Printf("Pour un total de %v boissons\n", distributeur.Drinks)

		} else if manageType == "Remove" || manageType == "remove" {
			fmt.Print("Combien de boissons voulez vous retirer ? \n")
			scanner.Scan()
			howMuch := scanner.Text()
			howMany, err := strconv.Atoi(howMuch)
			if howMany < 0 || howMany > stocks[drinks] {
				fmt.Printf("%v Error : Action Impossible\n", err)
				break
			}
			stocks[drinks] = stocks[drinks] - howMany
			fmt.Printf("%v drinks has been removed.\n", howMuch)
			fmt.Printf("Il y a maintenant %v boissons de ce type dans le distributeur %v\n", stocks[drinks], distributeur.Name)
			distributeur.Drinks = distributeur.Drinks - howMany
			fmt.Printf("Pour un total de %v boissons\n", distributeur.Drinks)
		} else {
			break
		}
	}
}
