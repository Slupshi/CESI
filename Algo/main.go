package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	drinkType        string
	isDrinkAvailable bool
	boissons         = []string{"Eau", "Cafe", "The"}
	com              []string
	commands         [][]string
)

type Machine struct {
	Name            string `json:"Name"`
	Drinks          int    `json:"Nombre de boissons"`
	isInMaintenance bool   `json:"Le distributeur est en maintenance"`
}

const baseStocks int = 10

func main() {
	blipbloop := Machine{
		Drinks:          30,
		Name:            "BlipBloop",
		isInMaintenance: false}

	blipbloop.RunMachine()

}

func PrintBoissons() {
	fmt.Print("( ")
	s := boissons[:]
	for index, value := range s {
		index++
		fmt.Printf("#%v %v, ", index, value)
	}
	fmt.Print(" )\n")
}

//Methodes

func (distributeur Machine) JsonHTTP() {
	jsonMachine, _ := json.Marshal(distributeur)

	// Set routing rules
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(jsonMachine)
	})

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func (distributeur Machine) SayHello() {
	fmt.Printf("Bonjour je suis %v, je dispose de %v boissons parmi ceci ", distributeur.Name, distributeur.Drinks)
}

func (distributeur Machine) Serve(boissons *int) {
	*boissons--
}

func (distributeur Machine) PrintStocks(stocks map[string]int) {
	s := boissons[:]
	for _, value := range s {
		fmt.Printf("| %v %v ", stocks[value], value)

	}
	fmt.Print("|\n")
}

func (distributeur Machine) GetInput(stocks map[string]int) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	drinkType = scanner.Text()
	if drinkType == "*" {
		distributeur.StopMachine()
	} else if drinkType == "m" {
		distributeur.MachineMaintenance(stocks)
	} else if drinkType == "p" {
		distributeur.JsonHTTP()
	} else {
		s := boissons[:]
		for index, value := range s {
			if drinkType == s[index] {
				drinkType = value
				if stocks[drinkType] > 0 {
					isDrinkAvailable = true
					stocks[drinkType]--
					break
				} else {
					isDrinkAvailable = false
				}
			} else {
				isDrinkAvailable = false
			}
		}
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
		distributeur.isInMaintenance = false
		fmt.Print("Quelle boisson voulez vous ? ")
		distributeur.PrintStocks(stocks)
		fmt.Print("Appuyez sur \"m\" pour entrer en mode maintenance\n")
		distributeur.GetInput(stocks)

		if distributeur.Drinks > 0 && isDrinkAvailable == true {
			distributeur.Serve(&distributeur.Drinks)
			commandeTime := time.Now().Format(time.Kitchen)
			commandelog := fmt.Sprintf("[%v] 1 %v sur le distributeur %v", commandeTime, drinkType, distributeur.Name)
			com = append(com, commandelog)
			commands = append(commands, com)
			com = nil
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
		distributeur.isInMaintenance = true
		fmt.Print("De quelle boisson voulez vous vous occupez ?\n")
		distributeur.PrintStocks(stocks)
		fmt.Print("Press \"logs\" to save the commands logs\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		drinkType = scanner.Text()
		s := boissons[:]
		for index, value := range s {
			if drinkType == s[index] {
				drinkType = value
				distributeur.ManageDrinks(stocks, drinkType)
				break
			} else if drinkType == "logs" {
				distributeur.SaveLog()
			} else {
				distributeur.isInMaintenance = false
				return
			}
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

func (distributeur Machine) SaveLog() {
	file, err := os.Create("commands_logs.csv")
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range commands {
		writer.Write(value)
	}
}
