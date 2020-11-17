package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var beers = map[string]string{
	"01D9X58E7NPXX5MVCR9QN794CH": "Mad Jack Mixer \n",
	"01D9X5BQ5X48XMMVZ2F2G3R5MS": "Keystone Ice \n",
	"01D9X5CVS1M9VR5ZD627XDF6ND": "Belgian Moon",
}

func main() {
	//param := os.Args[1] //Args -> x recollir parametres

	/*b := flag.Bool("beers", false, "print beers") //flag -> es un paquet. Permetra parsejar un d'aquests metodes.
	flag.Parse()                                  //s'ha de declarar flag.parse(), pq sigui capaç de llegir els flags
	//aixo retornara un puntero, i volem accedir al seu valor
	// es posarà *b
	if *b {
		fmt.Println(beers)
	}*/

	/* Flagset*/
	beersCmd := flag.NewFlagSet("beers", flag.ExitOnError)
	flag.Parse()

	//quan tens 2 parametres i has de mirar que existeix
	//NArg: qunats arguments hi ha a partir del nom
	if flag.NArg() == 0 {
		log.Fatal("You must specified a command beers")
		os.Exit(2)
	}

	//per recuperar aquest argument en concret
	switch flag.Arg(0) {
	case "beers":
		ID := beersCmd.String("id", "", "find by ID")
		beersCmd.Parse(os.Args[2:])

		if *ID != "" {
			fmt.Println(beers[*ID])
		} else {
			fmt.Println(beers)
		}
	}
}
