package answer

import (
	"fmt"
	"math/rand"
)

func Experiment10() {

	fmt.Println("\n\nExperiment 10")

	telco := []string{"Telkom", "Indosat", "XL", "Verizon", "AT&T", "Nippon", "Vodafone", "Orange", "KDDI", "Telefonica", "T-Mobile"}
	selectedIndex := rand.Intn(len(telco))
	company := telco[selectedIndex]

	fmt.Println("Company selected : ", company)

	findCompanyExperiment10(telco, company)
}

func findCompanyExperiment10(array []string, input string) {
	for i := 0; i < len(array); i++ {
		if array[i] == input {
			fmt.Println("Company found : ", array[i])
			break
		}
		fmt.Printf("Searching company... %d\n", (i + 1))
	}
}
